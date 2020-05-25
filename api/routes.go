package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bdkiran/lyr-service/elasticpersist"
	"github.com/gorilla/mux"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Home function called")
	const returnString = "Alive"
	response, _ := json.Marshal(returnString)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func singleArtistHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	artist, _ := vars["artistName"]
	log.Println(artist)

	data := elasticpersist.GetLyricstByArtistName(artist)

	response, _ := json.MarshalIndent(data, "", "    ")
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func singleSongHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	song, _ := vars["songName"]
	log.Println(song)

	data := elasticpersist.GetLyricstBySongName(song)

	response, _ := json.MarshalIndent(data, "", "    ")
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func songArtistHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	term, _ := vars["term"]
	log.Println(term)

	data := elasticpersist.GetLyricsByTerm(term)

	response, _ := json.MarshalIndent(data, "", "    ")
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
