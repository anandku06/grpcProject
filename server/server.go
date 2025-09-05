package server

import (
	"context"

	"github.com/anandku06/grpcProject/db"
	pb "github.com/anandku06/grpcProject/grpc_mongo_crud/proto"
	"github.com/anandku06/grpcProject/models"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer // gives all the unimplemented servies defined in the proto file
}

// createUser method
func (s *UserServer) CreateUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	user := models.User{
		UserId:   req.UserId,
		Name:     req.Name,
		Age:      req.Age,
		Email:    req.Email,
		Password: req.Password,
	}

	_, err := db.UserCollections.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{
		UserId: user.UserId,
		Name:   user.Name,
		Age:    user.Age,
		Email:  user.Email,
	}, nil
}
