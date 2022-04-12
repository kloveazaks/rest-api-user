package main

import (
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"rest-api-tutorial/internal/user"
	"rest-api-tutorial/pgk/logging"
	"time"
)

//func IndexHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
//	name := params.ByName("name")
//	w.Write([]byte(fmt.Sprintf("Hello %s", name)))
//}

func main() {
	logger := logging.GetLogger()
	logger.Info("create router")
	router := httprouter.New()

	logger.Info("Register user handler")
	handler := user.NewHandler(*logger)
	handler.Register(router)

	start(router)
}

func start(router *httprouter.Router) {
	logger := logging.GetLogger()
	logger.Info("Start application")
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Info("Server is listening port 0.0.0.0:1234")
	logger.Fatal(server.Serve(listener))
}
