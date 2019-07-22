#http://blog.wrouesnel.com/articles/Totally%20static%20Go%20builds/
GOBUILD := go build -a -ldflags '-extldflags "-static"'
BUILD_ENV := GO111MODULE=on CGO_ENABLED=0
EXE_NAME := th


.PHONY: gofmt
gofmt:
	GO111MODULE=on go fmt ./...

.PHONY: build
build: gofmt
	$(BUILD_ENV) GOOS=darwin GOARCH=386 $(GOBUILD) -o bin/macos/$(EXE_NAME)
	$(BUILD_ENV) GOOS=linux GOARCH=amd64 $(GOBUILD) -o bin/linux/$(EXE_NAME)
	$(BUILD_ENV) GOOS=windows GOARCH=386 $(GOBUILD) -o bin/windows/$(EXE_NAME).exe

.PHONY: goget
goget:
	GO111MODULE=on go get ./...

.PHONY: tidy
tidy:
	GO111MODULE=on go mod tidy
