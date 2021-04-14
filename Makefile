

.PHONY: test
test: test_prepare
	gotest -v ./...

build/tools/golangci-lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b build/tools/ v1.31.0
	chmod +x build/tools/golangci-lint

.PHONY: lint
lint: build/tools/golangci-lint
	build/tools/golangci-lint run -v --timeout 2m

.PHONY: test_prepare
test_prepare: 
	GO111MODULE=off go get github.com/rakyll/gotest