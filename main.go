package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go-build-microservice/handler"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	productHandler := handler.NewProduct(l)

	sm := mux.NewRouter()
	subRouterGet := sm.Methods(http.MethodGet).Subrouter()
	subRouterGet.HandleFunc("/", productHandler.GetProductList)

	subRouterPut := sm.Methods(http.MethodPut).Subrouter()
	subRouterPut.HandleFunc("/{id:[0-9]+}", productHandler.UpdateProduct)
	subRouterPut.Use(productHandler.MiddlewareValidateProduct)

	subRouterPost := sm.Methods(http.MethodPost).Subrouter()
	subRouterPost.HandleFunc("/", productHandler.AddProduct)
	subRouterPost.Use(productHandler.MiddlewareValidateProduct)
	server := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  time.Second,
		WriteTimeout: time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)
	<-signalChannel
	fmt.Println("Terminated server...")
	tc, _ := context.WithTimeout(context.Background(), time.Second*30)
	server.Shutdown(tc)
}
