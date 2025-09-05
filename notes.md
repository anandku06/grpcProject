# gRPC + GO

## gRPC
- Google Remote Procedure Call
- a powerful open-source tool made by Google.
- helps different applications or services talk to each other - fast, securely, and in any language.
- allows a client application to directly call a method on a server application on a different machine as if it were a local object
- achieved through a contact-first approach, where the service interface  and the structure of the payload messages are defined in a `.proto` file using **Protocol Buffers**.

### Install gRPC in Go
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

### Features of gRPC
- Uses **Protocol Buffers (ProtoBuf)** as the Interface Definition Language (IDL)
- Runs on **HTTPS/2** (supports multiplexing, streaming, flow control, header compression, etc.)
- Supports **multiple languages** (Go, Java, Python, C++, Rust, etc.)
- Provides Bi-directional streaming (client -> server can send multiple messages)
- Offers built-in features like **authentication**, **load balancing**, and **deadline propagation**.

### Types of RPC in gRPC
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
- `.proto` file = **Protocol Buffers** definition file
-  Its where you define:
    - **Messages** (the data structure you send & recieve)
    - **Services** (the RPC methods your server provide)
- Think of it like a **contract** between client and server.
    - Both agree to use this file, so they have always to understand each other's data and functions

### 1. Why do we need `.proto` file?
- Serialization → Converts data into a small, fast binary format (much faster than JSON).
- Language Neutral → Works with Go, Python, Java, C++, etc.
- Strongly Typed → Compiler checks for errors.
- Single Source of Truth → Instead of writing data structures manually in each language, you define once in .proto, then generate code.

### Structure of a `.proto` File
```proto
syntax = "proto3";    // Version of Protobuf

package user;         // Namespace

// Define a service (like an interface)
service UserService {
  rpc GetUser (UserRequest) returns (UserResponse);
}

// Define the request message
message UserRequest {
  int32 id = 1;        // Field type: int32, field name: id, field tag: 1
}

// Define the response message
message UserResponse {
  int32 id = 1;
  string name = 2;
}
```
- ` syntax = "proto3";` -> Protobuf has versions: proto2 and proto3. gRPC usually uses proto3 because it’s simpler.

- `package user;` -> Like Go/Java packages → helps organize code.

- `service` -> Defines RPC methods that server exposes.
```proto
// Example:
service UserService {
  rpc GetUser (UserRequest) returns (UserResponse);
}
```
- `rpc GetUser` = Remote Procedure Call.
- Input: `UserRequest`.
- Output: `UserResponse`.

- `message` -> Defines the data structure (like a Go struct).
```proto
// Example:
message UserRequest {
  int32 id = 1;
}
```

→ Same as:
```go
type UserRequest struct {
    Id int32
}
```
- Field numbers (= 1, = 2, …) -> Each field gets a unique number (called a tag). Tags are used in the binary encoding.
```proto
// Example:
message UserResponse {
  int32 id = 1;
  string name = 2;
}
```
→ Field id = 1, name = 2.
→ Even if names change later, numbers keep compatibility.

- This `.proto` File is compiled using `.proto` compiler
    - `user.pb.go` -> Go structs
    - `user_grpc.pb.go` -> Go interfaces for gRPC

## protoc

- `protoc` => **Protocol Buffer Compiler**
- write your *API definition* (message + services) inside the `.proto` file.
- `protoc` compiles that file into the source code, that can be read by Go (or any other language according to the dependencies used)
- Basically, generates the boilerplate code for you.

```bash
protoc --go_out=. --go-grpc_out=. user.proto # this is the command
```

- it generates two files:
  - `user.pb.go`
    - contains data structure (structs) for your message
    - In short, all your data models live here.
```proto
message User {
  string name = 1;
  int32 age = 2;
}
```
- then in Go, it looks like this:
```go
type User struct {
    Name string
    Age  int32
}
```

  - `user_grpc.pb.go`
    - contains **gRPC service code**.
    - all your gRPC networking boilerplate lives here.

```proto
service UserService {
  rpc GetUser(UserRequest) returns (UserResponse);
}
```

- Then in Go, it'll look like this:

```go
type UserServiceServer interface {
    GetUser(context.Context, *UserRequest) (*UserResponse, error)
}

// or

type UserServiceClient interface {
    GetUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
}

```

- refrain from any type of editing in these files

## pb.UnimplementedUserServiceServer

- got from the `user_grpc.pb.go`, contains all the unimplemented user services written in the `.proto` file
- all the services defined inside the `.proto` file gets an interface in the `user_grpc.pb.go` file, and a helper empty struct.

**What happens if you don't embed the line?**
- If tomorrow someone adds a new method in the `.proto`, your server would stop compiling because you didn’t implement that new method.
- By embedding `UnimplementedUserServiceServer`, your server automatically has a default "not implemented" version of any methods you didn’t write yet.
- That means your code keeps compiling, and you only override the methods you care about. 

- Think of `UnimplementedUserServiceServer` as a default skeleton house 🏠.
  - Even if you don’t build all rooms yet, you at least have a placeholder saying “room not ready”.You only decorate or add rooms (methods) you care about right now.