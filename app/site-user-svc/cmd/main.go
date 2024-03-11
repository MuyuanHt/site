package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"site/app/site-user-svc/pkg/dao"
	"site/app/site-user-svc/pkg/server"
	"site/app/site-user-svc/pkg/service"
	"site/app/site-user-svc/pkg/utils"
	"site/conf"
	pb "site/protocol/user"
)

func main() {
	allConf, err := conf.LoadAppConf()
	if err != nil {
		panic(err)
	}

	c := allConf.GetServiceConf("user")
	err = conf.LoadParamConf()
	if err != nil {
		panic(err)
	}
	dsn := c.Mysql.GetMysqlDsn()
	d, err := dao.InitDao(dsn)
	if err != nil {
		log.Fatalf("Failed to init dao: %v", err)
	}

	sObj, err := utils.InitSnowflake(2, 2)
	if err != nil {
		log.Fatalf("Failed to init snowflke: %v", err)
	}

	// 服务监听客户端连接端口
	l, err := net.Listen("tcp", c.GetAddress())
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Printf("User_server port: %v", c.Port)

	s := &server.UserServer{
		Svc: service.InitUserService(d, sObj),
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, s)
	if err = grpcServer.Serve(l); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
