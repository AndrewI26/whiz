package main

import (
	"fmt"

	logging "github.com/AndrewI26/whiz/logger"
)

func main() {
	logger := logging.NewLogger(logging.Info)
	fmt.Printf("%d", logger.LogLevel)

}
