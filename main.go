package main

import (
	"fmt"
	"net/http"
	"time"
  "github.com/violenti/server-pipeline/route"
	"github.com/gorilla/mux"
)

//var GITLAB_TOKEN = os.Getenv("GITLAB_TOKEN")
//var URI = os.Getenv("GITLAB")

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{branch}/{file}", route.Branch).Methods("GET")
	http.Handle("/", r)

	srv := &http.Server{
		Addr: "0.0.0.0:8000",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}
	fmt.Println("server-pipeline running on http://0.0.0.0:8000 ")
	srv.ListenAndServe()
}
