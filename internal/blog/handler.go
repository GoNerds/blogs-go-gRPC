package blog

import (
	"context"

	pb "blogs-go-grpc/internal/grpc/proto"
)

type BlogHandler struct {
	service Service
}

func NewBlogHandler(service Service) *BlogHandler {
	return &BlogHandler{service: service}
}

func (h *BlogHandler) CreateBlog(ctx context.Context, req *pb.CreateBlogRequest) (*pb.CreateBlogResponse, error) {
	blog := &Blog{
		Title: req.GetTitle(),
		Content: req.GetContent(),
		Author: req.GetAuthor(),
	}

	if err := h.service.CreateBlog(ctx, blog); err !- nil {
		return nil, err
	}

	return &pb.CreateBlogResponse{Id: blog.ID}, nil
}
