package user

import (
	"context"

	"github.com/anditakaesar/uwa-back/adapter/database"
	"github.com/anditakaesar/uwa-back/adapter/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	models.ChangesInfo
	GUID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Username    string    `gorm:"unique,index"`
	Email       string    `gorm:"unique"`
	Password    string
	PhoneNumber string
	RoleID      int
}

type UserModelInterface interface {
	CreateUser(ctx context.Context, user *User) error
	GetUserByUsername(ctx context.Context, username string) (*User, error)
}

type UserModel struct {
	db *gorm.DB
}

func NewUserModel(db database.DatabaseInterface) UserModelInterface {
	return &UserModel{
		db: db.Get(),
	}
}

func (m *UserModel) CreateUser(ctx context.Context, user *User) error {
	return m.db.WithContext(ctx).Create(user).Error
}

func (m *UserModel) GetUserByUsername(ctx context.Context, username string) (*User, error) {
	user := &User{}
	err := m.db.WithContext(ctx).Model(&user).Where("username = ?", username).First(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
