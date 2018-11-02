package mem

import (
	_"fmt"
	_"encoding/json"
	"sysmonitor/common"
	"sysmonitor/profile"
	"github.com/shirou/gopsutil/mem"
)


func MemMonitor() string {
	mem_info, _ := mem.VirtualMemory()
	mem_swap, _ := mem.SwapMemory()

	mem_status := new(profile.MemStatus)

	mem_status.Mem.Total =  mem_info.Total/1024/1024
	mem_status.Mem.Available =  mem_info.Available/1024/1024
	mem_status.Mem.Used =  mem_info.Used/1024/1024
	mem_status.Mem.Free =  mem_info.Free/1024/1024
	mem_status.Mem.MemUseage =  mem_info.UsedPercent

	mem_status.Swap.Total =  mem_swap.Total/1024/1024
	//mem_status.Swap.Available =  mem_swap.Available/1024/1024
	mem_status.Swap.Used =  mem_swap.Used/1024/1024
	mem_status.Swap.Free =  mem_swap.Free/1024/1024
	mem_status.Swap.SwapUseage =  mem_swap.UsedPercent
	mem_status.Swap.Sin =  mem_swap.Sin
	mem_status.Swap.Sout =  mem_swap.Sout

	/*
	result := common.JsonMarshal(mem_status)
    //fmt.Printf("        Mem        : %v  \n", string(result))
    return result
	*/

	mem_result := new (profile.Item)
    mem_result.Tag = "Mem"
    mem_result.Payload = mem_status
    result := common.JsonMarshal(mem_result)
    return result
}

