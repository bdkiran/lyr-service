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

func healthHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info.Println("Home function called")
	const returnString = "Alive"
	response, _ := json.Marshal(returnString)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	term, _ := vars["term"]
	logger.Info.Printf("Searching for %s", term)

	data := elasticpersist.GetLyricsByTerm(term)

	response, _ := json.MarshalIndent(data, "", "    ")
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
