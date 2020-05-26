package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	d, err := ioutil.ReadAll(r.Body)
	for err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
	}
	fmt.Fprintf(rw, "Hello %s %s", d, time.Now())
}
