package shipper

type Shipper struct {
	Reader    Reader
	Writer    Writer
	Transport Transport
}

func NewShipper(config Config) Shipper {
	if config.BatchSize == 0 {
		config.BatchSize = 10000
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
	}

	reader := Reader{
		SourcePath:     config.Source,
		SourceFileSize: config.SourceFileSize,
		BatchSize:      config.BatchSize,
	}

	writer := Writer{
		DestinationPath: config.Destination,
		BatchSize:       config.BatchSize,
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

	readErr, chunks := shipper.Reader.read()
	if readErr != nil {
		return readErr
	}

	if writeErr := shipper.Writer.write(chunks); writeErr != nil {
		return writeErr
	}

	return nil
}

func (shipper *Shipper) Ship() error {
	if connErr := shipper.Transport.connect(); connErr != nil {
		return connErr
	}

	return nil
}

func (shipper *Shipper) Dock() error {
	if connErr := shipper.Transport.receive(); connErr != nil {
		return connErr
	}

	return nil
}
