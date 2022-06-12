package cmd

import (
	"testing"

	"github.com/fsvxavier/golang-worker-skeleton/pkg/logger"
)

func TestInitializeRabbitMQError(t *testing.T) {

	logger := logger.NewGenericLogger()
	logger.SetModule("worker")
	logger.SetOperation("InitWorker")
	worker := Worker{
		Logger: logger,
	}

	worker.initialization()
}
