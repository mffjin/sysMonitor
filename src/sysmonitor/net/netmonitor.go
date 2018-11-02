package net

import (
	_"fmt"
	"strings"
	_"errors"
	_"encoding/json"
    "sysmonitor/common"
	models "sysmonitor/net/models"
    _"sysmonitor/profile"
	_"github.com/shirou/gopsutil/net"
)


func NetMonitor() models.Element {
  var element models.Element

  element.NetMetric, _ = NetIOCounters("test", strings.Split("eth1", ","), strings.Split("eth0", ","))

  /*
  result := common.JsonMarshal(element.NetMetric)
  return result
  */
  return element
}

//func CollectElement(src, dest models.Element, interv int) (models.Element, error) {
func CollectElement(src, dest models.Element, interv int) string {
    var element models.Element
    //var err error

    if len(src.NetMetric.Values) > 0 && len(dest.NetMetric.Values) > 0 {
        element.NetMetric,_ = NetMultiSub(src.NetMetric, dest.NetMetric, interv)
		/*
        if err != nil {
            log.Log.Error("net ", src, dest)
            return models.Element{}, err
        }
		*/
        //log.Log.Info("net time: ", element.NetMetric.Time)
    }

	/*
    if len(src.NetMetric.Values) == 0  {
        log.Log.Error("element error: ", src, dest)
        return models.Element{}, errors.New("element error")
    }
	*/

  	result := common.JsonMarshal(element.NetMetric)
  	return result

    //return element, nil
}
