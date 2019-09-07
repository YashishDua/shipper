package main

import (
	"fmt"
	"shipper"
)

func main() {
	shipper := shipper.NewShipper(shipper.Config{
		Source:      "test.log",
		Destination: "test2.log",
		BatchSize:   5000, // Total characters to be processed in 1 concurrent batch
	})

	if err := shipper.ShipAndDock(); err != nil {
		fmt.Printf("Error in ShipAndDock: %s", err.Error())
	}
}
