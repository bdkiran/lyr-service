package main

import (
	"log"
	"net/http"
	"time"

	"github.com/bdkiran/lyr-service/api"
	"github.com/bdkiran/lyr-service/elasticpersist"
)

func main() {
	elasticpersist.ConnectToEs()
	log.Println("Starting the server...")
	handleRoutes()

}

func handleRoutes() {
	router := api.InitilizeRouter()
	srv := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
