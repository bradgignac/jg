BINARY=jg
VERSION=1.0.0
REVISION=`git rev-parse HEAD`
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Revision=${REVISION}"

default: dependencies build

dependencies: glide.yaml
	glide install

build:
	go build ${LDFLAGS} -o bin/${BINARY}

install: build
	cp bin/${BINARY} ${GOPATH}/bin/${BINARY}

clean:
	rm -r bin

.PHONY: clean install
