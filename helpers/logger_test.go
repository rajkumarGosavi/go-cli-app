package helpers

import (
	"log"
	"testing"
)

func TestGetLoggerInstace(t *testing.T) {
	var loggerInstance *log.Logger
	if loggerInstance != nil {
		t.Error("logger Instance should be nil")
	}
	loggerInstance = GetLoggerInstace()
	if loggerInstance == nil {
		t.Error("logger Instance should not be nil")
	}
}
