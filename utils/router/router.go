package router

import (
	"log"
	"net/http"
	"time"

	"github.com/alactic/mygoproject/routes/routerindex"
	"github.com/gorilla/mux"
)

func Router() {
	// r := mux.NewRouter()
	// v1 := r.PathPrefix("/api").Subrouter()
	// routerindex.Routerindex(v1)
	// if err := http.ListenAndServe(":8000", r); err != nil {
	// 	log.Fatal("Serving error.")
	// }

	r := mux.NewRouter()
	v1 := r.PathPrefix("/api").Subrouter()
	routerindex.Routerindex(v1)

	srv := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8900",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
