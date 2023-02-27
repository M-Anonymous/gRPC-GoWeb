package service

import (
	"context"
	"log"
	"study.com/study-user/internal/server"
	pb "study.com/study-user/pkg/pb.service"
)

func init() {
	pb.RegisterUserServiceServer(server.GrpcServer, &UserService{})
}

type UserService struct {
	*pb.UnimplementedUserServiceServer
}

func (*UserService) SayHello(context.Context, *pb.Request) (*pb.Response, error) {
	log.Printf("Hello,I'm UserService!")
	return &pb.Response{
		Response: "Hello,I'm UserService!",
	}, nil
}
