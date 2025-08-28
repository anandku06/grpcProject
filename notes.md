# gRPC + GO

## gRPC
    -
- Google Remote Procedure Call
- a powerful open-source tool made by Google.
- helps different applications or services talk to each other - fast, securely, and in any language.
- allows a client application to directly call a method on a server application on a different machine as if it were a local object
- achieved through a contact-first approach, where the service interface  and the structure of the payload messages are defined in a `.proto` file using **Protocol Buffers**.

### Install gRPC in Go
    - 
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

### Features of gRPC
    -
- Uses **Protocol Buffers (ProtoBuf)** as the Interface Definition Language (IDL)
- Runs on **HTTPS/2** (supports multiplexing, streaming, flow control, header compression, etc.)
- Supports **multiple languages** (Go, Java, Python, C++, Rust, etc.)
- Provides Bi-directional streaming (client -> server can send multiple messages)
- Offers built-in features like **authentication**, **load balancing**, and **deadline propagation**.

### Types of RPC in gRPC
    - 
- gRPC supports 4 types of communication patterns:
1. **Unary RPC**
    -> Client sends **one request**, gets **one response** (most common)
2. **Server Streaming ROC**
    -> Client sends one request, **server stream mutiple responses**. like fetching a list of records
3. **Client Streaming RPC**
    -> Client streams multiple requests, server responds with **one final response**. like uploading chunks of a file.
4. **Bi-directional Streaming RPC**
    -> Both client & server stream messages independently. like chat app

## .proto File
    -
- `.proto` file = **Protocol Buffers** definition file
-  Its where you define:
    - **Messages** (the data structure you send & recieve)
    - **Services** (the RPC methods your server provide)
- Think of it like a **contract** between client and server.