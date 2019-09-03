package main

import (
	"fmt"
	"shipper"
)

func main() {
	shipper := shipper.NewShipper(shipper.Config{
		Source:    "test.log",
		BatchSize: 5, // Total characters to be processed in 1 concurrent batch
		TCP: shipper.TCP{
			Host: "127.0.0.1",
			Port: 8001,
		},
	})

	if err := shipper.Ship(); err != nil {
		fmt.Printf("Error in Shipping: %s", err.Error())
	}
}
