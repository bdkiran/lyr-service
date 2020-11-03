package api

import (
	"encoding/json"
	"net/http"

	"github.com/bdkiran/lyr-service/elasticpersist"
	"github.com/bdkiran/lyr-service/utils"
	"github.com/gorilla/mux"
)

//Initilize variable to access project logger
var logger = utils.NewLogger()

func testHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info.Println("Message Recieved")
}

//Check Elasticsearch health
func healthHandler(w http.ResponseWriter, r *http.Request) {
	esStatusColor, err := elasticpersist.GetHealthOfCluster()
	if err != nil {
		sendResponse("Error", "Issue getting response from elasticsearch cluster", http.StatusOK, w)
		return
	}

	switch esStatusColor {
	case "green":
		sendResponse("Healthy", "No Errors detected", http.StatusOK, w)
		break
	case "yellow":
		sendResponse("Warning", "Issues reported with the elasticsearch cluster", http.StatusOK, w)
		break
	case "red":
		sendResponse("Error", "The elasticsearch cluster is down", http.StatusOK, w)
		break
	default:
		sendResponse("Healthy", "No Errors detected", http.StatusOK, w)
		break
	}
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	term, _ := vars["term"]

	logger.Info.Printf("Searching for: %s", term)
	data, err := elasticpersist.GetLyricsByTerm(term)
	if err != nil {
		sendResponse("Search Term Not Found", "Unable to find matches for "+term, http.StatusNotFound, w)
		return
	}

	response, _ := json.MarshalIndent(data, "", "    ")
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func randomHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info.Println("Getting random lyrics")
	data, err := elasticpersist.GetRandomLyrics()
	if err != nil {
		sendResponse("Something went wrong fetching random lyrics", "Unable to ger random lyrics", http.StatusNotFound, w)
		return
	}
	response, _ := json.MarshalIndent(data, "", "    ")
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
