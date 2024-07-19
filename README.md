# goevm

A simulation of EVM (Ethereum Virtual Machine) in go. 

> [!WARNING]
> This work is experimental and should not be used in production.

The implementation contains the fundamental modules needed for the EVM i.e. stack, memory and storage. As of now, it only contains bunch of isolated opcodes (e.g. arithmetic operations) and opcodes interacting with memory and underlying storage/state. The whole list of opcodes supported can be found [here](./evm/jump_table.go)

### Storage

The [storage interface](./evm/storage.go) defines some generic methods which any storage should implement. There are 2 storage designs supported.
1. A [simple storage](./evm/simple_storage.go) -- A basic in-memory storage using map for storing account and state.  
2. A [remote storage](./evm/remote_storage.go) -- A storage which is pluggable to any geth based datadir (using level db and hash based scheme).

The simple storage is helpful to perform isolated simulations and testing. The remote storage provides a neat interface to interact with the underlying state of any existing EVM chain (which follows the same structure). To prevent data corruption on any existing chain's db, setter functions are not implemented for remote storage. It allows you to read balance, nonce and state data (e.g. contract slots) from any existing chain. Opcodes like `SLOAD` and `BALANCE` can read data from remote db.

### Tracing

The implementation contains a very simple tracer which logs the following things for each opcode.
- Stack changes (pre and post executing an opcode)
- Memory changes (including expansion, pre and post executing an opcode)
- Storage changes
    - Tracks all the reads for the remote storage
    - Track the state-diffs for the in-memory storage (i.e. pre and post execution values)
