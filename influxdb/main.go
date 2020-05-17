package main

import (
	"context"
	"fmt"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go"
)

func main() {
	client := influxdb2.NewClientWithOptions("http://localhost:8086", "",
		influxdb2.DefaultOptions().SetLogLevel(5))
	// Get non-blocking write client
	writeApi := client.WriteApiBlocking("", "test/autogen")
	line := fmt.Sprintf("stat,unit=temperature avg=%f,max=%f", 23.5, 45.0)
	_ = writeApi.WriteRecord(context.Background(), line)
	// write some points
	for i := 0; i < 100; i++ {
		// create point
		p := influxdb2.NewPoint(
			"system",
			map[string]string{"unit": "Ms"},
			map[string]interface{}{
				"mem_free": double(45.6),
			},
			time.Now())
		// write asynchronously
		err := writeApi.WritePoint(context.Background(), p)
		if err != nil {
			panic(err)
		}
	}
	// Force all unwritten data to be sent
	// writeApi.Flush()
	// Ensures background processes finishes
	client.Close()
}
