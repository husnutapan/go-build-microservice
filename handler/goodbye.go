package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (h *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello Goodbye")
	d, err := ioutil.ReadAll(r.Body)
	for err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
	}
	fmt.Fprintf(rw, "Hello goodbye %s", d)
}
