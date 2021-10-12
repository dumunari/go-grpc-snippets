# Go gRPC Snippets
Simple project created to learn more about some gRPC concepts

## Running

### Setting some things up
Firstly, install:

* ```go install google.golang.org/protobuf/cmd/protoc-gen-go```

Then run:
* ```protoc --proto_path=proto proto/*.proto --go_out=pb```
* ```protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb```

### Running the server
* ```go run cmd/server/server.go```

### Running the clients
* Unary RPC Client
  * Client sends a single request and gets a single response from the server.
    * ```go run cmd/clients/unary/client.go ```

* Client Stream RPC Client
  * Client sends a sequence of requests using a stream and gets a single response from the server.
    * ```go run cmd/clients/client-stream/client.go ```

* Server Stream RPC Client
  * Client sends a single request and gets a stream to read a sequence of responses from the server.
    * ```go run cmd/clients/server-stream/client.go ```

* Bidirectional RPC Stream Client
  * Client and server send a sequence of messages using a read-write stream.
    * ```go run cmd/clients/bidirectional-stream/client.go```



