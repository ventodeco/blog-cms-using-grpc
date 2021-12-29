package main

import (
	"blog-grpc/blogpb"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	fmt.Println("Blog Client...")

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	cc, err := grpc.Dial("localhost:50051", opts...)

	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	// close connection when all logic executed
	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)

	fmt.Println("Creating a blog")

	blog := &blogpb.Blog{
		AuthorId: "tes",
		Content:  "tes",
		Title:    "Tess",
	}

	created, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})

	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}

	fmt.Printf("Blog has been created: %v", created)
}
