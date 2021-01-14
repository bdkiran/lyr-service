package api

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//InitilizeRouter returns a new http Handler with all initilized routes, also providing cors access
func InitilizeRouter() http.Handler {
	//Get jwt token secret.....
	tokenInit()

	router := mux.NewRouter()
	router.HandleFunc("/", healthHandler).Methods("GET")
	router.HandleFunc("/search/{term}", searchHandler).Methods("GET")
	router.HandleFunc("/random", randomHandler).Methods("GET")
	router.HandleFunc("/random/{artist}", randomArtistHandler).Methods("GET")
	router.HandleFunc("/form/newsletter", subscriptionFormHandler).Methods("POST")
	router.HandleFunc("/form/lyricsub", lyricSubmissionFormHandler).Methods("POST")
	router.HandleFunc("/upvote", upvoteHandler).Methods("POST")
	router.HandleFunc("/upvote", undoUpvoteHandler).Methods("PUT")

	router.Use(authMiddleWare)

	headersOk := handlers.AllowedHeaders([]string{"Accept", "Accept-Language", "X-Requested-With", "Content-Type", "Authorization"})
	//change this to our website address...
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	returnCors := handlers.CORS(headersOk, originsOk, methodsOk)(router)

	return returnCors
}

func authMiddleWare(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		urlListToCheck := []string{"/upvote"}
		for _, uri := range urlListToCheck {
			if r.RequestURI != uri {
				h.ServeHTTP(w, r)
				return
			}
		}
		isAuthenticated, parseUserID := validateToken(w, r)
		if isAuthenticated {
			context.Set(r, "token", parseUserID)
			h.ServeHTTP(w, r)
		}
	})
}
