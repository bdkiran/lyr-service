package api

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/bdkiran/lyr-service/utils"
	"github.com/dgrijalva/jwt-go"
)

var jwtKey []byte

type authenticationData struct {
	UserID string `json:"user_id,omitempty"`
	Token  string `json:"token,omitempty"`
}

func tokenInit() {
	key, tokenErr := utils.GetEnvVariableString("TOKEN_SECRET")
	if tokenErr != nil {
		logger.Error.Fatal(tokenErr)
		return
	}
	jwtKey = []byte(key)
}

func validateToken(w http.ResponseWriter, r *http.Request) (bool, int) {
	var authenticatedUser int
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
		return false, authenticatedUser
	}
	jwtToken := strings.TrimSpace(splitToken[1])

	//obtain token from authorization header
	claim := jwt.MapClaims{}

	tkn, err := jwt.ParseWithClaims(jwtToken, claim, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			authResponse = authenticationData{
				Token: "Unauthroized Request",
			}
			//return 403
			sendResponse("Unauthorized", "invalid jwt singnature", http.StatusUnauthorized, w)
			return false, authenticatedUser
		}
		authResponse = authenticationData{
			Token: "Invalid Request",
		}
		//return 400
		sendResponse("Unauthorized", "failed to authenticate", http.StatusUnauthorized, w)
		return false, authenticatedUser
	}
	if !tkn.Valid {
		authResponse = authenticationData{
			Token: "Unauthroized Request",
		}
		//return 400
		sendResponse("Unauthorized", "invalid token", http.StatusUnauthorized, w)
		return false, authenticatedUser
	}
	authenticatedUser, err = strconv.Atoi(claim["sub"].(string))
	if err != nil {
		logger.Warning.Printf("Unable to convert %s to int", claim["sub"])
		sendResponse("Error", "Error authorizing user", http.StatusUnauthorized, w)
		return false, authenticatedUser
	}
	return true, authenticatedUser
}
