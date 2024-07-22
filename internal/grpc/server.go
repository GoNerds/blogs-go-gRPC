package grpc

import (
	"context"

	pb "blogs-go-grpc/internal/grpc/proto"

	"go.mongodb.org/mongo-driver/mongo"
)

type BlogServer struct {
	pb.UnimplementedBlogServiceServer
	dbClient *mongo.Client
}

func NewBlogServer(dbClient *mongo.Client) *BlogServer {
	return &BlogServer{dbClient: dbClient}
}

func (s *BlogServer) CreateBlog(ctx context.Context, req *pb.CreateBlogRequest) (*pb.CreateBlogResponse, error) {
	return &pb.CreateBlogResponse{}, nil
}

