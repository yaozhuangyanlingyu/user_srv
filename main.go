package main

import (
	"fmt"
	"user_srv/api_handler"
	"user_srv/grpc_client"
	"user_srv/handler"
	"user_srv/lib/helper"
	pb "user_srv/proto/go/user"
	"user_srv/wrappers"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/web"
	"github.com/micro/go-plugins/registry/consul/v2"
	limiter "github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2"
	"github.com/spf13/cast"
)

const (
	QPS = 1
)

func main() {
	// 服务配置
	address := "0.0.0.0:8001"
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

	// http服务
	RunHttp(consulAddr)

	// 启动服务
	srv := micro.NewService(
		// 服务名
		micro.Name("user_srv"),

		// 服务地址
		micro.Address(address),

		// 服务注册consul配置
		micro.Registry(consul.NewRegistry(func(o *registry.Options) {
			o.Addrs = []string{consulAddr}
		})),

		// 限流配置
		micro.WrapHandler(limiter.NewHandlerWrapper(QPS)),

		// 熔断配置
		micro.WrapClient(wrappers.NewProductWrapper),
	)

	// Register handler
	pb.RegisterUserHandler(srv.Server(), handler.New())

	// 设置grpc客户端
	grpc_client.SetClient(srv.Client())

	// Run service
	if err := srv.Run(); err != nil {
		fmt.Println(err)
		//logger.Fatal(err)
	}
}

/**
 * 启动http服务
 */
func RunHttp(consulAddr string) {
	// 原生路由实现http
	/*
		go func() {
			server := web.NewService(
				web.Name("web-api"),
				web.Address(":8002"),
			)
			server.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
				writer.Write([]byte("hello world"))
			})
			server.Run()
		}()*/

	// gin框架实现http
	go func() {
		consulReg := consul.NewRegistry(func(o *registry.Options) {
			o.Addrs = []string{consulAddr}
		})
		ginRouter := gin.Default()

		// 获取用户信息
		ginRouter.Handle("GET", "/user", api_handler.UserInfo)

		// 获取商品列表数据
		ginRouter.Handle("GET", "/product/list", api_handler.GetProductList)
		server := web.NewService(
			web.Name("gin-api"),
			web.Address(":8003"),
			web.Handler(ginRouter),
			web.Registry(consulReg),
		)

		// 注：这里可以使用命令行参数，启动端口，例：go run main.go server_address :8004
		server.Init()
		server.Run()
	}()
}
