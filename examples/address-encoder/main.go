package main

import (
	"encoding/binary"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/diem/client-sdk-go/diemid"
	"github.com/diem/client-sdk-go/diemkeys"
	"github.com/diem/client-sdk-go/diemtypes"
)

// networkToPrefix converts from a human friendly format to a prefix usable by the bech32 address format
func networkToPrefix(network string) (diemid.NetworkPrefix, error) {
	switch network {
	case "mainnet":
		return diemid.MainnetPrefix, nil
	case "premainnet":
		return diemid.PreMainnetPrefix, nil
	case "testnet":
		return diemid.TestnetPrefix, nil
	default:
		return diemid.NetworkPrefix(""), fmt.Errorf("Invalid network=%s supplied, no network prefix")
	}
}

// encode converts a onchainAddress + subAddress or publickey to a bech32 address format
// if you provide both, it will pick the onchainAddress and ignore publickey
func encode(networkPrefix diemid.NetworkPrefix, onchainAddress, publicKey string, subAddressNum uint64) (string, error) {
	var (
		accountAddress diemtypes.AccountAddress
		subAddress     diemtypes.SubAddress
		err            error
	)

	if onchainAddress == "" && publicKey == "" {
		return "", errors.New("Need at least onchain address or publickey to encode")
	}

	if onchainAddress != "" {
		accountAddress, err = diemtypes.MakeAccountAddress(onchainAddress)
		if err != nil {
			return "", fmt.Errorf("failed to make account address for %s: %w", onchainAddress, err)
		}
	}

	if publicKey != "" {
		pubKey, err := diemkeys.NewEd25519PublicKeyFromString(publicKey)
		if err != nil {
			return "", fmt.Errorf("failed to parse public key %w", err)
		}
		authKey := diemkeys.NewAuthKey(pubKey)
		accountAddress = authKey.AccountAddress()
	}

	if subAddressNum == 0 {
		subAddress = diemtypes.EmptySubAddress
	} else {
		subAddressBuf := make([]byte, 8)
		// Write the buffer as 00000000000000<num>
		binary.BigEndian.PutUint64(subAddressBuf, subAddressNum)
		subAddress, err = diemtypes.MakeSubAddressFromBytes(subAddressBuf)
		if err != nil {
			return "", fmt.Errorf("Error making subaddress from bytes %q: %w", subAddressBuf, err)
		}
	}

	return diemid.EncodeAccount(networkPrefix, accountAddress, subAddress)
}

// decode decodes the given bech32 encoded network public address and returns the hex account address and subAddress
func decode(networkPrefix diemid.NetworkPrefix, encodedAddress string) (string, string, error) {
	account, err := diemid.DecodeToAccount(networkPrefix, encodedAddress)
	if err != nil {
		return "", "", fmt.Errorf("Failed to decode to account: %w", err)
	}
	return account.AccountAddress.Hex(), account.SubAddress.Hex(), nil
}

func main() {
	var network = flag.String("network", "testnet", "Network to encode or decode addresses")
	var encodedAddress = flag.String("encoded-address", "", "Encoded address to use")
	var onChainAddress = flag.String("onchain-address", "", "Onchain address to use")
	var publicKey = flag.String("publickey", "", "Public key in hex format to use for generating address")
	var subAddress = flag.Uint64("subaddress", 0, "SubAddress to use, this will be left-padded with 0 to make 8 bytes to create a bech32 address")
	var task = flag.String("task", "", "Default task to do - encode or decode")
	flag.Parse()

	prefix, err := networkToPrefix(*network)
	if err != nil {
		panic(err)
	}

	switch *task {
	case "encode":
		address, err := encode(prefix, *onChainAddress, *publicKey, *subAddress)
		if err != nil {
			panic(err)
		}
		fmt.Println(address)
	case "decode":
		onChainAddress, subAddress, err := decode(prefix, *encodedAddress)
		if err != nil {
			panic(err)
		}
		subAddressNum, err := strconv.ParseUint(subAddress, 16, 64)
		if err != nil {
			panic(err)
		}
		fmt.Printf("OnchainAddress: %s\nsubAddress(hex): %s\nsubAddress(int): %d\n", onChainAddress, subAddress, subAddressNum)
	default:
		fmt.Printf("Unknown task %s\n", task)
		os.Exit(1)
	}
}
