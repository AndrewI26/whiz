package main

import (
	"strconv"

	logging "github.com/AndrewI26/whiz/logger"
	"github.com/AndrewI26/whiz/routing"
	"github.com/AndrewI26/whiz/server"
)

func defaultHandler(params map[string]string) *routing.Response {
	res := routing.Response{200, "Hello", map[string]string{"My-Header": "Hello world"}}
	res.Headers["Content-Length"] = strconv.Itoa(len(res.Data))
	return &res
}

func hello(params map[string]string) *routing.Response {
	res := routing.Response{200, "Hello world", map[string]string{"My-Header": "Hello world"}}
	res.Headers["Content-Length"] = strconv.Itoa(len(res.Data))
	return &res
}

func main() {
	logger := logging.NewLogger(
		logging.Info,
		"LOG_",
		logging.OneKB,
	)

	logger.Info("Hello")

	err := logger.Open()
	if err != nil {
		panic(err)
	}
	defer logger.Close()

	router := routing.NewRouter()
	router.Get("/user/rentals", defaultHandler)
	router.Get("/user/:user", defaultHandler)
	router.PrintRoutes(router.Root, 1)

	server := server.NewServer(8000)
	server.AddRouter(router)
	server.Serve()
}
