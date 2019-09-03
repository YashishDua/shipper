package shipper

import (
	"os"
	"sync"
)

type Writer struct {
	DestinationPath string
	BatchSize       int
	DestinationFile *os.File
}

func (writer *Writer) open() error {
	destinationFile, err := os.OpenFile(writer.DestinationPath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	writer.DestinationFile = destinationFile

	return nil
}

func (writer *Writer) write(wg *sync.WaitGroup, packet Packet) {
	go writer.chunkWrite(wg, packet)
}

func (writer *Writer) chunkWrite(wg *sync.WaitGroup, packet Packet) error {
	defer wg.Done()

	offset := int64(packet.Index * writer.BatchSize)
	if _, writeErr := writer.DestinationFile.WriteAt([]byte(packet.Value), offset); writeErr != nil {
		return writeErr
	}

	return nil
}

func (writer *Writer) close() error {
	return writer.DestinationFile.Close()
}
