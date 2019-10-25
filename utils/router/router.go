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
