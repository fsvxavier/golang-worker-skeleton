package cmd

import (
	"fmt"

	loggerInterfaces "github.com/fsvxavier/golang-worker-skeleton/pkg/logger/interfaces"
)

// Worker is the main structure of the application
type Worker struct {
	Logger loggerInterfaces.GenericLogger
}

func (w *Worker) initialization() {
	w.Logger.SetModule("worker")
	w.Logger.SetOperation("InitWorker")

	w.Logger.LogIt("INFO", fmt.Sprintf("Starting worker..."), nil)

}

// StartWorker is responsible for starting the execution of the process
func (w *Worker) StartWorker() {
	w.initialization()

}
