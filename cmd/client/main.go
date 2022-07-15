package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/cnpog/grpc-example/pkg/grpc_blog"

	"google.golang.org/grpc"
)

func main() {
	c, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can't connect to backend: %v", err)
	}

	client := grpc_blog.NewBlogClient(c)
	for {
		blogpost, err := client.CreateBlogPost(context.Background(), &grpc_blog.BlogPostRequest{
			Username:      "MaxMustermann",
			Title:         "Hello",
			Content:       "World",
			Unixtimestamp: time.Now().Unix(),
		})
		if err != nil {
			log.Fatalf("failed to create blogpost: %v", err)
		}

		fmt.Println(blogpost)
	}
}
