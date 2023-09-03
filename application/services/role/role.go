package role

import (
	"context"

	roleModel "github.com/anditakaesar/uwa-back/adapter/models/role"
)

type RoleSeviceInterface interface {
	GetRoleByName(ctx context.Context, name string) (*roleModel.Role, error)
}

type RoleService struct {
	RoleModel roleModel.RoleModelInterface
}

func NewRoleService(roleModel roleModel.RoleModelInterface) RoleSeviceInterface {
	return &RoleService{
		RoleModel: roleModel,
	}
}

func (s *RoleService) GetRoleByName(ctx context.Context, name string) (*roleModel.Role, error) {
	return s.RoleModel.GetRoleByName(ctx, name)
}
