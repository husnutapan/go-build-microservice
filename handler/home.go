package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (svr *ServerInformations) Home(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode("Welcome To This Awesome API")
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}
