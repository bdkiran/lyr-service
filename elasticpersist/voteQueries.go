package elasticpersist

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

//UpvoteLyricElastic handles upvoting lyrics in the elasticsearch database
func UpvoteLyricElastic(docID string, userID int) error {
	var r map[string]interface{}
	index := "music_lyrics"

	logger.Info.Printf("Attempting to upvote for document: %s", docID)

	updateBody := `{
		"script": {
		  "source": "if(ctx._source.upvotes == null){ctx._source.upvotes = [params.user]} else if(!ctx._source.upvotes.contains(params.user)) { ctx._source.upvotes.add(params.user)} else { ctx.op = 'none'}",
		  "lang": "painless",
		  "params": {
			"user": %d
		  }
		}
	  }`
	s := fmt.Sprintf(updateBody, userID)
	var sb strings.Builder
	sb.WriteString(s)
	read := strings.NewReader(sb.String())
	res, err := es.Update(index, docID, read)
	if err != nil {
		logger.Warning.Printf("Error with update query in elasticsearch: %s", err)
	}
	defer res.Body.Close()
	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			logger.Error.Printf("Error parsing the response body from elasticsearch: %s", err)
			return err
		}
		// Print the response status and error information.
		esErrorString := fmt.Sprintf("[%s] %s: %s",
			res.Status(),
			e["error"].(map[string]interface{})["type"],
			e["error"].(map[string]interface{})["reason"],
		)

		logger.Warning.Printf(esErrorString)
		return errors.New("no update was made by the ES script")
	}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		logger.Warning.Printf("Error parsing the response body: %s", err)
	}

	if r["result"] == "noop" {
		return errors.New("no update was made by the ES script")
	}
	//if updated return nil
	return nil
}

//UndoUpvoteLyricElastic handles undoing a like for a single user
func UndoUpvoteLyricElastic(docID string, userID int) error {
	var r map[string]interface{}
	index := "music_lyrics"

	logger.Info.Printf("Attempting to undo upvote for document: %s", docID)

	updateBody := `{
		"script": {
		  "source": "if(ctx._source.upvotes == null){ctx.op = 'none'} else if (ctx._source.upvotes.contains(params.user)) { ctx._source.upvotes.remove(ctx._source.upvotes.indexOf(params.user)) } else {ctx.op = 'none'}",
		  "lang": "painless",
		  "params": {
			"user": %d
		  }
		}
	  }`
	s := fmt.Sprintf(updateBody, userID)
	var sb strings.Builder
	sb.WriteString(s)
	read := strings.NewReader(sb.String())
	res, err := es.Update(index, docID, read)
	if err != nil {
		logger.Warning.Printf("Error with update query in elasticsearch: %s", err)
	}
	defer res.Body.Close()
	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			logger.Error.Printf("Error parsing the response body from elasticsearch: %s", err)
			return err
		}
		// Print the response status and error information.
		esErrorString := fmt.Sprintf("[%s] %s: %s",
			res.Status(),
			e["error"].(map[string]interface{})["type"],
			e["error"].(map[string]interface{})["reason"],
		)

		logger.Warning.Printf(esErrorString)
		return errors.New("No update was made by the ES script")
	}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		logger.Warning.Printf("Error parsing the response body: %s", err)
	}

	if r["result"] == "noop" {
		return errors.New("No update was made by the ES script")
	}
	//if updated return nil
	return nil
}
