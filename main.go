package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "My go programming App with mux")
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
}

func main() {
	r := mux.NewRouter()
	fmt.Println("Go Web App Started on Port 3000")
	setupRoutes()
	http.ListenAndServe(":5000", r)
}
