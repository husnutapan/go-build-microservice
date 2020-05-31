package handler

import (
	"github.com/husnutapan/go-build-microservice/utility"
	"net/http"
)

func (svr *ServerInformations) Home(w http.ResponseWriter, r *http.Request) {
	utility.ResponseData(w, "This is base endpoint...", success)
}
