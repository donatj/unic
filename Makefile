BIN=unic
USER=$(shell whoami)
HEAD=$(shell ([ -n "$${CI_TAG}" ] && echo "$$CI_TAG" || exit 1) || git describe --tags 2> /dev/null || git rev-parse --short HEAD)
STAMP=$(shell date -u '+%Y-%m-%d_%I:%M:%S%p')
DIRTY=$(shell test $(shell git status --porcelain | wc -l) -eq 0 || echo '(dirty)')


LDFLAGS="-X main.buildStamp=$(STAMP) -X main.buildUser=$(USER) -X main.buildHash=$(HEAD) -X main.buildDirty=$(DIRTY)"
all: install

.PHONY: build
build: release/darwin_universal release/linux_amd64 release/freebsd_amd64 release/windows_386 release/windows_amd64

.PHONY: test
test:
	go test './...'

.PHONY: clean
clean:
	-rm -f $(BIN)
	-rm -rf release dist
	go clean -i ./cmd/$(BIN)

.PHONY: install
install:
	go install -ldflags $(LDFLAGS) ./cmd/$(BIN)

release/darwin_amd64:
	env GOOS=darwin GOARCH=amd64 go clean -i ./cmd/$(BIN)
	env GOOS=darwin GOARCH=amd64 go build -ldflags $(LDFLAGS) -o release/darwin_amd64/$(BIN) ./cmd/$(BIN)

release/darwin_arm64:
	env GOOS=darwin GOARCH=arm64 go clean -i ./cmd/$(BIN)
	env GOOS=darwin GOARCH=arm64 go build -ldflags $(LDFLAGS) -o release/darwin_arm64/$(BIN) ./cmd/$(BIN)

release/darwin_universal: release/darwin_amd64 release/darwin_arm64
	mkdir release/darwin_universal
	lipo -create -output release/darwin_universal/$(BIN) release/darwin_amd64/$(BIN) release/darwin_arm64/$(BIN)

release/linux_amd64:
	env GOOS=linux GOARCH=amd64 go clean -i ./cmd/$(BIN)
	env GOOS=linux GOARCH=amd64 go build -ldflags $(LDFLAGS) -o release/linux_amd64/$(BIN) ./cmd/$(BIN)

release/freebsd_amd64:
	env GOOS=freebsd GOARCH=amd64 go clean -i ./cmd/$(BIN)
	env GOOS=freebsd GOARCH=amd64 go build -ldflags $(LDFLAGS) -o release/freebsd_amd64/$(BIN) ./cmd/$(BIN)

release/windows_386:
	env GOOS=windows GOARCH=386 go clean -i ./cmd/$(BIN)
	env GOOS=windows GOARCH=386 go build -ldflags $(LDFLAGS) -o release/windows_386/$(BIN).exe ./cmd/$(BIN)

release/windows_amd64:
	env GOOS=windows GOARCH=amd64 go clean -i ./cmd/$(BIN)
	env GOOS=windows GOARCH=amd64 go build -ldflags $(LDFLAGS) -o release/windows_amd64/$(BIN).exe ./cmd/$(BIN)

.PHONY: release
release: clean build
	mkdir dist
	zip -j 'dist/$(BIN).darwin_universal.$(HEAD)$(DIRTY).zip'  release/darwin_universal/$(BIN)
	zip -j 'dist/$(BIN).linux_amd64.$(HEAD)$(DIRTY).zip'       release/linux_amd64/$(BIN)
	zip -j 'dist/$(BIN).freebsd_amd64.$(HEAD)$(DIRTY).zip'     release/freebsd_amd64/$(BIN)
	zip -j 'dist/$(BIN).windows_386.$(HEAD)$(DIRTY).exe.zip'   release/windows_386/$(BIN).exe
	zip -j 'dist/$(BIN).windows_amd64.$(HEAD)$(DIRTY).exe.zip' release/windows_amd64/$(BIN).exe
