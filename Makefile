test:
	go test ./...

cover:
	mkdir -p .tmp
	go test -covermode=count -coverprofile=.tmp/count.out ./...
	go tool cover -html=.tmp/count.out

fetch-libra-testnet:
	cd libra && git fetch && git reset --hard origin/testnet
