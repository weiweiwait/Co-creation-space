package auth_service_v1

import (
	"context"
	"github.com/jinzhu/copier"
	"my_project/project-common/encrypts"
	"my_project/project-common/errs"
	"my_project/project-grpc/auth"
	"my_project/project-project/internal/dao"
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
