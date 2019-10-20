package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alactic/mygoproject/routes/customers"
	"github.com/alactic/mygoproject/utils/connection"
	"github.com/gorilla/mux"
	"gopkg.in/couchbase/gocb.v1"
)

var bucket *gocb.Bucket

func main() {
	fmt.Println("Starting application ...")

	bucket = connection.Connection()
	router := mux.NewRouter()
	// routerindex.Routerindex()
	customers.Customers()

	// router.HandleFunc("/customer", customers.CreateCustomerEndpoint).Methods("POST")
	// router.HandleFunc("/customer/{id}", customers.GetCustomerEndpoint).Methods("GET")
	// router.HandleFunc("/customer/creditcard/{id}", creditcards.AddCreditCardEndpoint).Methods("PUT")
	// router.HandleFunc("/customer/creditcard/{id}", creditcards.AddCreditCardEndpoint).Methods("GET")
	// router.HandleFunc("/customers", customers.GetCustomersEndpoint).Methods("GET")

	// router.HandleFunc("/product", products.CreateProductEndpoint).Methods("POST")
	// router.HandleFunc("/product/{id}", products.GetProductEndpoint).Methods("GET")
	// router.HandleFunc("/products", products.GetProductsEndpoint).Methods("GET")
	// router.HandleFunc("/receipt", receipts.CreateReceiptEndpoint).Methods("POST")
	// router.HandleFunc("/receipt/{id}", receipts.GetReceiptEndpoint).Methods("GET")
	// http.Handle("/", customers.Customers())
	fmt.Println("Listening at :12345")
	log.Fatal(http.ListenAndServe(":12345", router))
}
