package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/couchbase/gocb.v1"
)

var bucket *gocb.Bucket

func main() {
	fmt.Println("Starting application ...")

	// bucket = connection.Connection()
	// router.Router()

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// an example API handler
		fmt.Fprintf(w, "My go application Application 22")
	})
	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8800",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
