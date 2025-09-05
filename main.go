package main

import (
	"fmt"
	"log"
	"net"

	"github.com/anandku06/grpcProject/db"
	pb "github.com/anandku06/grpcProject/grpc_mongo_crud/proto"
	"github.com/anandku06/grpcProject/server"
	"google.golang.org/grpc"
)


func main() {
	err := db.InitMango()

	if err != nil {
		log.Fatal("Mongo connection failed!", err)
	}

	lis, err := net.Listen("tcp", ":50051") // Listen announces on the local network address. The network must be "tcp", "tcp4", "tcp6", "unix" or "unixpacket". 50051 is default port number

	if err != nil {
		log.Fatal("Failed to listen ", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &server.UserServer{})
	fmt.Println("gRPC server is running!!")

	// Serve accepts incoming connections on the listener lis, creating a new ServerTransport and service goroutine for each.
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to serve the gRPC server!!", err)
	}
}