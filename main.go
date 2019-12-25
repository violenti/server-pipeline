package main

import (
    "io/ioutil"
    "log"
    "fmt"
    "net/http"
    "os"
)

var GITLAB_TOKEN = os.Getenv("TOKEN")
var URI = os.Getenv("GITLAB")
var FULL_URI= URI + GITLAB_TOKEN

func main ()  {

  http.HandleFunc("/master/default.yml", Master)

  http.ListenAndServe(":8080", nil)
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
