package resposiitory

import (
	"context"
	"gorm.io/gorm"
	"maples/maples-service/module/entity"
)

type UserModel interface {
	Create(ctx context.Context, db *gorm.DB, user *entity.User) error
}

type userModel struct {
}

func NewUserModel() UserModel {
	return &userModel{}
}

func (u *userModel) Create(ctx context.Context, db *gorm.DB, user *entity.User) error {
	db = db.WithContext(ctx).Create(user)
	return db.Error
}
