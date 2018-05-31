BINARY = sloc

ARCH32_386 = 386
ARCH32_arm = arm
ARCH64_amd64 = amd64
ARCH64_arm64 = arm64

FLAGS = src/main.go

all: clean test linux darwin windows freebsd solaris

linux:
	GOOS=linux GOARCH=${ARCH32_386} go build -o bin/linux/${ARCH32_386}/${BINARY} ${FLAGS}
	GOOS=linux GOARCH=${ARCH32_arm} go build -o bin/linux/${ARCH32_arm}/${BINARY} ${FLAGS}
	GOOS=linux GOARCH=${ARCH64_amd64} go build -o bin/linux/${ARCH64_amd64}/${BINARY} ${FLAGS}
	GOOS=linux GOARCH=${ARCH64_arm64} go build -o bin/linux/${ARCH64_arm64}/${BINARY} ${FLAGS}

darwin:
	GOOS=darwin GOARCH=${ARCH32_386} go build -o bin/darwin/${ARCH32_386}/${BINARY} ${FLAGS}
	GOOS=darwin GOARCH=${ARCH64_amd64} go build -o bin/darwin/${ARCH64_amd64}/${BINARY} ${FLAGS}

windows:
	GOOS=windows GOARCH=${ARCH32_386} go build -o bin/windows/${ARCH32_386}/${BINARY}.exe ${FLAGS}
	GOOS=windows GOARCH=${ARCH64_amd64} go build -o bin/windows/${ARCH64_amd64}/${BINARY}.exe ${FLAGS}

freebsd:
	GOOS=freebsd GOARCH=${ARCH32_386} go build -o bin/freebsd/${ARCH32_386}/${BINARY}.out ${FLAGS}
	GOOS=freebsd GOARCH=${ARCH32_arm} go build -o bin/freebsd/${ARCH32_arm}/${BINARY}.out ${FLAGS}
	GOOS=freebsd GOARCH=${ARCH64_amd64} go build -o bin/freebsd/${ARCH64_amd64}/${BINARY}.out ${FLAGS}

solaris:
   GOOS=solaris GOARCH=${ARCH64_amd64} go build -o bin/solaris/${ARCH64_amd64}/${BINARY}.bin ${FLAGS}

test:
	go test ./test/...

clean:
	-rm -rf bin/

.PHONY: linux darwin windows freebsd solaris test clean
