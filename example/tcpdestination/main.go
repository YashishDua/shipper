package main

import (
	"fmt"
	"shipper"
)

func main() {
	shipper := shipper.NewShipper(shipper.Config{
		Destination: "test2.log",
		TCP: shipper.TCP{
			Port: 8001,
		},
	})

	if err := shipper.Dock(); err != nil {
		fmt.Printf("Error in Docking: %s", err.Error())
	}
}
