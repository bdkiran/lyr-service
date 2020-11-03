package api

import (
	"encoding/json"
	"net/http"
)

//create subscription struct
type subscription struct {
	Email string `json:"email"`
}

//create lyric submission struct
type lyricSubmission struct {
	SongName   string `json:"songName"`
	ArtistName string `json:"artistName"`
	AlbumName  string `json:"albumName"`
	Lyrics     string `json:"songLyrics"`
}

func subscriptionFormHandler(w http.ResponseWriter, r *http.Request) {
	var sub subscription
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&sub)
	if err != nil {
		logger.Error.Println("Error occured when processing message")
		sendResponse("Request error", "Unable to proccess form", http.StatusNotFound, w)
	}
	logger.Info.Println(sub)
	sendResponse("Message Proccesed", "Successfully processed form", http.StatusOK, w)
}

func lyricSubmissionFormHandler(w http.ResponseWriter, r *http.Request) {
	var lyrSub lyricSubmission
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&lyrSub)
	if err != nil {
		logger.Error.Println("Error occured when processing message")
		sendResponse("Request error", "Unable to proccess form", http.StatusNotFound, w)
	}
	logger.Info.Println(lyrSub)
	sendResponse("Message Proccesed", "Successfully processed form", http.StatusOK, w)
}
