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

var GITLAB_TOKEN= os.Getenv("GITLAB_TOKEN")
var URI = os.Getenv("GITLAB")

func main ()  {
  r := mux.NewRouter()
  r.HandleFunc("/master/{file}", Master)
  r.HandleFunc("/development/{file}", Development)
  http.Handle("/", r)

  srv := &http.Server{
        Addr: "0.0.0.0:8000",
        // Good practice to set timeouts to avoid Slowloris attacks.
        WriteTimeout: time.Second * 15,
        ReadTimeout:  time.Second * 15,
        IdleTimeout:  time.Second * 60,
        Handler: r, // Pass our instance of gorilla/mux in.
    }
    fmt.Println("server-pipeline running on http://0.0.0.0:8000 ")
    log.Fatal(srv.ListenAndServe() )
}


func Master (w http.ResponseWriter, r *http.Request)  {

  var varrequest = mux.Vars(r)
  var FILE = varrequest["file"]
  var FULL_URI= URI + FILE + "/raw?ref=master" + "&private_token=" + GITLAB_TOKEN

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

func Development (w http.ResponseWriter, r *http.Request)  {


  var varrequest = mux.Vars(r)
  var FILE = varrequest["file"]
  var FULL_URI= URI + FILE + "/raw?ref=develop" + "&private_token=" + GITLAB_TOKEN

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
