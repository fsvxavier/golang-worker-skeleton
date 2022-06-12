package main

import (
	"flag"
	"log"
	"os"
	"runtime"

	"github.com/fsvxavier/golang-worker-skeleton/cmd"
	"github.com/fsvxavier/golang-worker-skeleton/pkg/enviroment"
	_ "github.com/fsvxavier/golang-worker-skeleton/pkg/gostatsclient"
	"github.com/fsvxavier/golang-worker-skeleton/pkg/logger"
)

func main() {
	flag.Parse()

	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)

	env := enviroment.NewEnviroment()
	os.Setenv("ENV", "production")
	env.SetFileConfig("./config/env.json")

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	logger := logger.NewGenericLogger()

	w := cmd.Worker{
		Logger: logger,
	}
	w.StartWorker()

}
