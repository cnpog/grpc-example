package server

import (
	"context"

	"github.com/cnpog/grpc-example/pkg/grpc_blog"
)

type Server struct {
	grpc_blog.UnimplementedBlogServer
}

func (s Server) CreateBlogPost(ctx context.Context, request *grpc_blog.BlogPostRequest) (*grpc_blog.BlogPostResponse, error) {
	r := &grpc_blog.BlogPostResponse{
		Username:      request.Username,
		Title:         request.Title,
		Content:       request.Content,
		Unixtimestamp: request.Unixtimestamp,
	}

	return r, nil
}
