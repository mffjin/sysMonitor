package engine
import (
	_"time"
	_"fmt"
	"os"
    "os/signal"
    _"strings"
    "sync"
    "syscall"
	"github.com/go-ini/ini"
)

var WorkConfigs map[string]map[string]interface{}
var WorkConfig map[string]interface{}

func InitConfig(cfg *ini.File) map[string]map[string]interface{}{
  WorkConfigs = make(map[string]map[string]interface{})
  allConfigSection := cfg.SectionStrings() 	
  for _,name := range allConfigSection[1:]{
	WorkConfig = make(map[string]interface{})
	  if (name == "interval") {
		keys := cfg.Section(name).KeyStrings()
		for _,key := range keys {
		  flag, _ := cfg.Section(name).Key(key).Int()
			WorkConfig[key] = flag
		}
	  }else{
		keys := cfg.Section(name).KeyStrings()
		for _,key := range keys {
		  flag, _ := cfg.Section(name).Key(key).Bool()
			if (flag) {
			  WorkConfig[key] = flag
			}
		}
	  }
	WorkConfigs[name] = WorkConfig
  }
 return WorkConfigs
}

func MonitorEngine(cfg *ini.File, itemChan chan interface{}) {
  worker := InitConfig(cfg)
  signals := make(chan os.Signal)
  signal.Notify(signals, os.Interrupt, syscall.SIGHUP)
  var wg sync.WaitGroup

  wg.Add(1)
  go func() {
        defer wg.Done()
        select {
        case sig := <-signals:
            if sig == os.Interrupt {
                os.Exit(0)
            }
        }
   }()

   wg.Add(1)
   go func() {
        defer wg.Done()
		if _, ok :=worker["sys"]["cpu"]; ok{
			SysMonitorProcess(worker, itemChan)
		}
   }()

   wg.Add(1)
   go func() {
        defer wg.Done()
		if _, ok :=worker["net"]["net_eth"]; ok{
			NetMonitorProcess(worker, itemChan)
		}
   }()
   wg.Wait()
}
