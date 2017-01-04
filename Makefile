
.PHONY: build
build:
	mkdir -p ./bin
	go build -o bin/ftest _functional_tests/*.go
	go build -o bin/quokka cmd/quokka/*.go

.PHONY: test
test: build
	go test -v ./pkg/...

.PHONY: up
up:
	glide up --strip-vendor

.PHONY: bootstrap
bootstrap:
	glide install --strip-vendor

.PHONY: ftest
ftest:
	bin/ftest
