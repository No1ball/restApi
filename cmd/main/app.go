package main

import (
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"restApi/internal/user"
	"restApi/pkg/logging"
	"time"

	"fmt" // to delete
)

func main() {
	logger := logging.GetLogger()
	logger.Info("create router")
	router := httprouter.New() //add httprouter from github

	//create handler
	handler := user.NewHandler(logger)
	handler.Register(router)

	start(router)

	fmt.Print("hello")
}

func start(router *httprouter.Router) {
	logger := logging.GetLogger()
	logger.Info("Start ")
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	logger.Info("server is listening port 1234")
	logger.Fatalln(server.Serve(listener))
}
