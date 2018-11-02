package main

import (
	_"time"
	_"fmt"
	"sysmonitor/persist"
	"sysmonitor/engine"
	"github.com/go-ini/ini"
)

func initConfig() (*ini.File, error) {
  cfg, err := ini.InsensitiveLoad("./config/monitor.ini")
	if err != nil {
	  return nil, err
	}
  return cfg, err
}

func main() {
    cfg, _ := initConfig()
	itemChan, _ := persist.ItemSaver()
	engine.MonitorEngine(cfg, itemChan)
}
