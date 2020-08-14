
# Basics

- [ ] module structure:
  - libra
    - libraclient: high level APIs interface, should support application to do easy mock / stub development.
      - types: data transfer object types for jsonrpc client, should match server side JSON-RPC spec data types.
      - utils: includes crypto, data types converting and other utils functions
        - signing, sha3 hashing, address parsing and converting, hex encoding / decoding
        - LCS utils
    - jsonrpc: jsonrpc client
    - stdlib: move stdlib script utils.
    - testnet: testnet utils, should include FaucetService for handling testnet mint.
    - types: Libra onchain data structure types.
- [x] JSON-RPC 2.0 Spec:
  - [x] spec version validation.
  - [x] batch requests and responses handling.
- [x] JSON-RPC client error handling should distinguish the following 3 type errors:
  - Transport layer error, e.g. HTTP call failure.
  - JSON-RPC protocol error: e.g. server responds to non json data, or can't be parsed into [Libra JSON-RPC SPEC][1] defined data structure, or missing result & error field.
  - JSON-RPC error: error returned from server.
- [x] https
- [x] Client connection pool: keep connection alive for less likely getting inconsistent data from connecting to multiple servers.
- [ ] Handle stale responses:
  - [ ] client tracks latest server response block version and timestamp, raise error when received server response contains stale version / timestamp.
  - [ ] parse and use libra_chain_id, libra_ledger_version and libra_ledger_tiemstamp in the JSONRPC response.
- [ ] Parsing and gen Libra Account Identifier (see [LIP-5][2])
  - bech32 addresses/subaddresses support
- [ ] language specific standard release publish: e.g. java maven central repo, python pip
- [ ] Multi-network: initialize Client with chain id, JSON-RPC server URL
- [x] Handle unsigned int64 data type properly
- [ ] Validate server chain id: client should be initialized with chain id and validate server response chain id is the same.
- [ ] Validate input parameters, e.g. invalid account address: "kkk". Should return / raise InvalidArgumentError.

# High Level API

- [ ] transfer: wrap peer to peer transfer with metadata script and submit transaction
  - may have the option to wait until the transaction executed successfully or failed.
- [ ] waitForTransactionExecuted(String accountAddress, long sequence, String signedTranscationHash, long timeout):
  - for given signed transaction sender address, sequence number, expiration time (or 5 sec timeout) to wait and validate execution result is executed, otherwise return/raise an error / flag to tell it is not executed.
  - when signedTransactionHash validation failed, it should return / raise TransactionSequenceNumberConflictError
  - when transaction execution vm_status is not "executed", it should return / raise TransactionExecutionFailure

# Read from Blockchain

- [x] Get metadata
- [x] Get currencies
- [x] Get events
- [x] Get transactions
- [x] Get account
- [x] Get account transaction
- [x] Get account transactions
- [ ] Get account events
- [x] Handle error response
- [x] Serialize result JSON to typed data structure

# Submit Transaction

- [ ] Submit [p2p transfer][3] transaction
- [ ] Submit other [Move Stdlib scripts][4]
- [ ] Wait for transaction executed:
  - wait for a transaction by get_transaction by account and transaction sequence, no validation of vm_status and signature. (low level API, consider not exposing, only for internal or test usage.)

# Testnet support

- [x] Generate ed25519 private key, derive ed25519 public keys from private key.
- [x] Generate Single auth-keys
- [ ] Generate MultiSig auth-keys
- [x] Mint coins through [Faucet service][6]

See [doc][5] for above concepts.

# Examples

- [ ] Query blockchain example
- [ ] Submit p2p transfer transaction example

# Nice to have

- [ ] Async client
- [ ] CLI connects to testnet for trying out features.

[1]: https://github.com/libra/libra/blob/master/json-rpc/json-rpc-spec.md "Libra JSON-RPC SPEC"
[2]: https://github.com/libra/lip/blob/master/lips/lip-5.md "LIP-5"
[3]: https://github.com/libra/libra/blob/master/language/stdlib/transaction_scripts/doc/peer_to_peer_with_metadata.md "P2P Transafer"
[4]: https://github.com/libra/libra/tree/master/language/stdlib/transaction_scripts/doc "Move Stdlib scripts"
[5]: https://github.com/libra/libra/blob/master/client/libra-dev/README.md "Libra Client Dev Doc"
[6]: https://github.com/libra/libra/blob/master/json-rpc/docs/service_testnet_faucet.md "Faucet service"
