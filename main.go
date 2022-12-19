package main

import (
	"fmt"
	"user_srv/handler"
	pb "user_srv/proto/go/user"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
)

func main() {
	// Create service
	address := "0.0.0.0:8261"
	consulAddr := "192.168.107.151:8500"
	srv := micro.NewService(
		micro.Name("user_srv"),
		micro.Address(address),
		micro.Registry(consul.NewRegistry(func(o *registry.Options) {
			o.Addrs = []string{consulAddr}
		})),
	)

	// Register handler
	pb.RegisterUserHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		fmt.Println(err)
		//logger.Fatal(err)
	}
}
