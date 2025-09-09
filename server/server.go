package server

import (
	"context"

	"github.com/anandku06/grpcProject/db"
	pb "github.com/anandku06/grpcProject/grpc_mongo_crud/proto"
	"github.com/anandku06/grpcProject/models"
	"go.mongodb.org/mongo-driver/bson"
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

func (s *UserServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	user := models.User{}
	err := db.UserCollections.FindOne(ctx, bson.M{"user_id": req.UserId}).Decode(&user)
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

func (s *UserServer) UpdateUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	filter := bson.M{"user_id": req.UserId}
	update := bson.M{"$set": bson.M{
		"name":     req.Name,
		"age":      req.Age,
		"email":    req.Email,
		"password": req.Password,
	}}

	_, err := db.UserCollections.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{
		UserId: req.UserId,
		Name:   req.Name,
		Age:    req.Age,
		Email:  req.Email,
	}, nil
}

func (s *UserServer) DeleteUser(ctx context.Context, req *pb.GetUserRequest) (*pb.DeleteResponse, error) {
	_, err := db.UserCollections.DeleteOne(ctx, bson.M{"user_id": req.UserId})
	if err != nil {
		return nil, err
	}
	return &pb.DeleteResponse{
		Mssg: "User deleted successfully",
	}, nil
}