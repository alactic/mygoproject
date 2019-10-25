package router

import (
	"log"
	"net/http"

	"github.com/alactic/mygoproject/routes/routerindex"
	"github.com/gorilla/mux"
)

func Router() {
	r := mux.NewRouter()
	v1 := r.PathPrefix("/api").Subrouter()
	routerindex.Routerindex(v1)

	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal("Serving error.")
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "My Awesome Go App")
  }
  
  func setupRoutes() {
	http.HandleFunc("/", homePage)
  }
  
  func main() {
	fmt.Println("Go Web App Started on Port 3000")
	setupRoutes()
	http.ListenAndServe(":3091", nil)
  }