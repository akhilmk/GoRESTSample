To generate the proto files

**protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative  proto/routeguide.proto**

Run the server:

**$ go run server/server.go**

From another terminal, run the client:

**$ go run client/client.go**
