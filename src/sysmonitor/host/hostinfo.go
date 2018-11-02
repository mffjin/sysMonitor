package host

import (
	_"fmt"
	_"encoding/json"
	"sysmonitor/common"
	"sysmonitor/profile"
	"github.com/shirou/gopsutil/host"	
)

/*
type HostInfo struct {
	Platform string	
	PlatformFamily string 
	PlatformVersion string
	Hostname string
	OS string
}
*/

func GetSysInfo() string {
	host_info, _ := host.Info()

	host_status := new(profile.HostInfo)

	host_status.Platform = host_info.Platform
	host_status.PlatformFamily = host_info.PlatformFamily
	host_status.PlatformVersion = host_info.PlatformVersion
	host_status.Hostname = host_info.Hostname
	host_status.OS = host_info.OS

	//result := common.JsonMarshal(host_status)
    //fmt.Printf("        Host        : %v  \n", string(result))
    //return result

	host_result := new (profile.Item)
    host_result.Tag = "Host"
    host_result.Payload = host_status
    result := common.JsonMarshal(host_result)
    return result
}
