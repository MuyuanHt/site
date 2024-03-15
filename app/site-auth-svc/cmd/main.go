package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"site/app/site-auth-svc/pkg/client"
	"site/app/site-auth-svc/pkg/server"
	"site/app/site-auth-svc/pkg/service"
	"site/app/site-auth-svc/pkg/utils"
	"site/common/inits"
	"site/conf"
	pb "site/protocol/auth"
)

func main() {
	inits.AppInitialize()
	authConf := conf.GlobalAppConfig.GetServiceConf("auth")
	userConf := conf.GlobalAppConfig.GetServiceConf("user")
	authConf.InitRunningServerName()
	inits.ToolInitialize()

	// 服务监听客户端连接端口
	l, err := net.Listen("tcp", authConf.GetAddress())
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Printf("Auth_server port: %v", authConf.Port)

	// 初始化其他服务客户端
	userSvc := client.InitUserServiceClient(userConf.GetAddress())
	jwt, err := utils.InitJwt()
	if err != nil {
		log.Fatalf("Failed to initialize JWT: %v", err)
	}

	// 注册rpc服务
	s := &server.AuthServer{
		Svc: service.AuthService{
			UserSvc: userSvc,
			Jwt:     jwt,
		},
	}
	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, s)
	if err = grpcServer.Serve(l); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
