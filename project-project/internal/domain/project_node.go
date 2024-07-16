package domain

import (
	"context"
	"my_project/project-common/errs"
	"my_project/project-project/internal/dao"
	"my_project/project-project/internal/data"
	"my_project/project-project/internal/repo"
	"my_project/project-project/pkg/model"
)

type ProjectNodeDomain struct {
	projectNodeRepo repo.ProjectNodeRepo
}

func (d *ProjectNodeDomain) TreeList() ([]*data.ProjectNodeTree, *errs.BError) {
	//node表都查出来 转换成treelist结构
	list, err := d.projectNodeRepo.FindAll(context.Background())
	if err != nil {
		return nil, model.DBError
	}
	return data.ToNodeTreeList(list), nil
}

func (d *ProjectNodeDomain) NodeList() ([]*data.ProjectNode, *errs.BError) {
	list, err := d.projectNodeRepo.FindAll(context.Background())
	if err != nil {
		return nil, model.DBError
	}
	return list, nil
}

func NewProjectNodeDomain() *ProjectNodeDomain {
	return &ProjectNodeDomain{
		projectNodeRepo: dao.NewProjectNodeDao(),
	}
}
