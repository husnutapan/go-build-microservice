package handler

import (
	"github.com/gorilla/mux"
	"github.com/husnutapan/go-build-microservice/utility"
	"net/http"
)

func (svr *ServerInformations) UpServer() {
	prepareRoutings(svr)
	http.ListenAndServe(":8080", svr.Router)
}

func prepareRoutings(svr *ServerInformations) {
	svr.Router = mux.NewRouter()
	svr.Router.HandleFunc("/", utility.AddHeaderToJSON(svr.Home)).Methods("GET")

}
