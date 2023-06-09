.PHONY: all linux macos

all: linux macos

linux:
	cd src && GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o ../scanner-linux-amd64

macos:

	cd src && GOOS=darwin GOARCH=arm64 go build -ldflags "-w -s" -o ../scanner-macos-arm64