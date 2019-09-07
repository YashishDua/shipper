package shipper

import (
	"fmt"
	"os"
)

type Reader struct {
	SourcePath     string
	BatchSize      int
	SourceFile     *os.File
	SourceFileSize int
}

func (reader *Reader) open() error {
	sourceFile, err := os.Open(reader.SourcePath)
	if err != nil {
		return err
	}

	reader.SourceFile = sourceFile

	return nil
}

func (reader *Reader) init() (int, error) {
	var fileSize int

	if reader.SourceFileSize != 0 {
		fileSize = reader.SourceFileSize
	} else {
		fileinfo, err := reader.SourceFile.Stat()
		if err != nil {
			return 0, err
		}

		fileSize = int(fileinfo.Size())
	}

	routines := fileSize/reader.BatchSize + 1

	return routines, nil
}

func (reader *Reader) read(routines int, p chan Packet) {
	for i := 0; i < routines; i++ {
		go reader.chunkRead(p, i)
	}
}

func (reader *Reader) chunkRead(p chan Packet, index int) {
	offset := int64(index * reader.BatchSize)
	buffer := make([]byte, reader.BatchSize)
	bytesread, _ := reader.SourceFile.ReadAt(buffer, offset)

	p <- Packet{
		Type:  Chunk,
		Value: string(buffer[:bytesread]),
		Index: index,
	}
}

func (reader *Reader) close() {
	if closeErr := reader.SourceFile.Close(); closeErr != nil {
		fmt.Println(closeErr)
	}
}
