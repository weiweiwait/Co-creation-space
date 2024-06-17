package user

import (
	"my_project/project-user/pkg/dao"
	"my_project/project-user/pkg/repo"
)

type HandlerUser struct {
	cache repo.Cache
}

func New() *HandlerUser {
	return &HandlerUser{
		cache: dao.Rc,
	}
}
