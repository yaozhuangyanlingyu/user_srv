package wrappers

import (
	"context"

	apiService "user_srv/api_handler/service"

	user_proto "user_srv/proto/go/user"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/v2/client"
)

type ProductWrapper struct {
	client.Client
}

func NewProductWrapper(c client.Client) client.Client {
	return &ProductWrapper{
		c,
	}
}

func (c *ProductWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	cmdName := req.Service() + "." + req.Endpoint()

	// 第一步：配置config
	configA := hystrix.CommandConfig{
		MaxConcurrentRequests:  1,    // 并发量超过多少触发
		Timeout:                1000, // 接口超时时间
		RequestVolumeThreshold: 2,    // 只要有2个请求过来，开始判断错误情况
		ErrorPercentThreshold:  50,   // 错误超过50%，开启熔断降级
		SleepWindow:            3000, // 降级方法，响应时长。超过该时间后，再次探测响应是否恢复正常
	}

	// 第二步：配置command
	hystrix.ConfigureCommand(cmdName, configA)

	// 第三步：请求接口
	return hystrix.Do(cmdName, func() error {
		return c.Client.Call(ctx, req, rsp, opts...)
	}, func(err error) error {
		// 需要降级方法
		DefaultData(rsp)
		return nil
	})
}

/**
 * 降级数据处理
 */
func DefaultData(rsp interface{}) error {
	switch rsp.(type) {
	case *user_proto.GetProductListResponse:
		apiService.GetDefaultProductWrapper(rsp)
	}
	return nil
}
