package user

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	common "my_project/project-common"
	"my_project/project-common/logs"
	"my_project/project-user/pkg/dao"
	"my_project/project-user/pkg/repo"
	"net/http"
	"time"
)

type HandlerUser struct {
	cache repo.Cache
}

func New() *HandlerUser {
	return &HandlerUser{
		cache: dao.Rc,
	}
}
func (*HandlerUser) getCaptcha(ctx *gin.Context) {
	rsp := &common.Result{}
	go func() {
		time.Sleep(2 * time.Second)
		zap.L().Info("成功 INFO")
		logs.LG.Debug("成功 DEBUG")
		zap.L().Error("成功 ERROR")
		log.Println("成功")

	}()
	ctx.JSON(http.StatusOK, rsp.Success("123456"))
}
