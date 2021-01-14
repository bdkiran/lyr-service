package api

import (
	"encoding/json"
	"net/http"

	"github.com/bdkiran/lyr-service/elasticpersist"
	"github.com/gorilla/context"
)

type vote struct {
	UserID     int    `json:"user_id"`
	LyricDocID string `json:"docID"`
}

func upvoteHandler(w http.ResponseWriter, r *http.Request) {
	var upvote vote
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&upvote)
	if err != nil {
		logger.Error.Printf("Error occured when processing message: %s", err)
		sendResponse("Request error", "Unable to proccess form", http.StatusNotFound, w)
		return
	}
	userIDFromToken := context.Get(r, "token")
	if userIDFromToken != upvote.UserID {
		sendResponse("Request error", "User_id is unauthroized", http.StatusNotFound, w)
		return
	}
	err = elasticpersist.UpvoteLyricElastic(upvote.LyricDocID, upvote.UserID)
	if err != nil {
		sendResponse("Unable to Update", "Unable to update", http.StatusNotFound, w)
		return
	}
	sendResponse("Updated", "Updated based on request", http.StatusOK, w)
	return
}

func undoUpvoteHandler(w http.ResponseWriter, r *http.Request) {
	var upvote vote
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&upvote)
	if err != nil {
		logger.Error.Println("Error occured when processing message")
		sendResponse("Request error", "Unable to proccess form", http.StatusNotFound, w)
		return
	}
	userIDFromToken := context.Get(r, "token")
	if userIDFromToken != upvote.UserID {
		sendResponse("Request error", "User_id is unauthroized", http.StatusNotFound, w)
		return
	}
	err = elasticpersist.UndoUpvoteLyricElastic(upvote.LyricDocID, upvote.UserID)
	if err != nil {
		sendResponse("Unable to Update", "Unable to update", http.StatusNotFound, w)
		return
	}
	sendResponse("Updated", "Updated based on request", http.StatusOK, w)
	return
}
