darwin64:
	env GOOS=darwin GOARCH=amd64 go clean -i
	env GOOS=darwin GOARCH=amd64 go build -o release/darwin64/unic

linux64:
	env GOOS=linux GOARCH=amd64 go clean -i
	env GOOS=linux GOARCH=amd64 go build -o release/linux64/unic

freebsd64:
	env GOOS=freebsd GOARCH=amd64 go clean -i
	env GOOS=freebsd GOARCH=amd64 go build -o release/freebsd64/unic

build: darwin64 linux64 freebsd64

.PHONY: release
release: build
	zip release/unic.darwin_amd64.zip release/darwin64/unic
	tar cJf release/unic.linux_amd64.tar.xz release/linux64/unic
	tar cJf release/unic.freebsd_amd64.tar.xz release/freebsd64/unic