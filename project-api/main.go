package main

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	_ "my_project/project-api/api"
	"my_project/project-api/api/midd"
	"my_project/project-api/config"
	"my_project/project-api/router"
	srv "my_project/project-common"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Use(midd.RequestLog())
	r.StaticFS("/upload", http.Dir("upload"))
	//路由
	router.InitRouter(r)
	//开启pprof 默认的访问路径是/debug/pprof
	pprof.Register(r)

	srv.Run(r, config.C.SC.Name, config.C.SC.Addr, nil)
}
