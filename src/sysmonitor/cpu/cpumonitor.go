package cpu

import (
	_"fmt"
	"time"
	_"encoding/json"
	"sysmonitor/profile"
	"sysmonitor/common"
	"github.com/shirou/gopsutil/cpu"
)

func CpuMonitor() string {
	cpu_info, _ := cpu.Info()
	cpu_percent, _ := cpu.Percent(time.Second, false)

	cpustatus := new(profile.CpuStatus)
	cpustatus.CPU = make([]profile.CpuInfo, len(cpu_info))

	for i, ci := range cpu_info {
        	cpustatus.CPU[i].ModelName = ci.ModelName
        	cpustatus.CPU[i].Cores = ci.Cores
        }
	cpustatus.CPUUseage =  cpu_percent[0]

	cpu_result := new (profile.Item)
	cpu_result.Tag = "CPU"
    cpu_result.Payload = cpustatus
	result := common.JsonMarshal(cpu_result)
	return result
}
