package main

import (
	"./handler"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	helloHandler := handler.NewHello(l)
	goodbyeHandler := handler.NewGoodbye(l)
	productHandler := handler.NewProduct(l)

	sm := http.NewServeMux()
	sm.Handle("/", helloHandler)
	sm.Handle("/goodbye", goodbyeHandler)
	sm.Handle("/product", productHandler)

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
