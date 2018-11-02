package common

import (
	_"time"
	_"fmt"
	"encoding/json"
	_"sysmonitor/profile"
)

func JsonMarshal(Network interface{}) string {
	   b, err := json.Marshal(Network)
       if err != nil {
          return ""
       } else {
          return string(b)
       }
}

/*
func GetLocalTime() string {
	local, err := time.LoadLocation("Local")
    if err != nil {
        fmt.Println("error")
    }

    TimeZone := time.Now().In(local)
    fmt.Println(TimeZone)
	return TimeZone
}
*/
