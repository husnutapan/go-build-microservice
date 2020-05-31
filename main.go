package main

import "github.com/husnutapan/go-build-microservice/handler"

func main() {
	var server = handler.ServerInformations{}
	server.UpServer()
}
