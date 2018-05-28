BINARY = sloc
GOARCH = amd64

FLAGS = src/sloc.go

all: clean test linux darwin windows

linux:
	GOOS=linux GOARCH=${GOARCH} go build -o bin/linux-${GOARCH}/${BINARY} ${FLAGS}

darwin:
	GOOS=darwin GOARCH=${GOARCH} go build -o bin/darwin-${GOARCH}/${BINARY} ${FLAGS}

windows:
	GOOS=windows GOARCH=${GOARCH} go build -o bin/windows-${GOARCH}/${BINARY}.exe ${FLAGS}

test:
	go test ./test/...

clean:
	-rm -rf bin/

.PHONY: linux darwin windows test clean
