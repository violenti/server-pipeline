package route

import (
	"fmt"
	"net/http"
)


func Health(w http.ResponseWriter, r *http.Request)  {

  fmt.Printf("ok")

}
