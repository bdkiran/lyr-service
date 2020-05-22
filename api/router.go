package api

import "github.com/gorilla/mux"

//InitilizeRouter returns a new routher with all initilized routes
func InitilizeRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", healthHandler).Methods("GET")
	router.HandleFunc("/artist/{artistName}", singleArtistHandler).Methods("GET")
	router.HandleFunc("/song/{songName}", singleSongHandler).Methods("Get")
	return router
}
