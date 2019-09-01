package shipper

type Config struct {
	Source      string // Source file path
	Destination string // Destination file path
	BatchSize   int    // Total characters to be processed in 1 async batch
}
