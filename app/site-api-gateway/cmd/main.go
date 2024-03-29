package main

import (
	"github.com/gin-gonic/gin"
	"site/app/site-api-gateway/pkg/auth"
	"site/app/site-api-gateway/pkg/collaborate"
	"site/app/site-api-gateway/pkg/user"
	"site/common/inits"
	"site/conf"
)

func main() {
	inits.AppInitialize()
	gateConf := conf.GlobalAppConfig.GetServiceConf("gate")
	authConf := conf.GlobalAppConfig.GetServiceConf("auth")
	userConf := conf.GlobalAppConfig.GetServiceConf("user")
	teamConf := conf.GlobalAppConfig.GetServiceConf("team")
	gateConf.InitRunningServerName()
	inits.ToolInitialize()
	r := gin.Default()
	authSvc := auth.RegisterRoutes(r, authConf)
	user.RegisterRoutes(r, userConf, authSvc)
	collaborate.RegisterRoutes(r, teamConf, authSvc)
	err := r.Run(gateConf.GetAddress())
	if err != nil {
		panic(err)
	}
}
