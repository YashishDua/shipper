package main

import (
	"fmt"

	"github.com/yashishdua/shipper"
)

func main() {
	shipper := shipper.NewShipper(shipper.Config{
		Source:    "test.log",
		BatchSize: 5000, // Total characters to be processed in 1 concurrent batch
		TCP: shipper.TCP{
			Host: "127.0.0.1",
			Port: 8001,
		},
	})

	if err := shipper.Ship(); err != nil {
		fmt.Printf("Error in Shipping: %s", err.Error())
	}
}
