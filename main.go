package main

import (
    "io/ioutil"
    "time"
    "log"
    "fmt"
    "net/http"
    "os"
    "github.com/gorilla/mux"
)

var GITLAB_TOKEN = os.Getenv("TOKEN")
var URI = os.Getenv("GITLAB")
var FULL_URI= URI + GITLAB_TOKEN

func main ()  {
  r := mux.NewRouter()
  r.HandleFunc("/master/default.yml", Master)

  srv := &http.Server{
        Addr:         "0.0.0.0:8080",
        // Good practice to set timeouts to avoid Slowloris attacks.
        WriteTimeout: time.Second * 15,
        ReadTimeout:  time.Second * 15,
        IdleTimeout:  time.Second * 60,
        Handler: r, // Pass our instance of gorilla/mux in.
    }
    log.Fatal(srv.ListenAndServe())
}

func Master (w http.ResponseWriter, r *http.Request)  {
  resp, err:= http.Get(FULL_URI)
  if err != nil {
    log.Fatalln(err)
    fmt.Printf("%s", err)
    os.Exit(1)
  } else {
    defer resp.Body.Close()
    contents, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln(err)
        os.Exit(1)
    }
    fmt.Fprint(w,string(contents))
    fmt.Print(FULL_URI)

  }
}
