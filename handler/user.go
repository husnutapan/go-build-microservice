package handler

import (
	"encoding/json"
	"github.com/husnutapan/go-build-microservice/pojo"
	"github.com/husnutapan/go-build-microservice/service"
	"io/ioutil"
	"net/http"
)

func SaveUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	user := pojo.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		return
	}
	result, err := service.User{}.SaveUser(&user)
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(result)
}
