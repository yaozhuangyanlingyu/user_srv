package service

import (
	"user_srv/grpc_client"

	"user_srv/proto/go/user"
	user_proto "user_srv/proto/go/user"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
)

/**
 * 熔断代码，获取商品列表数据
 * 注：这是正常使用时的方法，这样太麻烦了，需要看框架中如何使用
 */
func GetProductList(c *gin.Context) (*user_proto.GetProductListResponse, error) {
	// 第一步：配置config
	configA := hystrix.CommandConfig{
		Timeout: 1000,
	}

	// 第二部：配置command
	hystrix.ConfigureCommand("get_product_list", configA)

	// 第三步：执行Do方法
	var productsRsp *user_proto.GetProductListResponse
	err := hystrix.Do("get_product_list", func() (err error) {
		// 正常请求接口
		productsRsp, err = grpc_client.Grpc.UserSrv.GetProductList(c, &user_proto.GetProductListRequest{})
		return err
	}, func(err error) error {
		// 熔断降级
		tmpProductsRsp, tmpErr := GetDefaultProduct(c)
		if tmpErr != nil {
			return tmpErr
		}
		productsRsp = tmpProductsRsp
		return nil
	})
	return productsRsp, err
}

/**
 * 获取默认数据
 */
func GetDefaultProduct(c *gin.Context) (*user_proto.GetProductListResponse, error) {
	rsp := &user_proto.GetProductListResponse{}
	products := make([]*user.ProductListRow, 0)
	products = append(products, &user.ProductListRow{
		Id:   101,
		Name: "默认-香奈儿手提包",
	})
	products = append(products, &user.ProductListRow{
		Id:   102,
		Name: "默认-lv单肩包",
	})
	products = append(products, &user.ProductListRow{
		Id:   103,
		Name: "默认-鳄鱼皮带",
	})
	products = append(products, &user.ProductListRow{
		Id:   104,
		Name: "默认-爱马仕双肩包",
	})
	products = append(products, &user.ProductListRow{
		Id:   105,
		Name: "默认-江诗丹顿手表",
	})
	rsp.Products = products
	return rsp, nil
}
