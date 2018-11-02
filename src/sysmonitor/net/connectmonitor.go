package net

import (
    _"fmt"
    _"encoding/json"
    "sysmonitor/common"
    "sysmonitor/profile"
    "github.com/shirou/gopsutil/net"
)

func ConnectMonitor() string {
	net_connect, _ := net.Connections("tcp")
	net_result := new (profile.Item)
    net_result.Tag = "Net_connect"
    net_result.Payload = net_connect 
    result := common.JsonMarshal(net_result)
    return result

}
