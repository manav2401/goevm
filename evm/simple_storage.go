package evm

// SimpleStorage is an in-memory store with a simple map underneath
type SimpleStorage struct {
	data map[Key]Value
}

func NewSimpleStorage() *SimpleStorage {
	return &SimpleStorage{
		data: make(map[Key]Value),
	}
}

func (s *SimpleStorage) Store(key Key, value Value) {
	s.data[key] = value
}

func (s *SimpleStorage) Load(key Key) (value Value) {
	return s.data[key]
}
