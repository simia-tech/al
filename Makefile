
DOCKER ?= docker

wasm_path := build/test.wasm
package := github.com/simia-tech/al

clean:
	rm -f $(wasm_path)

$(wasm_path):
	$(DOCKER) run --rm \
	  -v "$(GOPATH)":/go \
		-v "$(PWD)":/target \
		-e GOOS=js -e GOARCH=wasm \
		golang:1.11-rc \
		go test -v -c -o /target/$(wasm_path) $(package)/storage/local
