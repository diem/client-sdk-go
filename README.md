# libra-client-sdk-go

[![API Reference](https://img.shields.io/badge/api-reference-blue.svg)](https://github.com/libra/libra/blob/master/json-rpc/json-rpc-spec.md) [![Apache V2 License](https://img.shields.io/badge/license-Apache%20V2-blue.svg)](../master/LICENSE)

libra-client-sdk-go is the official Libra Client SDK for the Go programming language.

## Overview of SDK's Packages

- libraclient: libra JSON-RPC APIs client
- jsonrpc: a JSON-RPC 2.0 SPEC client
- librakeys: keys utils, including generating public & private keys for testing, creating auth key and account address from public key.
- librasigner: sign transaction logic
- txnmetadata: utils for creating peer to peer transaction metadata. (LIP-4)
- libraid: encoding & decoding Libra Account Identifier and Intent URL. (LIP-5)
- testnet: testnet utils
- stdlib: move stdlib script utils. This is generated code, for constructing transaction script playload.
- libratypes: Libra on-chain data structure types. Mostly generated code with small extension code for attaching handy functions to generated types.
- [examples](../../tree/master/examples): examples of how to use this SDK.
  - [submit transaction and wait](../master/examples/exampleutils/submit_and_wait.go): this example shows how to submit a transaction and wait for its result; it also shows how to handle a stale response error in various cases.
  - [create child VASP account](../master/examples/create-child-vasp-account/main.go): this example shows how to create ChildVASP account for a ParentVASP account.
  - [p2p transfer](../master/examples/p2p-transfers/main.go): this example shows 4 different types p2p transfers between custodial accounts and non-custodial accounts.
  - [refund](../master/examples/refund/main.go): this example shows peer to peer transfers from custodial accounts to non-custodial accounts, and then refunds the amount.
  - [intent identifier](../master/examples/intent-identifier/main.go): this example shows how to use libraid for encoding and decoding the intent identifier / url.

## Installing

Use `go get` to retrieve the SDK to add it to your `GOPATH` workspace, or
project's Go module dependencies.

	go get github.com/libra/libra-client-sdk-go

To update the SDK use `go get -u` to retrieve the latest version of the SDK.

	go get -u github.com/libra/libra-client-sdk-go


## Development

*Run test*

```
make test
```

*Generate libratypes & move stdlib script encoder & decoder*

```
git submodule update
make gen
```

# API Documentation

The Go Client SDK API documentation is currently available at [godoc.org](https://godoc.org/github.com/libra/libra-client-sdk-go).

# License

[Apache License V2](../master/LICENSE)


# Contributing

[CONTRIBUTING](../master/CONTRIBUTING.md)
