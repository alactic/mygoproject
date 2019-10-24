package auth

import (
	"encoding/json"
	"net/http"

	"github.com/alactic/mygoproject/models/auth"
	hashed "github.com/alactic/mygoproject/utils/hash"
	jwtFile "github.com/alactic/mygoproject/utils/jwt"
)

func LoginEndpoint(response http.ResponseWriter, request *http.Request) {
	// jwtFile.DecodeJWT(request.Header["Authorization"][0])
	response.Header().Set("Content-Type", "application/json")
	var auth auth.Auth
	_ = json.NewDecoder(request.Body).Decode(&auth)
	hashed.Hash(auth.Password)
	hashed.CompareHashValue("password", "hashpassword")
	details := make(map[string]string)
	details["firstname"] = "echezona"
	details["lastname"] = "okafor"
	details["email"] = "okaforechezona1992@gmail.com"
	var token = jwtFile.GenerateJWT(details)
	json.NewEncoder(response).Encode(token)
}

func SignupEndpoint(response http.ResponseWriter, request *http.Request) {
	// jwtFile.DecodeJWT(request.Header["Authorization"][0])
	response.Header().Set("Content-Type", "application/json")
	var auth auth.Auth
	_ = json.NewDecoder(request.Body).Decode(&auth)
	hashed.Hash(auth.Password)
	details := make(map[string]string)
	details["firstname"] = "echezona"
	details["lastname"] = "okafor"
	details["email"] = "okaforechezona1992@gmail.com"
	var token = jwtFile.GenerateJWT(details)
	details["token"] = token
	json.NewEncoder(response).Encode(details)
}
