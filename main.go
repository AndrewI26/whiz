package main

import (
	logging "github.com/AndrewI26/whiz/logger"
)

func main() {
	logger := logging.NewLogger(
		logging.Info,
		"LOG_",
		logging.RollingConfig{
			TimeThreshold: logging.Minutely,
			SizeThreshold: logging.HalfMB,
		})

	logger.Info("Hello")

	err := logger.Open()
	if err != nil {
		panic(err)
	}
	defer logger.Close()

	logger.Info("Hello first log so cool")
}
