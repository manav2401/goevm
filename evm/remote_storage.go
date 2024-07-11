package evm

// RemoteStorage represents the disk based store of an existing
// geth based EVM node. This acts as an interface to interact
// with the state trie (i.e. accounts, contract, etc) of an
// existing node.
type RemoteStorage struct {
	// TODO: set a db instace here to interact with state
	data map[Key]Value
}

func NewRemoteStorage() *RemoteStorage {
	// Initialize a geth based db here
	return &RemoteStorage{
		data: make(map[Key]Value),
	}
}

func (s *RemoteStorage) Store(key Key, value Value) {
	// TODO: Write into db, not map
	s.data[key] = value
}

func (s *RemoteStorage) Load(key Key) (value Value) {
	// TODO: Read from db, not map
	return s.data[key]
}
