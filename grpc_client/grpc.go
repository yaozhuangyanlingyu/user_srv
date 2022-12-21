package grpc_client

import (
	user_proto "user_srv/proto/go/user"

	"github.com/micro/go-micro/v2/client"
)

var Grpc *GrpcClient

type GrpcClient struct {
	UserSrv user_proto.UserService
}

func SetClient(cli client.Client) {
	Grpc = &GrpcClient{
		UserSrv: user_proto.NewUserService("user_srv", cli),
	}
}
