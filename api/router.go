package api

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//InitilizeRouter returns a new http Handler with all initilized routes, also providing cors access
func InitilizeRouter() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/", healthHandler).Methods("GET")
	router.HandleFunc("/search/{term}", searchHandler).Methods("GET")

	corsObj := handlers.AllowedOrigins([]string{"*"})

	returnCors := handlers.CORS(corsObj)(router)

	return returnCors
}
