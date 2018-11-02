package load

import (
	_"fmt"
	_"encoding/json"
	"sysmonitor/common"
    "sysmonitor/profile"
	"github.com/shirou/gopsutil/load"
)

/*
type LoadStatus struct {
	Load1 float64	
	Load5 float64	
	Load15 float64
}
*/

func LoadMonitor() string {
	load_, _ := load.Avg()
	load_status := new(profile.LoadStatus)
	load_status.Load1  = load_.Load1
	load_status.Load5 = load_.Load5
	load_status.Load15 = load_.Load15
	//result := common.JsonMarshal(load_status)
    //return result

	load_result := new (profile.Item)
    load_result.Tag = "Load"
    load_result.Payload = load_status
    result := common.JsonMarshal(load_result)
    return result
}
