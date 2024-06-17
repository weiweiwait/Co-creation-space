package user

import (
	"github.com/gin-gonic/gin"
	"log"
	"my_project/project-user/router"
)

func init() {
	log.Println("init user router")
	router.Register(&RouterUser{})
}

type RouterUser struct {
}

func (*RouterUser) Route(r *gin.Engine) {
	h := &HandlerUser{}
	r.POST("/user/login", h.getCaptcha)
}
