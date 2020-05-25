package api

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//InitilizeRouter returns a new routher with all initilized routes
func InitilizeRouter() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/", healthHandler).Methods("GET")
	router.HandleFunc("/artist/{artistName}", singleArtistHandler).Methods("GET")
	router.HandleFunc("/song/{songName}", singleSongHandler).Methods("GET")
	router.HandleFunc("/search/{term}", songArtistHandler).Methods("GET")

	corsObj := handlers.AllowedOrigins([]string{"*"})

	returnCors := handlers.CORS(corsObj)(router)

	return returnCors
}
