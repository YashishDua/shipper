package shipper

type Packet struct {
	Type  PacketType
	Value string
	Index int
}

type PacketType string

const (
	Chunk     PacketType = "Chunk"
	EOF       PacketType = "EOF"
	BatchSize PacketType = "BatchSize"
)
