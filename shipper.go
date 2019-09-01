package shipper

type Shipper struct {
	Reader Reader
	Writer Writer
}

func NewShipper(config Config) Shipper {
	if config.BatchSize == 0 {
		config.BatchSize = 10000
	}

	reader := Reader{
		SourcePath: config.Source,
		BatchSize:  config.BatchSize,
	}

	writer := Writer{
		DestinationPath: config.Destination,
		BatchSize:       config.BatchSize,
	}

	shipper := Shipper{
		Reader: reader,
		Writer: writer,
	}

	return shipper
}

func (shipper *Shipper) Ship() error {
	if openErr := shipper.Reader.open(); openErr != nil {
		return openErr
	}
	defer shipper.Reader.close()

	if openErr := shipper.Writer.open(); openErr != nil {
		return openErr
	}
	defer shipper.Writer.close()

	readErr, chunks := shipper.Reader.read()
	if readErr != nil {
		return readErr
	}

	if writeErr := shipper.Writer.write(chunks); writeErr != nil {
		return writeErr
	}

	return nil
}
