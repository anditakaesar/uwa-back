package database

import "github.com/anditakaesar/uwa-back/domain"

func (db *Database) GetRoleByName(roleName string) *domain.Role {
	var role domain.Role
	db.mainDB.First(&role, "name = ?", roleName)
	return &role
}

func (db *Database) CreateRole(role *domain.Role) {
	db.create(role)
}

func (db *Database) CreateUserRole(userRole *domain.UserRole) {
	db.create(userRole)
}
