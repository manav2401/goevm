package evm

type Memory struct {
	data []byte
}

// NewMemory simply initializes a new memory instance
func NewMemory() *Memory {
	return &Memory{}
}

// InitWithSize initializes the underlying byte array with given `size`
func (m *Memory) InitWithSize(size uint64) {
	m.data = make([]byte, size)
}

// Store sets the data in underlying array. It assumes that the array is already initialized.
func (m *Memory) Store(offset, size uint64, data []byte) {
	copy(m.data[offset:offset+size], data)
}

// Load returns the data in a specific index range
func (m *Memory) Load(offset, size uint64) []byte {
	return m.data[offset : offset+size]
}

// Len returns length of underlying data instance
func (m *Memory) Len() uint64 {
	return uint64(len(m.data))
}
