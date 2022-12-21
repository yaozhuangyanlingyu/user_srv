package api_handler

import (
	"net/http"
	"user_srv/grpc_client"

	user_proto "user_srv/proto/go/user"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
)

func UserInfo(c *gin.Context) {
	// 调用服务端，获取数据
	data, err := _GetData(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 50000,
			"msg":  err.Error(),
		})
		return
	}

	// 调用rpc接口
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": data,
	})
}

/**
 * 熔断代码，获取用户数据
 */
func _GetData(c *gin.Context) (*user_proto.GetUserInfoResponse, error) {
	// 第一步：配置config
	configA := hystrix.CommandConfig{
		Timeout: 1000,
	}

	// 第二部：配置command
	hystrix.ConfigureCommand("get_user_info", configA)

	// 第三步：执行Do方法
	var userRsp *user_proto.GetUserInfoResponse
	err := hystrix.Do("get_user_info", func() (err error) {
		userRsp, err = grpc_client.Grpc.UserSrv.GetUserInfo(c, &user_proto.GetUserInfoRequest{})
		return err
	}, nil)
	return userRsp, err
}
