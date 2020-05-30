package handler

import (
	"github.com/husnutapan/go-build-microservice/utility"
	"net/http"
)

func (svr *ServerInformations) UpServer() {
	http.ListenAndServe(":8080", svr.Router)
}

func (s *ServerInformations) prepareRoutings() {
	s.Router.HandleFunc("/", utility.AddHeaderToJSON(s.Home))

}
