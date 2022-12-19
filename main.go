package main

import (
	"fmt"
	"user_srv/handler"
	"user_srv/lib/helper"
	pb "user_srv/proto/go/user"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"github.com/spf13/cast"
)

func main() {
	// Create service
	address := "0.0.0.0:8261"
	consulHost := "192.168.107.151"
	consulPort := "8500"
	consulAddr := consulHost + ":" + cast.ToString(consulPort)

	// 读取配置
	config, err := helper.GetConsulConfig(consulHost, consulPort, "aplum")
	if err != nil {
		panic(err)
	}
	mysqlConfig := &helper.MysqlConfig{}
	helper.GetMySQLConsulConfig(config, &mysqlConfig, "mysql")

	// 启动服务
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
