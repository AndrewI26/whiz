package main

import (
	"fmt"
	"strconv"
	"strings"

	logging "github.com/AndrewI26/whiz/logger"
	"github.com/AndrewI26/whiz/routing"
	"github.com/AndrewI26/whiz/server"
)

func main() {
	logger := logging.NewLogger(
		logging.Info,
		"LOG_",
		logging.OneKB,
	)

	err := logger.Open()
	if err != nil {
		panic(err)
	}
	defer logger.Close()

	router := routing.NewRouter(logger)
	router.Get("/hello/:name", func(params map[string]string) *routing.Response {
		// Dynamic parameters can be accessed through the params map
		name := params["name"]
		greeting := fmt.Sprintf("<h1>Hello there, %s! Nice to meet you.</h1>", name)

		res := routing.Response{
			Status:  200,
			Data:    greeting,
			Headers: map[string]string{"Content-Type": "text/html; charset=utf-8"},
		}
		res.Headers["Content-Length"] = strconv.Itoa(len(res.Data))

		return &res
	})
	router.Get("/hellos/:numhellos/:name", func(params map[string]string) *routing.Response {
		name := params["name"]
		numHellosParam := params["numhellos"]
		numHellos, err := strconv.Atoi(numHellosParam)
		if err != nil {
			logger.Error(fmt.Sprintf("Could not convert string to int (%s)", numHellosParam))
		}

		hellos := strings.Repeat(fmt.Sprintf("<h3>Hello %s!<h3>", name), numHellos)
		res := routing.Response{
			Status:  200,
			Data:    hellos,
			Headers: map[string]string{"Content-Type": "text/html; charset=utf-8"},
		}
		res.Headers["Content-Length"] = strconv.Itoa(len(res.Data))
		return &res
	})

	server := server.NewServer(logger, 8000)
	server.AddRouter(router)
	server.Serve()
}
