package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"site/app/site-auth-svc/pkg/client"
	"site/app/site-auth-svc/pkg/server"
	"site/app/site-auth-svc/pkg/service"
	"site/app/site-auth-svc/pkg/utils"
	"site/conf"
	pb "site/protocol/auth"
)

func main() {
	appConf, err := conf.LoadAppConf()
	if err != nil {
		panic(err)
	}

	err = conf.LoadParamConf()
	if err != nil {
		panic(err)
	}

	c := appConf.GetServiceConf("auth")

	// 服务监听客户端连接端口
	l, err := net.Listen("tcp", c.GetAddress())
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	userConf := appConf.GetServiceConf("user")
	userSvc := client.InitUserServiceClient(userConf.GetAddress())

	log.Printf("Auth_server port: %v", c.Port)

	jwt, err := utils.InitJwt()
	if err != nil {
		log.Fatalf("Failed to initialize JWT: %v", err)
	}

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
