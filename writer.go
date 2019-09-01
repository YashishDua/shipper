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

func (writer *Writer) write(chunks []string) error {
	routines := len(chunks)

	var (
		wg sync.WaitGroup
	)

	wg.Add(routines)

	for i := 0; i < routines; i++ {
		go writer.chunkWrite(&wg, chunks[i], i)
	}

	wg.Wait()

	return nil
}

func (writer *Writer) chunkWrite(wg *sync.WaitGroup, chunk string, index int) error {
	defer wg.Done()

	offset := int64(index * writer.BatchSize)
	if _, writeErr := writer.DestinationFile.WriteAt([]byte(chunk), offset); writeErr != nil {
		return writeErr
	}

	return nil
}

func (writer *Writer) close() error {
	return writer.DestinationFile.Close()
}
