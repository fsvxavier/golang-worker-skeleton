package udp

import (
	"fmt"

	"github.com/fsvxavier/golang-worker-skeleton/pkg/logger"
)

func ErrorCheck(err error, where string, kill bool) {
	logger := logger.NewGenericLogger()
	if err != nil {
		if kill {
			logger.LogIt("FATAL", "Script Terminated...", nil)
		} else {
			logger.LogIt("WARN", fmt.Sprintf("@ %s\n", where), nil)
		}
	}
}
