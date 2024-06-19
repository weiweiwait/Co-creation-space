package project

import (
	"context"
	"github.com/gin-gonic/gin"
	common "my_project/project-common"
	"my_project/project-common/errs"
	"my_project/project-grpc/project"
	"net/http"
	"time"
)

type HandlerProject struct {
}

func (p *HandlerProject) index(c *gin.Context) {
	result := &common.Result{}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	msg := &project.IndexMessage{}
	indexResponse, err := ProjectServiceClient.Index(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	c.JSON(http.StatusOK, result.Success(indexResponse.Menus))
}
func New() *HandlerProject {
	return &HandlerProject{}
}
