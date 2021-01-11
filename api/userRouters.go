package api

import (
	"encoding/json"
	"net/http"

	"github.com/bdkiran/lyr-service/elasticpersist"
)

type vote struct {
	UserID     int64  `json:"user_id"`
	LyricDocID string `json:"docID"`
}

func upvoteHandler(w http.ResponseWriter, r *http.Request) {
	var upvote vote
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&upvote)
	if err != nil {
		logger.Error.Println("Error occured when processing message")
		sendResponse("Request error", "Unable to proccess form", http.StatusNotFound, w)
	}
	err = elasticpersist.UpvoteElastic(upvote.LyricDocID, upvote.UserID)
	if err != nil {
		sendResponse("Unable to Update", "Unable to update", http.StatusNotFound, w)
		return
	}
	sendResponse("Updated", "Updated based on request", http.StatusOK, w)
	return
}
