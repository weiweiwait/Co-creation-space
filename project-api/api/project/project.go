package project

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"my_project/project-api/pkg/model"
	"my_project/project-api/pkg/model/menu"
	"my_project/project-api/pkg/model/pro"
	common "my_project/project-common"
	"my_project/project-common/errs"
	"my_project/project-grpc/project"
	"net/http"
	"strconv"
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
	}
	menus := indexResponse.Menus
	var ms []*menu.Menu
	copier.Copy(&ms, menus)
	c.JSON(http.StatusOK, result.Success(ms))
}
func (p *HandlerProject) myProjectList(c *gin.Context) {
	result := &common.Result{}
	//1. 获取参数
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	memberId := c.GetInt64("memberId")
	memberName := c.GetString("memberName")
	page := &model.Page{}
	page.Bind(c)
	selectBy := c.PostForm("selectBy")
	msg := &project.ProjectRpcMessage{
		MemberId:   memberId,
		MemberName: memberName,
		SelectBy:   selectBy,
		Page:       page.Page,
		PageSize:   page.PageSize}
	myProjectResponse, err := ProjectServiceClient.FindProjectByMemId(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
	}

	var pms []*pro.ProjectAndMember
	copier.Copy(&pms, myProjectResponse.Pm)
	if pms == nil {
		pms = []*pro.ProjectAndMember{}
	}
	c.JSON(http.StatusOK, result.Success(gin.H{
		"list":  pms, //null nil -> []
		"total": myProjectResponse.Total,
	}))
}
func (p *HandlerProject) projectTemplate(c *gin.Context) {
	result := &common.Result{}
	//1. 获取参数
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	memberId := c.GetInt64("memberId")
	memberName := c.GetString("memberName")
	page := &model.Page{}
	page.Bind(c)
	viewTypeStr := c.PostForm("viewType")
	viewType, _ := strconv.ParseInt(viewTypeStr, 10, 64)
	msg := &project.ProjectRpcMessage{
		MemberId:         memberId,
		MemberName:       memberName,
		ViewType:         int32(viewType),
		Page:             page.Page,
		PageSize:         page.PageSize,
		OrganizationCode: c.GetString("organizationCode")}
	templateResponse, err := ProjectServiceClient.FindProjectTemplate(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
	}

	var pms []*pro.ProjectTemplate
	copier.Copy(&pms, templateResponse.Ptm)
	if pms == nil {
		pms = []*pro.ProjectTemplate{}
	}
	for _, v := range pms {
		if v.TaskStages == nil {
			v.TaskStages = []*pro.TaskStagesOnlyName{}
		}
	}
	c.JSON(http.StatusOK, result.Success(gin.H{
		"list":  pms, //null nil -> []
		"total": templateResponse.Total,
	}))
}

func (p *HandlerProject) projectSave(c *gin.Context) {
	result := &common.Result{}
	//1. 获取参数
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	memberId := c.GetInt64("memberId")
	organizationCode := c.GetString("organizationCode")
	var req *pro.SaveProjectRequest
	c.ShouldBind(&req)
	msg := &project.ProjectRpcMessage{
		MemberId:         memberId,
		OrganizationCode: organizationCode,
		TemplateCode:     req.TemplateCode,
		Name:             req.Name,
		Id:               int64(req.Id),
		Description:      req.Description,
	}
	saveProject, err := ProjectServiceClient.SaveProject(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
	}
	var rsp *pro.SaveProject
	copier.Copy(&rsp, saveProject)
	c.JSON(http.StatusOK, result.Success(rsp))
}
func New() *HandlerProject {
	return &HandlerProject{}
}
