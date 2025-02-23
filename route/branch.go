package route

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func Branch(w http.ResponseWriter, r *http.Request) {

	var GITLAB_TOKEN = os.Getenv("GITLAB_TOKEN")
	var URI = os.Getenv("GITLAB")
	var varrequest = mux.Vars(r)
	var FILE = varrequest["file"]
	var BRANCH = varrequest["branch"]
	var REF = "/raw?ref=" + BRANCH
	var FULL_URI = URI + FILE + REF +  "&private_token=" + GITLAB_TOKEN

	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.JSONFormatter{})


	resp, err := http.Get(FULL_URI)
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
		fmt.Fprint(w, string(contents))
		log.Printf(FILE)

	}
}
