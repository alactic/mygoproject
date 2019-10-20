package customers

import (
	"net/http"

	"github.com/alactic/mygoproject/controllers/customers"
	"github.com/alactic/mygoproject/utils/router"
)

func Customers() {
	router := router.Router()
	router.HandleFunc("/customer", customers.CreateCustomerEndpoint).Methods("POST")
	router.HandleFunc("/customer/{id}", customers.GetCustomerEndpoint).Methods("GET")
	router.HandleFunc("/customers", customers.GetCustomersEndpoint).Methods("GET")
	http.Handle("/", router)
}
