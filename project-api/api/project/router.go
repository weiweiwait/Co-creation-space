package project

import (
	"github.com/gin-gonic/gin"
	"log"
	"my_project/project-api/api/midd"
	"my_project/project-api/router"
)

type RouterProject struct {
}

func init() {
	log.Println("init project router")
	ru := &RouterProject{}
	router.Register(ru)
}

func (*RouterProject) Route(r *gin.Engine) {
	//初始化grpc的客户端连接
	InitRpcProjectClient()
	h := New()
	group := r.Group("/project/index")
	group.Use(midd.TokenVerify1())
	group.POST("", h.index)
	group1 := r.Group("/project/project")
	group1.Use(midd.TokenVerify1())
	group1.POST("/selfList", h.myProjectList)
}
