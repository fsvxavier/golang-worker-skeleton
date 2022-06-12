package gostatsclient

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/fsvxavier/golang-worker-skeleton/pkg/ip"
	"github.com/fsvxavier/golang-worker-skeleton/pkg/udp"
)

const maxPacketSize = 65536 - 8 - 20 // 8-byte UDP header, 20-byte IP header

// Collector implements the periodic grabbing of informational data from the
// runtime package and outputting the values to a GaugeFunc.
type GoStatsClient struct {
	// PauseDur represents the interval inbetween each set of stats output.
	// Defaults to 10 seconds.
	PauseDur time.Duration

	// Done, when closed, is used to signal Collector that is should stop collecting
	// statistics and the Run function should return. If Done is set, upon shutdown
	// all gauges will be sent a final zero value to reset their values to 0.
	Done <-chan struct{}
}

// New creates a new Collector that will periodically output statistics to gaugeFunc. It
// will also set the values of the exported fields to the described defaults. The values
// of the exported defaults can be changed at any point before Run is called.
func New() *GoStatsClient {
	return &GoStatsClient{
		PauseDur: 10 * time.Second,
	}
}

// Run gathers statistics from package runtime and outputs them to the configured GaugeFunc every
// PauseDur. This function will not return until Done has been closed (or never if Done is nil),
// therefore it should be called in its own goroutine.
func (c *GoStatsClient) Run() {
	defer c.zeroStats()
	c.outputStats()

	// Gauges are a 'snapshot' rather than a histogram. Pausing for some interval
	// aims to get a 'recent' snapshot out before statsd flushes metrics.
	tick := time.NewTicker(c.PauseDur)
	defer tick.Stop()
	for {
		select {
		case <-c.Done:
			return
		case <-tick.C:
			c.outputStats()
			c.zeroStats()
		}
	}
}

func (c *GoStatsClient) Runner(hostServer string, portServer int64) {
	client := udp.NewClient()

	addressUdp := hostServer + ":" + strconv.FormatInt(portServer, 10)

	client.SetupConnection(addressUdp)

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	go client.ReadFromSocket(maxPacketSize)
	go client.ProcessPackets()
	go client.ProcessMessages()

	for {
		select {
		case <-ticker.C:
			client.Send(c.outputStats())
		}
	}
}

// zeroStats sets all the stat guages to zero. On shutdown we want to zero them out so they don't persist
// at their last value until we start back up.
func (c *GoStatsClient) zeroStats() {
	output := CollectorEntity{}
	c.defaultData(&output)
	c.outputCPUStats(&output)

	mStats := runtime.MemStats{}
	c.outputMemStats(&mStats, &output)
	c.outputGCStats(&mStats, &output)
}

func (c *GoStatsClient) defaultData(ce *CollectorEntity) {

	ce.OsSystem = runtime.GOOS
	ce.OsSystemArch = runtime.GOARCH
	ce.App = os.Getenv("APP_NAME")
	ce.Version = os.Getenv("APP_VERSION")

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}
	ce.Hostname = hostname

	ce.CurrentTime = string(time.Now().UTC().Format(time.RFC3339))

	localIp, err := ip.InternalIP()
	if err != nil {
		fmt.Println(err)
	}
	ce.LocalIp = localIp
}

func (c *GoStatsClient) outputStats() string {

	output := CollectorEntity{}

	c.defaultData(&output)

	// Colect data CPU
	c.outputCPUStats(&output)

	// Colect data Memory
	m := &runtime.MemStats{}
	runtime.ReadMemStats(m)
	c.outputMemStats(m, &output)

	// Colect data Garbage Collection
	c.outputGCStats(m, &output)

	outputString, err := json.Marshal(output)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return ""
	}

	return string(outputString)
}

func (c *GoStatsClient) outputCPUStats(ce *CollectorEntity) {

	ce.CPU.Goroutines = uint64(runtime.NumGoroutine())
	ce.CPU.CgoCalls = uint64(runtime.NumCgoCall())
}

func (c *GoStatsClient) outputMemStats(m *runtime.MemStats, ce *CollectorEntity) {
	// General
	ce.Mem.Alloc = m.Alloc
	ce.Mem.Total = m.TotalAlloc
	ce.Mem.Sys = m.Sys
	ce.Mem.Lookups = m.Lookups
	ce.Mem.Malloc = m.Mallocs
	ce.Mem.Frees = m.Frees
	ce.Mem.Othersys = m.OtherSys

	ce.Mem.PauseNs = m.PauseNs
	ce.Mem.PauseTotalNs = m.PauseTotalNs
	ce.Mem.GCCPUFraction = m.GCCPUFraction

	// Heap
	ce.Mem.Heap.Alloc = m.HeapAlloc
	ce.Mem.Heap.Sys = m.HeapSys
	ce.Mem.Heap.Idle = m.HeapIdle
	ce.Mem.Heap.Inuse = m.HeapInuse
	ce.Mem.Heap.Released = m.HeapReleased
	ce.Mem.Heap.Objects = m.HeapObjects

	// Stack
	ce.Mem.Stack.Inuse = m.StackInuse
	ce.Mem.Stack.Sys = m.StackSys
	ce.Mem.Stack.MspanInuse = m.MSpanInuse
	ce.Mem.Stack.MspanSys = m.MSpanSys
	ce.Mem.Stack.McacheInuse = m.MCacheInuse
	ce.Mem.Stack.McacheSys = m.MCacheSys

}

func (c *GoStatsClient) outputGCStats(m *runtime.MemStats, ce *CollectorEntity) {

	// Garbage Collection
	ce.Mem.Gc.Sys = m.GCSys
	ce.Mem.Gc.Next = m.NextGC
	ce.Mem.Gc.Last = m.LastGC
	ce.Mem.Gc.PauseTotal = m.PauseTotalNs
	ce.Mem.Gc.Pause = m.PauseNs[(m.NumGC+255)%256]
	ce.Mem.Gc.Count = uint64(m.NumGC)
	ce.Mem.Gc.NumForced = uint64(m.NumForcedGC)
}
