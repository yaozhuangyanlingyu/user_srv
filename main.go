package main

import (
	"fmt"
	"user_srv/handler"
	pb "user_srv/proto/go/user"

	"github.com/micro/go-micro/v2"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name("user"),
	)

	// Register handler
	pb.RegisterUserHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		fmt.Println(err)
		//logger.Fatal(err)
	}
}
