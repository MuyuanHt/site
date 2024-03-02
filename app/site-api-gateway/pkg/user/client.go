package user

import (
	"google.golang.org/grpc"
	"log"
	"site/conf"
	pb "site/protocol/user"
)

type ServiceClient struct {
	Client pb.UserServiceClient
}

// InitServiceClient 初始化客户端
func InitServiceClient(c *conf.ServiceConf) pb.UserServiceClient {
	cc, err := grpc.Dial(c.GetAddress(), grpc.WithInsecure())
	if err != nil {
		log.Printf("Could't connect: %v", err)
	}

	return pb.NewUserServiceClient(cc)
}
