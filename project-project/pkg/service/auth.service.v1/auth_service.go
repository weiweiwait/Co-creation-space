package auth_service_v1

import (
	"context"
	"github.com/jinzhu/copier"
	"my_project/project-common/encrypts"
	"my_project/project-common/errs"
	"my_project/project-grpc/auth"
	"my_project/project-project/internal/dao"
	"my_project/project-project/internal/database"
	"my_project/project-project/internal/database/tran"
	"my_project/project-project/internal/domain"
	"my_project/project-project/internal/repo"
)

type AuthService struct {
	auth.UnimplementedAuthServiceServer
	cache             repo.Cache
	transaction       tran.Transaction
	projectAuthDomain *domain.ProjectAuthDomain
}

func New() *AuthService {
	return &AuthService{
		cache:             dao.Rc,
		transaction:       dao.NewTransaction(),
		projectAuthDomain: domain.NewProjectAuthDomain1(),
	}
}
func (a *AuthService) AuthList(ctx context.Context, msg *auth.AuthReqMessage) (*auth.ListAuthMessage, error) {
	organizationCode := encrypts.DecryptNoErr(msg.OrganizationCode)
	listPage, total, err := a.projectAuthDomain.AuthListPage(organizationCode, msg.Page, msg.PageSize)
	if err != nil {
		return nil, errs.GrpcError(err)
	}
	var prList []*auth.ProjectAuth
	copier.Copy(&prList, listPage)
	return &auth.ListAuthMessage{List: prList, Total: total}, nil
}
func (a *AuthService) Apply(ctx context.Context, msg *auth.AuthReqMessage) (*auth.ApplyResponse, error) {
	if msg.Action == "getnode" {
		//获取列表
		list, checkedList, err := a.projectAuthDomain.AllNodeAndAuth(msg.AuthId)
		if err != nil {
			return nil, errs.GrpcError(err)
		}
		var prList []*auth.ProjectNodeMessage
		copier.Copy(&prList, list)
		return &auth.ApplyResponse{List: prList, CheckedList: checkedList}, nil
	}
	if msg.Action == "save" {
		//先删除 project_auth_node表 在新增  事务
		//保存
		nodes := msg.Nodes
		//先删在存 加事务
		authId := msg.AuthId
		err := a.transaction.Action(func(conn database.DbConn) error {
			err := a.projectAuthDomain.Save(conn, authId, nodes)
			return err
		})
		if err != nil {
			return nil, errs.GrpcError(err.(*errs.BError))
		}
	}
	return &auth.ApplyResponse{}, nil
}
