test:
	go clean -cache -testcache
	go list ./... | grep examples | grep -v transaction-builder | xargs go build
	go list ./... | grep -v /examples/ | xargs go test

cover:
	mkdir -p .tmp
	go test -covermode=count -coverprofile=.tmp/count.out ./...
	go tool cover -html=.tmp/count.out

gen:
	cd libra && cargo build -p transaction-builder-generator && target/debug/generate-transaction-builders \
		--language go \
		--module-name stdlib \
		--libra-package-name github.com/libra/libra-client-sdk-go \
		--with-libra-types "testsuite/generate-format/tests/staged/libra.yaml" \
		--target-source-dir ".." \
		"language/stdlib/compiled/transaction_scripts/abi"
