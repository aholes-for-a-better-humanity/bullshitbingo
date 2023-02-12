#! bash

env GOOS=js GOARCH=wasm go build -o game.wasm
ls -lh game.wasm
[[ 15951600 -gt  $( wc -c game.wasm | grep -o '[0-9]\+' ) ]] && echo OK
