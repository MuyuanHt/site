package collaborate

import (
	"google.golang.org/grpc"
	"log"
	"site/conf"
	pb "site/protocol/collaborate"
)

type ServiceClient struct {
	Client pb.CollaborateServiceClient
}

// InitServiceClient 初始化客户端
func InitServiceClient(c *conf.ServiceConf) pb.CollaborateServiceClient {
	cc, err := grpc.Dial(c.GetAddress(), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could't connect: %v", err)
	}

	return pb.NewCollaborateServiceClient(cc)
}
