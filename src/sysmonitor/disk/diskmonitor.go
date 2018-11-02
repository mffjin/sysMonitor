package disk
import (
	_"fmt"
	_"encoding/json"
	"sysmonitor/common"	
	"sysmonitor/profile"	
	"github.com/shirou/gopsutil/disk"
)

/*
type DiskStatus struct {
	Total uint64	
	Free uint64
	DiskUsage float64
}
*/

func DiskMonitor() string {
	disk, _ := disk.Usage("/")
	disk_status := new(profile.DiskStatus)	
	disk_status.Total = disk.Total/1024/1024/1024	
	disk_status.Free = disk.Free/1024/1024/1024	
	disk_status.DiskUsage = disk.UsedPercent

	disk_result := new (profile.Item)
    disk_result.Tag = "Disk"
    disk_result.Payload = disk_status
    result := common.JsonMarshal(disk_result)
    return result
}


