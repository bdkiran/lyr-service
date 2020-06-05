package elasticpersist

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
)

//GetLyricsByTerm gets documents matchine a term
func GetLyricsByTerm(term string) ([]Lyric, error) {
	bufferedQuery, err := createBufferedQuery(term)
	if err != nil {
		logger.Warning.Print(err)
		return nil, err
	}

	elasticResponse, err := makeRequestToElasticSearch(bufferedQuery, "music_lyrics")
	if err != nil {
		logger.Warning.Print(err)
		return nil, err
	}

	lyrics, err := extractLyricsToReturn(elasticResponse)
	if err != nil {
		logger.Warning.Print(err)
		return nil, err
	}

	return lyrics, nil
}

func createBufferedQuery(keywordToSearch string) (bytes.Buffer, error) {
	//Generate a random salt string
	randomSalt := createRandomSeedString(10)
	// Build the request body.
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  keywordToSearch,
				"fields": []string{"artist", "title", "lyric"},
			},
		},
		"sort": map[string]interface{}{
			//First sorts by highest score
			"_score": map[string]interface{}{
				"order": "desc",
			},
			//Randomizes the order of the returned items, after score
			"_script": map[string]interface{}{
				"type": "number",
				"script": map[string]interface{}{
					"lang":   "painless",
					"source": "(doc['_id'].value + params.salt).hashCode()",
					"params": map[string]interface{}{
						"salt": randomSalt,
					},
				},
				"order": "asc",
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		logger.Warning.Printf("Error encoding query: %s", err)
		return buf, err
	}

	return buf, nil
}

func makeRequestToElasticSearch(buf bytes.Buffer, index string) (map[string]interface{}, error) {
	var r map[string]interface{}
	// Perform the search request.
	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex(index),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		logger.Warning.Printf("Error getting response from elasticsearch: %s", err)
		return r, err
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			logger.Error.Printf("Error parsing the response body from elasticsearch: %s", err)
			return r, err
		}
		// Print the response status and error information.

		esErrorString := fmt.Sprintf("[%s] %s: %s",
			res.Status(),
			e["error"].(map[string]interface{})["type"],
			e["error"].(map[string]interface{})["reason"],
		)

		logger.Warning.Printf(esErrorString)
		return r, errors.New(esErrorString)
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		logger.Warning.Printf("Error parsing the response body: %s", err)
		return r, err
	}

	// Print the response status, number of results, and request duration.
	logger.Info.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(r["took"].(float64)),
	)

	return r, nil
}

func extractLyricsToReturn(r map[string]interface{}) ([]Lyric, error) {
	//Array of lyrics that will be returned
	var returnData []Lyric
	// Print the ID and document source for each hit.
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		var songLyric Lyric
		data, err := json.Marshal(hit.(map[string]interface{})["_source"])
		if err != nil {
			logger.Warning.Println(err)
			return returnData, err
		}

		err = json.Unmarshal(data, &songLyric)
		if err != nil {
			logger.Warning.Println(err)
			return returnData, err
		}
		returnData = append(returnData, songLyric)
	}
	return returnData, nil
}
