PHONY: run

run:
	go run .

build-wasm:
	env GOOS=js GOARCH=wasm go build -o ./docs/yourgame.wasm .
