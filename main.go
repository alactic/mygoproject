package main

import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "My go programming App")
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
}

func main() {
	fmt.Println("Go Web App Started on Port 3000")
	setupRoutes()
	http.ListenAndServe(":4000", nil)
}
