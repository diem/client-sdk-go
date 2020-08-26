# libra-client-sdk-go

[![API Reference](https://img.shields.io/badge/api-reference-blue.svg)](https://github.com/libra/libra/blob/master/json-rpc/json-rpc-spec.md) [![Apache V2 License](https://img.shields.io/badge/license-Apache%20V2-blue.svg)](../blob/LICENSE)

libra-client-sdk-go is the official Libra Client SDK for the Go programming language.

## Overview of SDK's Packages

- libraclient: libra JSON-RPC APIs client
- jsonrpc: a JSON-RPC 2.0 SPEC client
- librakeys: keys utils, including generating public & private keys for testing, creating auth keys, and signing transaction
- libraid: encoding & decodeing Libra Account Identifier and Intent URL
- testnet: testnet utils
- stdlib: move stdlib script utils. This is generated code, for constructing transaction script playload.
- libratypes: Libra onchain data structure types. Mostly generated code with small extension code for attaching handy functions to generated types.
- [examples](../blob/examples): examples of how to use this SDK.

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

*Upgrade to latest libra testnet release*

```
make fetch-libra-testnet

```

*Generate libratypes & move stdlib script encoder & decoder*

```
make gen
```


# License

[Apache License V2](../blob/LICENSE)


# Contributing

[CONTRIBUTING](../blob/CONTRIBUTING.md)
