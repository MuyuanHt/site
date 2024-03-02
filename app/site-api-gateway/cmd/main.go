package main

import (
	"github.com/gin-gonic/gin"
	"site/app/site-api-gateway/pkg/auth"
	"site/app/site-api-gateway/pkg/user"
	"site/conf"
)

func main() {
	c, err := conf.LoadAppConf()
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	authSvc := auth.RegisterRoutes(r, c.GetServiceConf("auth"))
	user.RegisterRoutes(r, c.GetServiceConf("user"), authSvc)

	err = r.Run(c.GetServiceConf("gate").GetAddress())
	if err != nil {
		panic(err)
	}
}
