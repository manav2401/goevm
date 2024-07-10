# goevm

A simulation of EVM (Ethereum Virtual Machine) in go. Do not use in production.

The aim is to do the following things:
1. Have a simulated EVM (at least with limited opcodes to begin with).
2. Have tracing support.
3. Usable against a forked chain to access the state.

Checklist:

- [ ] Setup basic data structures required for EVM (e.g. stack)
- [ ] Implement very basic opcodes (e.g. arithmetic ones)
- [ ] Test against an isolated transaction
- [ ] Implement interface to interact with existing state. Need to think a bit more but the idea is to build an interface which interacts with the state. Hence, during a transaction execution, all the read/writes should happen in the underlying state.