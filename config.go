package shipper

type Config struct {
	Source         string // Source file path
	SourceFileSize int    // We do not recommend this
	Destination    string // Destination file path
	BatchSize      int    // Total characters to be processed in 1 concurrent batch
	TCP            TCP
}

type TCP struct {
	Host string
	Port int
}
