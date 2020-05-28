package elasticpersist

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
)

//search indexed documents
func multiMatchSearchIndexedDocument(keywordToSearch string) []Lyric {
	var r map[string]interface{}

	// Build the request body.
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  keywordToSearch,
				"fields": []string{"artist", "title", "lyric"},
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		logger.Warning.Printf("Error encoding query: %s", err)
	}

	// Perform the search request.
	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("song_lyrics"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		logger.Warning.Printf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			logger.Warning.Printf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		logger.Warning.Printf("Error parsing the response body: %s", err)
	}

	// Print the response status, number of results, and request duration.
	logger.Info.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(r["took"].(float64)),
	)

	//Array of lyrics that will be returned
	var returnData []Lyric
	// Print the ID and document source for each hit.
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		var songLyric Lyric
		data, err := json.Marshal(hit.(map[string]interface{})["_source"])
		if err != nil {
			logger.Warning.Println(err)
		}

		err = json.Unmarshal(data, &songLyric)
		if err != nil {
			logger.Warning.Println(err)
		}
		returnData = append(returnData, songLyric)
		// log.Printf(" ID=%s", hit.(map[string]interface{})["_id"])
		// log.Printf(" Artist=%s", hit.(map[string]interface{})["_source"].(map[string]interface{})["artist"])
		// log.Printf(" Song=%s", hit.(map[string]interface{})["_source"].(map[string]interface{})["title"])
		// log.Printf(" Song=%s", hit.(map[string]interface{})["_source"].(map[string]interface{})["lyric"])
		// log.Printf(" LineNumber=%d", int(hit.(map[string]interface{})["_source"].(map[string]interface{})["lineNumber"].(float64)))
	}
	return returnData
}
