package client

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "site/protocol/user"
)

type UserServiceClient struct {
	Client pb.UserServiceClient
}

// InitUserServiceClient 初始化用户服务客户端
func InitUserServiceClient(url string) UserServiceClient {
	cc, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Printf("Could't connect: %v", err)
	}
	return UserServiceClient{
		Client: pb.NewUserServiceClient(cc),
	}
}

// FindAccount 查询账户
func (c *UserServiceClient) FindAccount(accountType int32, account string) (*pb.FindAccountResp, error) {
	req := &pb.FindAccountReq{
		AccountType: accountType,
		Account:     account,
	}
	return c.Client.FindAccount(context.Background(), req)
}

// CreateUser 创建用户
func (c *UserServiceClient) CreateUser(accountType int32, account string, password string) (*pb.CreateUserResp, error) {
	req := &pb.CreateUserReq{
		AccountType: accountType,
		Account:     account,
		Password:    password,
	}
	return c.Client.CreateUser(context.Background(), req)
}
