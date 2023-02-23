#! bash
env GOOS=js GOARCH=wasm go build -o game.wasm &&  go run ./server
# nats-server --config nats-server.conf &
