package resposiitory

import (
	"context"
	"gorm.io/gorm"
	"maples/maples-service/infra/mysql"
	"maples/maples-service/module/entity"
)

type UserModel interface {
	Create(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, user *entity.User) error
}

type userModel struct {
	db *gorm.DB
}

func NewUserModel() UserModel {
	return &userModel{db: mysql.MysqlClient}
}

func (u *userModel) Create(ctx context.Context, user *entity.User) error {
	db := u.db.WithContext(ctx).Create(user)
	return db.Error
}

func (u *userModel) Update(ctx context.Context, user *entity.User) error {
	db := u.db.WithContext(ctx)
	db = db.Model(&entity.User{}).Where("phone = ?", user.Phone).Updates(entity.User{
		Name:     user.Name,
		Sex:      user.Sex,
		Birthday: user.Birthday,
	})
	return db.Error
}
