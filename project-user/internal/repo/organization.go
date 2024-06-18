package repo

import (
	"context"
	"my_project/project-user/internal/data/organization"
)

type OrganizationRepo interface {
	SaveOrganization(ctx context.Context, org *organization.Organization) error
	FindOrganizationByMemId(ctx context.Context, memId int64) ([]*organization.Organization, error)
}
