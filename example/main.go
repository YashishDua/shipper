package main

import (
	"fmt"
	"shipper"
)

func main() {
	shipper := shipper.NewShipper(shipper.Config{
		Source:      "test.log",
		Destination: "test2.log",
		BatchSize:   5, // Total characters to be processed in 1 async batch
	})

	if err := shipper.Ship(); err != nil {
		fmt.Printf("Error in Shipping: %s", err.Error())
	}
}
