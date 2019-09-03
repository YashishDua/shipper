package shipper

import (
	"encoding/gob"
	"fmt"
	"strconv"
	"sync"
)

type Shipper struct {
	Reader    Reader
	Writer    Writer
	Transport Transport
}

func NewShipper(config Config) Shipper {
	if config.BatchSize == 0 {
		config.BatchSize = 10000
	}

	reader := Reader{
		SourcePath:     config.Source,
		SourceFileSize: config.SourceFileSize,
		BatchSize:      config.BatchSize,
	}

	writer := Writer{
		DestinationPath: config.Destination,
	}

	transport := Transport{
		Type: DefaultTransport,
	}

	if config.TCP != (TCP{}) {
		// TCP enabled
		transport = Transport{
			Type: TCPTransport,
			Host: config.TCP.Host,
			Port: config.TCP.Port,
		}
	} else {
		writer.BatchSize = reader.BatchSize
	}

	shipper := Shipper{
		Reader:    reader,
		Writer:    writer,
		Transport: transport,
	}

	return shipper
}

func (shipper *Shipper) ShipAndDock() error {
	if openErr := shipper.Reader.open(); openErr != nil {
		return openErr
	}
	defer shipper.Reader.close()

	if openErr := shipper.Writer.open(); openErr != nil {
		return openErr
	}
	defer shipper.Writer.close()

	p := make(chan Packet)

	routines, readErr := shipper.Reader.read(p)
	if readErr != nil {
		fmt.Println(readErr)
	}

	var (
		wg sync.WaitGroup
	)

	for i := 0; i < routines; i++ {
		packet := <-p
		wg.Add(1)
		shipper.Writer.write(&wg, packet)
	}

	wg.Wait()

	return nil
}

func (shipper *Shipper) Ship() error {
	conn, connErr := shipper.Transport.connect()
	if connErr != nil {
		return connErr
	}
	defer conn.Close()

	if openErr := shipper.Reader.open(); openErr != nil {
		return openErr
	}
	defer shipper.Reader.close()

	p := make(chan Packet)
	encoder := gob.NewEncoder(conn)

	routines, readErr := shipper.Reader.read(p)
	if readErr != nil {
		fmt.Println(readErr)
	}

	// Passing BatchSize
	encoder.Encode(Packet{
		Type:  BatchSize,
		Value: strconv.Itoa(shipper.Reader.BatchSize),
	})

	for i := 0; i < routines; i++ {
		packet := <-p
		encoder.Encode(packet)
	}

	// Passing EOF
	encoder.Encode(Packet{
		Type: EOF,
	})

	fmt.Println("File read successfully")

	return nil
}

func (shipper *Shipper) Dock() error {
	conn, connErr := shipper.Transport.listen()
	if connErr != nil {
		return connErr
	}
	defer conn.Close()

	if openErr := shipper.Writer.open(); openErr != nil {
		return openErr
	}
	defer shipper.Writer.close()

	var (
		wg sync.WaitGroup
	)

	dec := gob.NewDecoder(conn)

	for {
		packet := &Packet{}
		dec.Decode(packet)

		if packet.Type == BatchSize {
			shipper.Writer.BatchSize, _ = strconv.Atoi(packet.Value)
		} else if packet.Type == EOF {
			break
		} else if packet.Type == Chunk {
			if shipper.Writer.BatchSize == 0 {
				return fmt.Errorf("BatchSize Packet lost, restart the process")
			}

			wg.Add(1)
			shipper.Writer.write(&wg, *packet)
		}
	}

	wg.Wait()

	fmt.Println("File written successfully")

	return nil
}
