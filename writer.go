package shipper

import (
	"os"
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
	for i := 0; i < len(chunks); i++ {
		if _, err := writer.DestinationFile.WriteString(chunks[i]); err != nil {
			return err
		}
	}

	return nil
}
