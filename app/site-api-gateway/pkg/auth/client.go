package auth

import (
	"google.golang.org/grpc"
	"log"
	"site/conf"
	pb "site/protocol/auth"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

// InitServiceClient 初始化客户端
func InitServiceClient(c *conf.ServiceConf) pb.AuthServiceClient {
	cc, err := grpc.Dial(c.GetAddress(), grpc.WithInsecure())
	if err != nil {
		log.Printf("Could't connect: %v", err)
	}

	return pb.NewAuthServiceClient(cc)
}
