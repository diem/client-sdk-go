# libra-client-sdk-go

Go lang client sdk for connect to Libra FullNode JSON-RPC APIs.

## Sub Packages

- libraclient: high level APIs interface
- jsonrpc: jsonrpc client
- librakeys: keys utils, generate public & private keys for testing, create auth keys, and sign transaction
- libraid: Libra Account Identifier and Intent URL encoding & decodeing
- librastd: move stdlib script utils.
- testnet: testnet utils
- libratypes: Libra onchain data structure types.


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
