test:
	go clean -cache -testcache
	go list ./... | grep examples | grep -v transaction-builder | xargs go build
	go list ./... | grep -v /examples/ | xargs go test

cover:
	mkdir -p .tmp
	go test -covermode=count -coverprofile=.tmp/count.out ./...
	go tool cover -html=.tmp/count.out

gen:
	cd diem && cargo build -p transaction-builder-generator && target/debug/generate-transaction-builders \
		--language go \
		--module-name stdlib \
		--diem-package-name github.com/diem/client-sdk-go \
		--with-diem-types "testsuite/generate-format/tests/staged/diem.yaml" \
		--target-source-dir ".." \
		"language/stdlib/compiled/transaction_scripts/abi"

protoc:
	# protoc  --go_out=. --go_opt=paths=source_relative ./diemjsonrpctypes/jsonrpc.proto
	protoc -Idiem/json-rpc/types/src/proto --go_out=./diemjsonrpctypes --go_opt=paths=source_relative jsonrpc.proto
