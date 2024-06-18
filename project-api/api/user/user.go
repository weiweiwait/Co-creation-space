package user

import (
	"context"
	"github.com/gin-gonic/gin"
	common "my_project/project-common"
	loginServiceV1 "my_project/project-user/pkg/service/login.service.v1"
	"net/http"
	"time"
)

type HandlerUser struct {
}

func New() *HandlerUser {
	return &HandlerUser{}
}

func (*HandlerUser) getCaptcha(ctx *gin.Context) {
	result := &common.Result{}
	mobile := ctx.PostForm("mobile")
	//调用超时多长时间
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	rsp, err := LoginServiceClient.GetCaptcha(c, &loginServiceV1.CaptchaMessage{Mobile: mobile})
	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(2001, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, result.Success(rsp.Code))
}
