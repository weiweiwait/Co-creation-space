package repo

import (
	"context"
	"my_project/project-project/internal/data"
)

type MenuRepo interface {
	FindMenus(ctx context.Context) ([]*data.ProjectMenu, error)
}
