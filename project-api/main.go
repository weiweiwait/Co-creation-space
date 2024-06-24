package main

import (
	"github.com/gin-gonic/gin"
	_ "my_project/project-api/api"
	"my_project/project-api/api/midd"
	"my_project/project-api/config"
	"my_project/project-api/router"
	srv "my_project/project-common"
)

func main() {
	r := gin.Default()
	r.Use(midd.RequestLog())
	//路由
	router.InitRouter(r)

	srv.Run(r, config.C.SC.Name, config.C.SC.Addr, nil)
}
