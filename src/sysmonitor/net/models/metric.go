package models

import (
	"time"
)

// Metric defines a single point measurement
type Metric struct {
	Tag      string        `json:"tag"`      //node_flow/vip_flow
	Fields   []string      `json:"fields"`   //metric name
	Values   []interface{} `json:"values"`   //metric datas
	Time     string        `json:"time"`     //gather metric time
	TimeZone time.Time     `json:"timezone"` //time zone
}

type Element struct {
	NetMetric         Metric `json:"net_metric"`
}
