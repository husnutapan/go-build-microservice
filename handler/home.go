package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	err := json.NewEncoder(w).Encode("This is base endpoint...")
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}
