package utility

import (
	"github.com/gorilla/mux"
	handler "github.com/husnutapan/go-build-microservice/handler"
	"net/http"
)

func (svr *ServerInformations) UpServer() {
	prepareRoutings(svr)
	http.ListenAndServe(":8080", svr.Router)
}

func prepareRoutings(svr *ServerInformations) {
	svr.Router = mux.NewRouter()
	svr.Router.HandleFunc("/", AddHeaderToJSON(handler.Home)).Methods("GET")
}
