#http://blog.wrouesnel.com/articles/Totally%20static%20Go%20builds/
GOBUILD := go build -a -ldflags '-extldflags "-static"'
EXE_NAME := th


.PHONY: gofmt
gofmt:
	GO111MODULE=on go fmt ./...

.PHONY: build
build: gofmt
	CGO_ENABLED=0 GOOS=darwin GOARCH=386 $(GOBUILD) -o bin/macos/$(EXE_NAME)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o bin/linux/$(EXE_NAME)
	CGO_ENABLED=0 GOOS=windows GOARCH=386 $(GOBUILD) -o bin/windows/$(EXE_NAME).exe

.PHONY: goget
goget:
	GO111MODULE=on go get ./...

.PHONY: tidy
tidy:
	GO111MODULE=on go mod tidy
