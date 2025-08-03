package main

import (
	logging "github.com/AndrewI26/whiz/logger"
	"github.com/AndrewI26/whiz/router"
	"github.com/AndrewI26/whiz/server"
)

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

	router := router.NewRouter()

	server := server.NewServer(8000)
	server.AddRouter(router)
	server.Serve()
}
