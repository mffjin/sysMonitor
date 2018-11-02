package profile

type CpuStatus struct {
    CPUUseage float64
    CPU []CpuInfo
}
type CpuInfo struct {
    ModelName string
    Cores int32
}

type MemStatus struct {
	Mem MemInfo
	Swap SwapInfo
}
type MemInfo struct {
    Total uint64
    Available uint64
    Used uint64
	Free uint64
    MemUseage float64
}
type SwapInfo struct {
    Total uint64
    Available uint64
    Used uint64
	Free uint64
    SwapUseage float64
	Sin uint64
	Sout uint64
}

type DiskStatus struct {
    Total uint64
    Free uint64
    DiskUsage float64
}

type HostInfo struct {
    Platform string
    PlatformFamily string
    PlatformVersion string
    Hostname string
    OS string
}

type LoadStatus struct {
    Load1 float64
    Load5 float64
    Load15 float64
}


type NetInfo struct {
    Addrs []string
    ByteSent uint64
    ByteRecv uint64
}

type Addr struct {
    IP   string
    Port uint32
}
type ConnectionStat struct {
    Family uint32
    Type   uint32
    Laddr  Addr
    Raddr  Addr
    Status string
}

type Item struct {
	Tag string
    Payload interface{}
}
