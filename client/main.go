package main

import (
	"fmt"

	"github.com/asmcos/requests"
	"github.com/micro/go-micro/v2/client/selector"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"github.com/spf13/cast"
)

func main() {
	// 从服务中获取一个IP
	consulHost := "192.168.107.151"
	consulPort := "8500"
	consulAddr := consulHost + ":" + cast.ToString(consulPort)
	consulReg := consul.NewRegistry(func(o *registry.Options) {
		o.Addrs = []string{consulAddr}
	})
	getService, err := consulReg.GetService("gin-api")
	if err != nil {
		panic(err)
	}

	// 随机获取一个IP
	/*
		next := selector.Random(getService)
		node, err := next()
		if err != nil {
			panic(err)
		}*/

	// 轮训获取一个IP
	next := selector.RoundRobin(getService)
	node, err := next()
	if err != nil {
		panic(err)
	}

	// 调用服务代码
	fmt.Println(node.Id, node.Address, node.Metadata)
	Call_01(node.Address)
}

/**
 * 原始方式调用接口
 */
func Call_01(address string) {
	url := "http://" + address + "/user"
	resp, err := requests.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Text())
}
