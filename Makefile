.PHONY: clean prepare build zip clean
all: clean prepare build zip clean

prepare:
	mkdir nd_darwin_386
	mkdir nd_darwin_amd64

build:
	env GOOS=darwin GOARCH=386 go build -ldflags="-X main.version=$(shell gov)" -o nd_darwin_386/nd
	env GOOS=darwin GOARCH=amd64 go build -ldflags="-X main.version=$(shell gov)" -o nd_darwin_amd64/nd

zip:
	zip nd_darwin_386.zip -r nd_darwin_386
	zip nd_darwin_amd64.zip -r nd_darwin_amd64

clean:
	rm -rf nd_darwin_386
	rm -rf nd_darwin_amd64
