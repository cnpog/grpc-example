package main

import (
	"log"
	"net"

	"github.com/cnpog/grpc-example/pkg/grpc_blog"
	"github.com/cnpog/grpc-example/pkg/server"

	"google.golang.org/grpc"
)

func main() {
	gs := grpc.NewServer()
	var s server.Server

	grpc_blog.RegisterBlogServer(gs, s)
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatalf("error starting sever: %v", err)
	}
	log.Fatal(gs.Serve(listener))
}
