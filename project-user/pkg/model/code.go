package model

import (
	"my_project/project-common/errs"
)

var (
	NoLegalMobile = errs.NewError(2001, "手机号不合法")
)
