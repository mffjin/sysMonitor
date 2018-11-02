package net

import (
	"errors"
	"fmt"
	//models "github.com/flowagent/models"
	models "sysmonitor/net/models"	
	"net"
	"os"
	"strconv"
	"strings"
	"time"
	psutilnet "github.com/shirou/gopsutil/net"
)

func isFoundInterface(eth string, eths []string) bool {
	var found bool
	for _, name := range eths {
		if strings.EqualFold(name, eth) {
			found = true
			break
		}
	}
	return found
}

func NetIOCounters(hosttype string, extra_eths []string, intra_eths []string) (models.Metric, error) {
	var metric models.Metric
	net_iocount, _ := psutilnet.IOCounters(true)
	hostname, err := os.Hostname()
	if err != nil {
		return models.Metric{}, err
	}
	nodename := hostname
	for _, net_ := range net_iocount {
		interfaceName := net_.Name
		if interfaceName == "" {
			continue
		}

		iface, err := net.InterfaceByName(interfaceName)
		if err != nil {
			continue
		}

		if iface.Flags&net.FlagLoopback == net.FlagLoopback {
			continue
		}
		if iface.Flags&net.FlagUp == 0 {
			continue
		}
		extra_found := isFoundInterface(interfaceName, extra_eths)
		intra_found := isFoundInterface(interfaceName, intra_eths)
		var extra int
		if intra_found {
			extra = 0
		} else if extra_found {
			extra = 1
		} else {
			continue
		}
		metric.Values = append(metric.Values, []interface{}{interfaceName,
			hosttype,
			hostname,
			nodename,
			extra,
			net_.BytesSent,
			net_.BytesRecv,
			net_.PacketsSent,
			net_.PacketsRecv,
			net_.Errin,
			net_.Errout,
			net_.Dropin,
			net_.Dropout})
		metric.Time = strconv.FormatUint(uint64(time.Now().UnixNano()/1e6), 10)
	}

	local, err := time.LoadLocation("Local")
	if err != nil {
		return models.Metric{}, err
	}
	metric.TimeZone = time.Now().In(local)
	metric.Tag = "node_flow"
	metric.Fields = []string{"eth",
		"hosttype",
		"hostname",
		"nodename",
		"extra",
		"outbps",
		"inbps",
		"outpps",
		"inpps",
		"errin",
		"errout",
		"dropin",
		"dropout"}

	return metric, nil
}

func TruncTime(mtime string, interv int) (string, error) {
	sunix, err := strconv.ParseInt(mtime, 10, 64)
	if err != nil {
		return "", err
	}
	stime := time.Unix(sunix/1000, 0).UTC().Format("2006-01-02 15:04:05")
	return stime, nil
}

func NetMultiSub(smetric, emetric models.Metric, interv int) (models.Metric, error) {
	fmt.Println("============== netmultisub ==========================")
	var metric models.Metric
	var err error
	metric.Time, err = TruncTime(emetric.Time, interv)

	if err != nil {
		return models.Metric{}, err
	}

	metric.TimeZone = emetric.TimeZone


	metric.Tag = emetric.Tag
	metric.Fields = emetric.Fields

	//src  des
	stime, _ := strconv.ParseUint(smetric.Time, 10, 64)
	etime, _ := strconv.ParseUint(emetric.Time, 10, 64)

	values := make([]interface{}, len(smetric.Values))

	for i := 0; i < len(smetric.Values); i++ {
		scols := smetric.Values[i].([]interface{})
		ecols := emetric.Values[i].([]interface{})

		temp := make([]interface{}, len(scols))
		for j := 0; j < 5; j++ {
			temp[j] = scols[j]
		}
		for j := 5; j < len(scols); j++ {
			svalue := (scols[j]).(uint64)
			evalue := (ecols[j]).(uint64)

			if evalue < svalue || (float64)(etime-stime) < 1000.0 {
				return models.Metric{}, errors.New(fmt.Sprintf("flow before:%s > after:%s [%s:%s]", svalue, evalue, smetric.Time, emetric.Time))
			}

			// byte outbps inbps outpps inpps......
			value := evalue - svalue
			if j == 5 || j == 6 {
				value *= 8
			}
			temp[j] = uint64(float64(value)/(float64)(etime-stime)) * 1000
		}
		values[i] = temp
	}

	metric.Values = values
	return metric, nil
}
