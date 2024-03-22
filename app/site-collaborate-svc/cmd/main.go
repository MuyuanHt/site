package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"site/app/site-collaborate-svc/pkg/client"
	"site/app/site-collaborate-svc/pkg/dao"
	"site/app/site-collaborate-svc/pkg/server"
	"site/app/site-collaborate-svc/pkg/service"
	"site/common/inits"
	"site/conf"
	pb "site/protocol/collaborate"
)

func main() {
	inits.AppInitialize()
	teamConf := conf.GlobalAppConfig.GetServiceConf("team")
	userConf := conf.GlobalAppConfig.GetServiceConf("user")
	teamConf.InitRunningServerName()
	inits.ToolInitialize()

	// 初始化数据库
	dsn := teamConf.Mysql.GetMysqlDsn()
	d, err := dao.InitDao(dsn)
	if err != nil {
		log.Fatalf("Failed to inits dao: %v", err)
	}
	defer d.Close()

	// 服务监听客户端连接端口
	l, err := net.Listen("tcp", teamConf.GetAddress())
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Printf("Team_server port: %v", teamConf.Port)

	// 初始化其他服务客户端
	userSvc := client.InitUserServiceClient(userConf.GetAddress())

	// 注册rpc服务
	s := &server.CollaborateServer{
		FriendSvc: service.InitFriendService(d, userSvc),
	}
	grpcServer := grpc.NewServer()
	pb.RegisterCollaborateServiceServer(grpcServer, s)
	if err = grpcServer.Serve(l); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
