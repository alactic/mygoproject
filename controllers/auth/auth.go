package auth

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"

	"github.com/alactic/mygoproject/models/auth"
	// "github.com/alactic/mygoproject/models/staff"
	"github.com/alactic/mygoproject/models/customers"
	"github.com/alactic/mygoproject/utils/connection"
	hashed "github.com/alactic/mygoproject/utils/hash"
	jwtFile "github.com/alactic/mygoproject/utils/jwt"
	"gopkg.in/couchbase/gocb.v1"
)

type Customer = customers.Customer

var bucket *gocb.Bucket = connection.Connection()

func LoginEndpoint(response http.ResponseWriter, request *http.Request) {
	// jwtFile.DecodeJWT(request.Header["Authorization"][0])
	response.Header().Set("Content-Type", "application/json")
	routerParams := mux.Vars(request)
	var customer Customer
	var auth auth.Auth
	_ = json.NewDecoder(request.Body).Decode(&auth)
	hashed.Hash(auth.Password)
	hashed.CompareHashValue("password", "hashpassword")
	_, err := bucket.Get(routerParams["id"], &customer)
	if err != nil {
		if err.Error() == "key not found" {
			response.Write([]byte(`{"data": "{}"}`))
			return
		}
		response.WriteHeader(500)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(customer)
	// details := make(map[string]string)
	// details["firstname"] = "echezona"
	// details["lastname"] = "okafor"
	// details["email"] = "okaforechezona1992@gmail.com"
	// var token = jwtFile.GenerateJWT(details)
	// json.NewEncoder(response).Encode(token)
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
