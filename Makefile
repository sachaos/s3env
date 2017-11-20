.PHONY: install
install:
	go install

.PHONY: test
test:
	go test -v

.PHONY: build
build:
	GOOS=linux GOARCH=amd64 go build -o artifacts/s3env_linux_amd64
	GOOS=linux GOARCH=arm go build -o artifacts/s3env_linux_arm
