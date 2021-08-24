package route

import (
  "os"
	"net/http"
  log "github.com/sirupsen/logrus"
)


func Health(w http.ResponseWriter, r *http.Request)  {

  log.SetOutput(os.Stdout)
	log.SetFormatter(&log.JSONFormatter{})
  
  log.Printf("ok")

}
