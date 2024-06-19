package project_service_v1

import (
	"my_project/project-grpc/project"
	"my_project/project-project/internal/dao"
	"my_project/project-project/internal/database/tran"
	"my_project/project-project/internal/repo"
)

type ProjectService struct {
	project.UnimplementedProjectServiceServer
	cache       repo.Cache
	transaction tran.Transaction
}

func New() *ProjectService {
	return &ProjectService{
		cache:       dao.Rc,
		transaction: dao.NewTransaction(),
	}
}
