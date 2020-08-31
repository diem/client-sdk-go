
# Basics

- [x] module structure:
  - libra
    - libraclient: high level APIs interface, should support application to do easy mock / stub development.
    - jsonrpc: jsonrpc client
    - librakeys: keys utils, generate public & private keys for testing, create auth key and account address.
    - librasigner: sign transaction logic
    - txnmetadata: utils for creating peer to peer transaction metadata. (LIP-4)
    - libraid: encoding & decodeing Libra Account Identifier and Intent URL. (LIP-5)
    - stdlib: move stdlib script utils.
    - testnet: testnet utils, should include FaucetService for handling testnet mint.
    - libratypes: Libra onchain data structure types.
- [x] JSON-RPC 2.0 Spec:
  - [x] spec version validation.
  - [x] batch requests and responses handling.
- [x] JSON-RPC client error handling should distinguish the following 3 type errors:
  - Transport layer error, e.g. HTTP call failure.
  - JSON-RPC protocol error: e.g. server responds to non json data, or can't be parsed into [Libra JSON-RPC SPEC][1] defined data structure, or missing result & error field.
  - JSON-RPC error: error returned from server.
- [x] https
- [x] Handle stale responses:
  - [x] client tracks latest server response block version and timestamp, raise error when received server response contains stale version / timestamp.
  - [x] parse and use libra_chain_id, libra_ledger_version and libra_ledger_tiemstamp in the JSONRPC response.
- [x] language specific standard release publish: e.g. java maven central repo, python pip
- [x] Multi-network: initialize Client with chain id, JSON-RPC server URL
  - [x] Validate server chain id: client should be initialized with chain id and validate server response chain id is the same.
- [x] Handle unsigned int64 data type properly
- [x] [Multi-signatures support](https://github.com/libra/libra/blob/master/specifications/crypto/spec.md#multi-signatures)
- [x] Transaction hash: for a given signed transaction, produce hash of the transaction executed.
  - hex-encode(sha3-256([]byte("Transaction")) + []byte {0} + signed transaction bytes)

# [LIP-4][7] support

- [x] Non-custodial to custodial transaction
- [x] Custodial to non-custodial transaction
- [x] Custodial to Custodial transaction
- [x] Refund

# [LIP-5][2] support

- [x] Encode and decode account identifier
- [x] Encode and decode intent identifier

# Read from Blockchain

- [x] Get metadata
- [x] Get currencies
- [x] Get events
- [x] Get transactions
- [x] Get account
- [x] Get account transaction
- [x] Get account transactions
- [x] Handle error response
- [x] Serialize result JSON to typed data structure

# Submit Transaction

- [x] Submit [p2p transfer][3] transaction
- [x] Submit other [Move Stdlib scripts][4]
- [x] waitForTransaction(accountAddress, sequence, transcationHash, expirationTimeSec, timeout):
  - for given signed transaction sender address, sequence number, expiration time (or 5 sec timeout) to wait and validate execution result is executed, otherwise return/raise an error / flag to tell it is not executed.
  - when signedTransactionHash validation failed, it should return / raise TransactionSequenceNumberConflictError
  - when transaction execution vm_status is not "executed", it should return / raise TransactionExecutionFailure
  - when transaction expired, it should return / raise TransactionExpiredError: compare the transaction expirationTimeSec with response latest ledger timestamp. If response latest ledger timestamp >= transaction expirationTimeSec, then we are sure the transaction will never be executed successfully.
    - Note: response latest ledger timestamp unit is microsecond, expirationTimeSec's unit is second.

# Testnet support

- [x] Generate ed25519 private key, derive ed25519 public keys from private key.
- [x] Generate Single auth-keys
- [x] Generate MultiSig auth-keys
- [x] Mint coins through [Faucet service][6]

See [doc][5] for above concepts.

# Examples

- [x] [p2p transfer examples](https://github.com/libra/lip/blob/master/lips/lip-4.md#transaction-examples)
- [x] refund p2p transfer example
- [x] create childVASP example
- [x] Intent identifier encoding, decoding example

# Nice to have

- [ ] Async client
- [ ] CLI connects to testnet for trying out features.
- [x] Client connection pool.

[1]: https://github.com/libra/libra/blob/master/json-rpc/json-rpc-spec.md "Libra JSON-RPC SPEC"
[2]: https://github.com/libra/lip/blob/master/lips/lip-5.md "LIP-5"
[3]: https://github.com/libra/libra/blob/master/language/stdlib/transaction_scripts/doc/peer_to_peer_with_metadata.md "P2P Transafer"
[4]: https://github.com/libra/libra/tree/master/language/stdlib/transaction_scripts/doc "Move Stdlib scripts"
[5]: https://github.com/libra/libra/blob/master/client/libra-dev/README.md "Libra Client Dev Doc"
[6]: https://github.com/libra/libra/blob/master/json-rpc/docs/service_testnet_faucet.md "Faucet service"
[7]: https://github.com/libra/lip/blob/master/lips/lip-4.md "Transaction Metadata Specification"
