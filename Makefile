install-buf:
	# Substitute GOBIN for your bin directory
	# Leave unset to default to $GOPATH/bin
	GO111MODULE=on GOBIN=/usr/local/bin go install github.com/bufbuild/buf/cmd/buf@v1.19.0

install-mockery:
	GO111MODULE=on GOBIN=/usr/local/bin go install github.com/vektra/mockery/v2@v2.27.1

install-tools: install-buf install-mockery

use-private:
	git config --global url."git@github.com".insteadOf "https://github.com"
	go env -w GOPRIVATE="github.com/releaseband"

buf-format:
	buf format proto --write --exit-code

buf-lint:
	buf lint proto

buf-gen:
	rm -rf gen
	buf generate

go-gen:
	find . -name 'mock_*_test.go' -type f -delete
	go generate ./...

run:
	go run .
