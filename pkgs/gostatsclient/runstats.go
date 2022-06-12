package gostatsclient

import (
	"flag"
	"runtime"
	"time"
)

var (
	pause   int  = 5
	publish bool = true
)

func init() {
	go runCollector()
}

func runCollector() {
	for !flag.Parsed() {
		// Defer execution of this goroutine.
		runtime.Gosched()

		// Add an initial delay while the program initializes to avoid attempting to collect
		// metrics prior to our flags being available / parsed.
		time.Sleep(10 * time.Second)
	}

	c := New()
	c.PauseDur = time.Duration(pause) * time.Second

	if publish {
		c.Runner("localhost", 8000)
	}
}
