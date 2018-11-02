package persist

import (
	"fmt"
	"context"
	"encoding/json"
	"sysmonitor/profile"
	"github.com/olivere/elastic"
)

func ItemSaver() (chan interface{}, error){
    //client, err := elastic.NewClient(elastic.SetURL("http://192.168.187.185:9200"))
	/*
    _, err := elastic.NewClient(elastic.SetURL("http://192.168.187.185:9200"))
    if err != nil {
        return nil, err
    }
	*/

    out := make(chan interface{})
    go func() {
        for {
            item := <- out
			var b profile.Item
			json.Unmarshal([]byte(item.(string)), &b)
			switch b.Tag {
				case "CPU":
            		fmt.Println("Item Saver: got item %v     %v",b.Tag, item)
					//save(client, "cpu_info", "cpu", item)
				case "Mem":
            		fmt.Println("Item Saver: got item %v     %v",b.Tag, item)
					//save(client, "mem_info", "mem", item)
				case "Disk":
            		fmt.Println("Item Saver: got item %v     %v",b.Tag, item)
					//save(client, "disk_info", "disk", item)
				case "Host":
            		fmt.Println("Item Saver: got item %v     %v",b.Tag, item)
					//save(client, "host_info", "host", item)
				case "Load":
            		fmt.Println("Item Saver: got item %v     %v",b.Tag, item)
					//save(client, "load_info", "load", item)
				case "Net_monitor":
            		fmt.Println("Item Saver: got item %v     %v",b.Tag, item)
					//save(client, "net_monitor_info", "net", item)
				case "Net_connect":
            		fmt.Println("Item Saver: got item %v     %v",b.Tag, item)
					//save(client, "net_connect_info", "net", item)
				//default:
				case "node_flow":
            		fmt.Println("Item Saver: got item %v     %v",b.Tag, item)
		    }
        }
    }()
    return out, nil
}

func save(client *elastic.Client, index string, index_type string,  item interface{}) {
    indexService := client.Index().Index(index).
        Type(index_type).BodyJson(item)

    _, err := indexService.Do(context.Background())

    if err != nil {
        fmt.Println("item saver error")
    }
}
