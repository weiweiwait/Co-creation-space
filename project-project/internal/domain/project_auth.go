package domain

import (
	"context"
	"go.uber.org/zap"
	"my_project/project-common/errs"
	"my_project/project-project/internal/dao"
	"my_project/project-project/internal/data"
	"my_project/project-project/internal/repo"
	"my_project/project-project/pkg/model"
	"time"
)

type ProjectAuthDomain struct {
	projectAuthRepo repo.ProjectAuthRepo
	userRpcDomain   *UserRpcDomain
}

func (d *ProjectAuthDomain) AuthList(orgCode int64) ([]*data.ProjectAuthDisplay, *errs.BError) {
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	list, err := d.projectAuthRepo.FindAuthList(c, orgCode)
	if err != nil {
		zap.L().Error("project AuthList projectAuthRepo.FindAuthList error", zap.Error(err))
		return nil, model.DBError
	}
	var pdList []*data.ProjectAuthDisplay
	for _, v := range list {
		display := v.ToDisplay()
		pdList = append(pdList, display)
	}
	return pdList, nil
}

func (d *ProjectAuthDomain) AuthListPage(orgCode int64, page int64, pageSize int64) ([]*data.ProjectAuthDisplay, int64, *errs.BError) {
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	list, total, err := d.projectAuthRepo.FindAuthListPage(c, orgCode, page, pageSize)
	if err != nil {
		zap.L().Error("project AuthList projectAuthRepo.FindAuthList error", zap.Error(err))
		return nil, 0, model.DBError
	}
	var pdList []*data.ProjectAuthDisplay
	for _, v := range list {
		display := v.ToDisplay()
		pdList = append(pdList, display)
	}
	return pdList, total, nil
}

func NewProjectAuthDomain1() *ProjectAuthDomain {
	return &ProjectAuthDomain{
		projectAuthRepo: dao.NewProjectAuthDao(),
		userRpcDomain:   NewUserRpcDomain(),
	}
}
