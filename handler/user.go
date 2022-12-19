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
