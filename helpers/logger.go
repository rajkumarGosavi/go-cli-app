package helpers

import (
	"log"
	"os"
)

var logger *log.Logger = GetLoggerInstace()

// GetLoggerInstace - Returns basic logger instance
func GetLoggerInstace() *log.Logger {

	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile("error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Panicln("Unable to open log file:", err)
	}
	return log.New(f, "[mycart]: ", log.Lshortfile)
}
