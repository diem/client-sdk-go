test:
	go test ./...

fetch-libra-testnet:
	cd libra && git fetch && git reset --hard origin/testnet
