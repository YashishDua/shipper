package shipper

import (
	"os"
	"sync"
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

func (reader *Reader) read() (error, []string) {
	var (
		wg       sync.WaitGroup
		fileSize int
	)

	if reader.SourceFileSize != 0 {
		fileSize = reader.SourceFileSize
	} else {
		fileinfo, err := reader.SourceFile.Stat()
		if err != nil {
			return err, nil
		}

		fileSize = int(fileinfo.Size())
	}

	routines := fileSize/reader.BatchSize + 1
	chunks := make([]string, routines)

	wg.Add(routines)

	for i := 0; i < routines; i++ {
		go reader.chunkRead(&wg, chunks, i)
	}

	wg.Wait()

	return nil, chunks
}

func (reader *Reader) chunkRead(wg *sync.WaitGroup, chunks []string, index int) {
	defer wg.Done()

	offset := int64(index * reader.BatchSize)
	buffer := make([]byte, reader.BatchSize)
	bytesread, _ := reader.SourceFile.ReadAt(buffer, offset)

	chunks[index] = string(buffer[:bytesread])
}

func (reader *Reader) close() error {
	return reader.SourceFile.Close()
}
