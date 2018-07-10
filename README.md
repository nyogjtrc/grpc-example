# grpc-example

example for https://nyogjtrc.github.io/posts/2018/07/try-grpc-with-go/

- start server: `go run server/main.go`
- start client: `go run client/main.go`
- send http request: `curl -X POST http://localhost:9999/echo/echo -H 'Content-Type: application/json' -d '{ "value": "hi, how are you?"  }'`
