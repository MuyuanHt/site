package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"site/app/site-user-svc/pkg/dao"
	"site/app/site-user-svc/pkg/server"
	"site/app/site-user-svc/pkg/service"
	"site/app/site-user-svc/pkg/utils"
	"site/common/inits"
	"site/conf"
	pb "site/protocol/user"
)

func main() {
	inits.AppInitialize()
	userConf := conf.GlobalAppConfig.GetServiceConf("user")
	userConf.InitRunningServerName()
	inits.ToolInitialize()

	// 初始化数据库
	dsn := userConf.Mysql.GetMysqlDsn()
	d, err := dao.InitDao(dsn)
	if err != nil {
		log.Fatalf("Failed to inits dao: %v", err)
	}
	defer d.Close()

	// 初始化雪花ID
	sObj, err := utils.InitSnowflake(2, 2)
	if err != nil {
		log.Fatalf("Failed to inits snowflke: %v", err)
	}

	// 服务监听客户端连接端口
	l, err := net.Listen("tcp", userConf.GetAddress())
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Printf("User_server port: %v", userConf.Port)

	// 注册rpc服务
	s := &server.UserServer{
		Svc: service.InitUserService(d, sObj),
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, s)
	if err = grpcServer.Serve(l); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
