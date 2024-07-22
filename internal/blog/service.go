package blog

import (
	"context"
)

type Service interface{
	CreateBlog(ctx context.Context, blog *Blog) error
}

type BlogService struct{
	repo Repository
}

func NewBlogService(repo Repository) *BlogService {
	return &BlogService{repo: repo}
}

func (s *BlogService) CreateBlog(ctx context.Context, blog *Blog) error {
	return s.repo.Create(ctx, blog)
}