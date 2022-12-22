package service

import (
	"user_srv/grpc_client"

	"user_srv/proto/go/user"
	user_proto "user_srv/proto/go/user"

	"github.com/gin-gonic/gin"
)

/**
 * 获取商品列表数据
 * @ 注：http调用
 */
func GetProductListWrapper(c *gin.Context) (*user_proto.GetProductListResponse, error) {
	// 正常请求接口
	productsRsp, err := grpc_client.Grpc.UserSrv.GetProductList(c, &user_proto.GetProductListRequest{})
	return productsRsp, err
}

/**
 * 获取默认数据
 */
func GetDefaultProductWrapper(rsp interface{}) {
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

	// 返回数据
	rspObj, ok := rsp.(*user_proto.GetProductListResponse)
	if ok {
		rspObj.Products = products
	}
}
