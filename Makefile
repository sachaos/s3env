ARTIFACTS_DIR=artifacts/${VERSION}
GITHUB_USERNAME=sachaos

.PHONY: install
install:
	go install

.PHONY: test
test:
	go test -v

.PHONY: release
release:
	GOOS=linux GOARCH=amd64 go build -o $(ARTIFACTS_DIR)/s3env_linux_amd64
	GOOS=linux GOARCH=arm go build -o $(ARTIFACTS_DIR)/s3env_linux_arm
	GOOS=darwin GOARCH=amd64 go build -o $(ARTIFACTS_DIR)/s3env_darwin_amd64
	ghr -u $(GITHUB_USERNAME) -t $(shell cat github_token) --replace ${VERSION} $(ARTIFACTS_DIR)
