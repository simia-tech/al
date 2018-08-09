
DOCKER ?= docker

GOPATH := $(shell go env GOPATH)

wasm_path := build/test.wasm
package := github.com/simia-tech/al

clean:
	$(DOCKER) run --rm \
		-v "$(PWD)":/target \
		golang:1.11-rc \
		rm -f /target/$(wasm_path)

$(wasm_path):
	$(DOCKER) run --rm \
	  -v "$(GOPATH)":/go \
		-v "$(PWD)":/target \
		-e GOOS=js -e GOARCH=wasm \
		golang:1.11-rc \
		go test -v -c -o /target/$(wasm_path) $(package)/http
