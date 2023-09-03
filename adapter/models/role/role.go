package role

import (
	"context"

	"github.com/anditakaesar/uwa-back/adapter/database"
	"github.com/anditakaesar/uwa-back/adapter/models"
	"gorm.io/gorm"
)

const (
	RoleSuperadmin = "superadmin"
	RoleAdmin      = "admin"
	RoleUser       = "user"
)

type Role struct {
	gorm.Model
	models.ChangesInfo
	Name        string
	Description string
}

type RoleModelInterface interface {
	GetRoleByName(ctx context.Context, name string) (*Role, error)
}

type RoleModel struct {
	db *gorm.DB
}

func NewRoleModel(db database.DatabaseInterface) RoleModelInterface {
	return &RoleModel{
		db: db.Get(),
	}
}

func (m *RoleModel) GetRoleByName(ctx context.Context, name string) (*Role, error) {
	role := &Role{}
	err := m.db.WithContext(ctx).Model(&role).Where("name = ?", name).First(role).Error
	if err != nil {
		return nil, err
	}

	return role, nil
}
