test:
	go test ./...

cover:
	go test -covermode=count -coverprofile=count.out ./...
	go tool cover -html=count.out

fetch-libra-testnet:
	cd libra && git fetch && git reset --hard origin/testnet
