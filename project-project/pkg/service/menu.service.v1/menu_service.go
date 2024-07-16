package menu_service_v1

import (
	"context"
	"github.com/jinzhu/copier"
	"my_project/project-common/errs"
	"my_project/project-grpc/menu"
	"my_project/project-project/internal/dao"
	"my_project/project-project/internal/database/tran"
	"my_project/project-project/internal/domain"
	"my_project/project-project/internal/repo"
)

type MenuService struct {
	menu.UnimplementedMenuServiceServer
	cache       repo.Cache
	transaction tran.Transaction
	menuDomain  *domain.MenuDomain
}

func New() *MenuService {
	return &MenuService{
		cache:       dao.Rc,
		transaction: dao.NewTransaction(),
		menuDomain:  domain.NewMenuDomain(),
	}
}

func (d *MenuService) MenuList(ctx context.Context, msg *menu.MenuReqMessage) (*menu.MenuResponseMessage, error) {
	list, err := d.menuDomain.MenuTreeList()
	if err != nil {
		return nil, errs.GrpcError(err)
	}
	var mList []*menu.MenuMessage
	copier.Copy(&mList, list)
	return &menu.MenuResponseMessage{List: mList}, nil
}
