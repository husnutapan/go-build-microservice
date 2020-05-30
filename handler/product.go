package handler

import (
	"encoding/json"
	"github.com/husnutapan/go-build-microservice/utility"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type Product struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

func (p *Product) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		p.getProductList(w, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(w, r)
		return
	}

	if r.Method == http.MethodPut {
		regex := regexp.MustCompile(`/([0-9]+)`)
		g := regex.FindAllStringSubmatch(r.URL.Path, -1)
		id, _ := strconv.Atoi(g[0][1])
		p.updateProduct(id, w, r)
		return
	}

}

func (p *Product) getProductList(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handled GET product Method")
	products := utility.GetProductList()
	data, err := json.Marshal(products)
	if err != nil {
		http.Error(w, "Unable marshal json", http.StatusInternalServerError)
	}
	w.Write(data)
}

func (p *Product) addProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handled POST product Method")
	newProduct := &utility.Product{}
	err := newProduct.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable marshal json", http.StatusInternalServerError)
	}
	utility.AddProduct(newProduct)
}

func (p *Product) updateProduct(id int, w http.ResponseWriter, r *http.Request) {
	newProduct := &utility.Product{}
	err := newProduct.FromJSON(r.Body)

	if err != nil {
		http.Error(w, "Unable marshal json", http.StatusInternalServerError)
	}

	utility.UpdateProduct(id, newProduct)
}
