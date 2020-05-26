package handler

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"go-build-microservice/utility"
	"log"
	"net/http"
	"strconv"
)

type Product struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

func (p *Product) GetProductList(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handled GET product Method")
	products := utility.GetProductList()
	data, err := json.Marshal(products)
	if err != nil {
		http.Error(w, "Unable marshal json", http.StatusInternalServerError)
	}
	w.Write(data)
}

func (p *Product) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST product Method")
	product := r.Context().Value(KeyProduct{}).(utility.Product)
	utility.AddProduct(&product)
}

func (p *Product) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable convert int to str", http.StatusBadGateway)
	}
	newProduct := r.Context().Value(KeyProduct{}).(utility.Product)

	utility.UpdateProduct(id, &newProduct)
}

type KeyProduct struct {
}

func (p Product) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		product := utility.Product{}
		err := product.FromJSON(r.Body)
		if err != nil {
			http.Error(rw, "Error reading product", http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(r.Context(), KeyProduct{}, product)
		r = r.WithContext(ctx)
		next.ServeHTTP(rw, r)
	})
}
