package project_service_v1

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
func (p *ProjectService) Index(context.Context, *project.IndexMessage) (*project.IndexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method")
}
