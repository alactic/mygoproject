package customers

import (
	"encoding/json"
	"net/http"

	"github.com/alactic/mygoproject/models/customers"
	"github.com/alactic/mygoproject/utils/connection"
	"github.com/gorilla/mux"

	uuid "github.com/satori/go.uuid"
	"gopkg.in/couchbase/gocb.v1"
)

type Customer = customers.Customer

var bucket *gocb.Bucket = connection.Connection()

//router.HandleFunc("/customer", CreateCustomerEndpoint).Methods("POST")
func CreateCustomerEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var customer Customer
	_ = json.NewDecoder(request.Body).Decode(&customer)
	id := uuid.Must(uuid.NewV4()).String()
	customer.Type = "customer"
	customer.Id = id
	_, err := bucket.Insert(id, customer, 0)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(`{"message": "` + err.Error() + `" }`))
		return
	}
	response.Write([]byte(`{ "id": "` + id + `"}`))
}

//router.HandleFunc("/customer/{id}", GetCustomerEndpoint).Methods("GET")
func GetCustomerEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	routerParams := mux.Vars(request)
	var customer Customer
	customer.Id = routerParams["id"]
	_, err := bucket.Get(routerParams["id"], &customer)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(customer)
}

//router.HandleFunc("/customers", GetCustomersEndpoint).Methods("GET")
// //router.HandleFunc("/customers", GetCustomersEndpoint).Methods("GET")
func GetCustomersEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	var customers []Customer
	query := gocb.NewN1qlQuery("SELECT META().id, " + bucket.Name() + ".* FROM " + bucket.Name() + " WHERE type = 'customer'")
	rows, err := bucket.ExecuteN1qlQuery(query, nil)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	var row Customer
	for rows.Next(&row) {
		customers = append(customers, row)
	}
	json.NewEncoder(response).Encode(customers)
}
