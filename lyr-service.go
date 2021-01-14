package main

import (
	"net/http"
	"time"

	"github.com/bdkiran/lyr-service/api"
	"github.com/bdkiran/lyr-service/elasticpersist"
	"github.com/bdkiran/lyr-service/utils"
)

//Initilize variable to access project logger,
//this initialization can be used accoss the whole package
var logger = utils.NewLogger()

func main() {
	logger.Info.Println("Starting the server...")
	utils.LoadEnvVariables()
	//Set up connection to elasticearch
	elasticpersist.ConnectToEs()
	//Set up Http listener.
	handleRoutes()
}

func handleRoutes() {
	corsHandler := api.InitilizeRouter()
	srv := &http.Server{
		Handler:      corsHandler,
		Addr:         ":9000",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	logger.Error.Fatal(srv.ListenAndServe())
}
