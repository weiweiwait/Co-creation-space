package domain

import (
	"context"
	"my_project/project-common/errs"
	"my_project/project-project/internal/dao"
	"my_project/project-project/internal/data"
	"my_project/project-project/internal/repo"
	"my_project/project-project/pkg/model"
	"time"
)

type MenuDomain struct {
	menuRepo repo.MenuRepo
}

func (d *MenuDomain) MenuTreeList() ([]*data.ProjectMenuChild, *errs.BError) {
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	menus, err := d.menuRepo.FindMenus(c)
	if err != nil {
		return nil, model.DBError
	}
	menuChildren := data.CovertChild(menus)
	return menuChildren, nil
}

func NewMenuDomain() *MenuDomain {
	return &MenuDomain{
		menuRepo: dao.NewMenuDao(),
	}
}
