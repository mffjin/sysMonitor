package engine

import (
	"time"
	_"fmt"
    "sysmonitor/net"
)

func NetMonitorProcess(worker  map[string]map[string]interface{}, itemChan chan interface{}) {
	interv := worker["interval"]["net_interval"].(int) //collect_interv = 30
	net_ticker := time.NewTicker(time.Second * time.Duration(interv))
	src := net.NetMonitor()
	for{
		select {	
		  case <-net_ticker.C: 
			if _, ok :=worker["net"]["net_eth"]; ok{	
			  go func() {
					dest := net.NetMonitor()
					tempSrc := src
            		src = dest

            		net_info := net.CollectElement(tempSrc, dest, interv)
					/*
            		if err != nil {
                		log.Log.Error("collect: ", err)
                		continue
            		}
					*/
					itemChan <- net_info
			  }()
			}

		   if _, ok :=worker["net_connect_protocol"]["tcp"]; ok{	
			 go func(){
				net_connect := net.ConnectMonitor()
			   	itemChan <- net_connect
			 }()
		   }
		}
    }
}
