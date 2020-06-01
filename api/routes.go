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

//Health handler function, should add more logic to actually provide a "Health" update
func healthHandler(w http.ResponseWriter, r *http.Request) {
	sendResponse("The server is alive", "No Errors detected", http.StatusOK, w)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	term, _ := vars["term"]

	logger.Info.Printf("Searching for: %s", term)
	data, err := elasticpersist.GetLyricsByTerm(term)
	if err != nil {
		sendResponse("Search Term Not Found", "Unable to find  for "+term, http.StatusNotFound, w)
	}

	response, _ := json.MarshalIndent(data, "", "    ")
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

type apiResponse struct {
	Message     string `json:"message"`
	Description string `json:"description"`
}

func sendResponse(message string, description string, statusCode int, w http.ResponseWriter) {

	responsePayload := apiResponse{
		Message:     message,
		Description: description,
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(responsePayload)
}
