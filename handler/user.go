package handler

import (
	"context"

	user "user_srv/proto/go/user"
)

type User struct{}

func New() *User {
	return &User{}
}

func (e *User) GetUserInfo(ctx context.Context, req *user.GetUserInfoRequest, rsp *user.GetUserInfoResponse) error {
	rsp.UserId = 198
	rsp.Name = "姚晓丰"
	return nil
}

func (e *User) GetProductList(ctx context.Context, req *user.GetProductListRequest, rsp *user.GetProductListResponse) error {
	products := make([]*user.ProductListRow, 0)
	products = append(products, &user.ProductListRow{
		Id:   1,
		Name: "香奈儿手提包",
	})
	products = append(products, &user.ProductListRow{
		Id:   2,
		Name: "lv单肩包",
	})
	products = append(products, &user.ProductListRow{
		Id:   3,
		Name: "鳄鱼皮带",
	})
	products = append(products, &user.ProductListRow{
		Id:   4,
		Name: "爱马仕双肩包",
	})
	products = append(products, &user.ProductListRow{
		Id:   5,
		Name: "江诗丹顿手表",
	})
	products = append(products, &user.ProductListRow{
		Id:   6,
		Name: "阿迪达斯袜子",
	})
	products = append(products, &user.ProductListRow{
		Id:   7,
		Name: "鳄鱼皮鞋",
	})
	products = append(products, &user.ProductListRow{
		Id:   8,
		Name: "花花公子西服",
	})
	products = append(products, &user.ProductListRow{
		Id:   9,
		Name: "安踏运动鞋",
	})
	products = append(products, &user.ProductListRow{
		Id:   10,
		Name: "雪白大号连衣裙",
	})
	rsp.Products = products
	return nil
}
