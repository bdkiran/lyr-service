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
	router.HandleFunc("/random", randomHandler).Methods("GET")
	router.HandleFunc("/random/{artist}", randomArtistHandler).Methods("GET")
	router.HandleFunc("/form/newsletter", subscriptionFormHandler).Methods("POST")
	router.HandleFunc("/form/lyricsub", lyricSubmissionFormHandler).Methods("POST")

	headersOk := handlers.AllowedHeaders([]string{"Accept", "Accept-Language", "X-Requested-With", "Content-Type", "Authorization"})
	//change this to our website address...
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	returnCors := handlers.CORS(headersOk, originsOk, methodsOk)(router)

	return returnCors
}
