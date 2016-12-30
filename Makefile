
.PHONY: build
build:
	mkdir -p ./bin
	go build -o bin/ftest _functional_tests/*.go

.PHONY: test
test: build
	go test -v ./pkg/...

.PHONY: up
up:
	glide up --skip-vendor

.PHONY: bootstrap
bootstrap:
	glide install --skip-vendor

.PHONY: ftest
ftest:
	bin/ftest
