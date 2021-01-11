package api

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("mySecret")

type claims struct {
	ID string `json:"user_id"`
	jwt.StandardClaims
}

type authenticationData struct {
	userID string `json:"user_id,omitempty"`
	Token  string `json:"token,omitempty"`
}

func validateToken(w http.ResponseWriter, r *http.Request) bool {
	var authResponse authenticationData
	authTokenString := r.Header.Get("Authorization")

	splitToken := strings.Split(authTokenString, "Bearer")
	if len(splitToken) != 2 {
		authResponse = authenticationData{
			Token: "Invalid Request",
		}
		//return 400
		sendResponse("Unauthorized", "Unauthorized", http.StatusUnauthorized, w)
		logger.Error.Println(authResponse)
		return false
	}
	jwtToken := strings.TrimSpace(splitToken[1])

	//obtain token from authorization header
	claim := &claims{}

	tkn, err := jwt.ParseWithClaims(jwtToken, claim, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			authResponse = authenticationData{
				Token: "Unauthroized Request",
			}
			//return 403
			sendResponse("Unauthorized", "Unauthorized", http.StatusUnauthorized, w)
			return false
		}
		authResponse = authenticationData{
			Token: "Invalid Request",
		}
		//return 400
		sendResponse("Unauthorized", "Unauthorized", http.StatusUnauthorized, w)
		return false
	}
	if !tkn.Valid {
		authResponse = authenticationData{
			Token: "Unauthroized Request",
		}
		//return 400
		sendResponse("Unauthorized", "Unauthorized", http.StatusUnauthorized, w)
		return false
	}
	return true
}
