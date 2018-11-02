package engine

import (
	"time"
	_"fmt"
    "sysmonitor/host"
    "sysmonitor/cpu"
    "sysmonitor/disk"
    "sysmonitor/mem"
    "sysmonitor/load"
)

func SysMonitorProcess(worker map[string]map[string]interface{}, itemChan chan interface{}) {
	sys_ticker := time.NewTicker(time.Second * time.Duration(worker["interval"]["sys_interval"].(int)))
	for {
	  select{	
		case <- sys_ticker.C:

		  if _, ok :=worker["sys"]["cpu"]; ok{
			go func(){
				sys_info := host.GetSysInfo()
				itemChan <- sys_info
			}()
		  }

		  if _, ok :=worker["sys"]["cpu"]; ok{
			go func(){
				cpu_info := cpu.CpuMonitor()
				itemChan <- cpu_info
			}()
		  }

		  if _, ok :=worker["sys"]["disk"]; ok{
			go func(){
				disk_info := disk.DiskMonitor()
			 	itemChan <- disk_info
			}()
		  }

		  if _, ok :=worker["sys"]["mem"]; ok{
			go func() {
				mem_info := mem.MemMonitor()
				itemChan <- mem_info
			}()
		  }

		  if _, ok :=worker["sys"]["load"]; ok{
			go func() { 
				load_info := load.LoadMonitor()
			 	itemChan <- load_info
			}()
		  }
 	 }
  }
}
