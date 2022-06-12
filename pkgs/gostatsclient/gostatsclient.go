package gostatsclient

type CollectorEntity struct {
	CPU          CPUEntity `json:"cpu"`
	Mem          MemEntity `json:"mem"`
	OsSystem     string    `json:"os_system"`
	OsSystemArch string    `json:"os_system_arch"`
	App          string    `json:"app"`
	Version      string    `json:"version"`
	Hostname     string    `json:"hostname"`
	LocalIp      string    `json:"local_ip"`
	CurrentTime  string    `json:"current_time"`
}

type CPUEntity struct {
	Goroutines uint64 `json:"goroutines"`
	CgoCalls   uint64 `json:"cgo_calls"`
}

type HeapEntity struct {
	Alloc    uint64 `json:"alloc"`
	Sys      uint64 `json:"sys"`
	Idle     uint64 `json:"idle"`
	Inuse    uint64 `json:"inuse"`
	Released uint64 `json:"released"`
	Objects  uint64 `json:"objects"`
}

type StackEntity struct {
	Inuse       uint64 `json:"inuse"`
	Sys         uint64 `json:"sys"`
	MspanInuse  uint64 `json:"mspan_inuse"`
	MspanSys    uint64 `json:"mspan_sys"`
	McacheInuse uint64 `json:"mcache_inuse"`
	McacheSys   uint64 `json:"mcache_sys"`
}

type GcEntity struct {
	Sys        uint64 `json:"sys"`
	Next       uint64 `json:"next"`
	Last       uint64 `json:"last"`
	PauseTotal uint64 `json:"pause_total"`
	Pause      uint64 `json:"pause"`
	Count      uint64 `json:"count"`
	NumForced  uint64 `json:"num_forced"`
}

type MemEntity struct {
	Alloc         uint64      `json:"alloc"`
	Total         uint64      `json:"total"`
	Sys           uint64      `json:"sys"`
	Lookups       uint64      `json:"lookups"`
	Malloc        uint64      `json:"malloc"`
	Frees         uint64      `json:"frees"`
	Othersys      uint64      `json:"othersys"`
	PauseNs       [256]uint64 `json:"pausens"`
	PauseTotalNs  uint64      `json:"pausetotalns"`
	GCCPUFraction float64     `json:"gccpufraction"`
	Heap          HeapEntity  `json:"heap"`
	Stack         StackEntity `json:"stack"`
	Gc            GcEntity    `json:"gc"`
}
