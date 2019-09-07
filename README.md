# shipper
High speed file transfer

## Features
1. Concurrent and high speed shipping
2. Ships in the same machine (Non-TCP)
3. Ships between two machines (TCP)


## Setup
```sh
go get github.com/yashishdua/shipper
```

#### Note: Example code to run shipper is placed in the /example directory

## 1. Ship in the same machine (non-TCP)

```go
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
```

## 2. Ship between two machines (TCP)

### TCP Source

```go
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
```

### TCP Destination

```go
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
```

