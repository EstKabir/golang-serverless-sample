PKG=./src
VERSION := $(shell git describe --always --long)
OUT=build/app.v.${VERSION}
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/)
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)

run_build:
	./${OUT}
build_app:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ${OUT} -ldflags="-X main.version=${VERSION}" ${PKG}
run:
	go run ${PKG}
test:
	@go test -short ${PKG_LIST}
lint:
	@for file in ${GO_FILES} ;  do \
		golint $$file ; \
	done
clean:
	go clean
	rm ${OUT}