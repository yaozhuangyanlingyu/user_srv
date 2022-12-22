package api_handler

import (
	"net/http"
	"user_srv/api_handler/service"

	"github.com/gin-gonic/gin"
)

func UserInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})
}

func GetProductList(c *gin.Context) {
	// 调用服务端，获取数据
	data, err := service.GetProductListWrapper(c)
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
