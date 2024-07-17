package evm

type Memory struct {
	data []byte
}

// NewMemory simply initializes a new memory instance
func NewMemory() *Memory {
	return &Memory{make([]byte, 0)}
}

func (m *Memory) Resize(size uint64) {
	if uint64(m.Len()) < size {
		m.data = append(m.data, make([]byte, size-uint64(m.Len()))...)
	}
}

// Store sets the data in underlying array. It assumes that the array is already initialized.
func (m *Memory) Store(offset, size uint64, data []byte) {
	copy(m.data[offset:offset+size], data)
}

// Load returns the data in a specific index range
func (m *Memory) Load(offset, size uint64) []byte {
	if size == 0 || offset >= m.Len() {
		data := make([]byte, size)
		return data
	}

	if offset+size > m.Len() {
		data := make([]byte, size)
		copy(data, m.data[offset:])
	}

	return m.data[offset : offset+size]
}

// Len returns length of underlying data instance
func (m *Memory) Len() uint64 {
	return uint64(len(m.data))
}
