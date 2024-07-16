package repo

import (
	"context"
	"my_project/project-project/internal/data"
)

type ProjectNodeRepo interface {
	FindAll(ctx context.Context) (list []*data.ProjectNode, err error)
}
