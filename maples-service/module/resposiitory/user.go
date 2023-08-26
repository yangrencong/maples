package resposiitory

import (
	"context"
	"gorm.io/gorm"
	"maples/maples-service/module/entity"
)

type UserModel interface {
	Create(ctx context.Context, db *gorm.DB, user *entity.User) error
	Find(ctx context.Context, db *gorm.DB, user *entity.User) (*entity.User, error)
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

func (u *userModel) Find(ctx context.Context, db *gorm.DB, user *entity.User) (*entity.User, error) {
	db = db.WithContext(ctx).Debug()
	if user.Name != "" {
		db = db.Where("name = ?", user.Name)
	}
	if user.Phone != "" {
		db = db.Where("phone = ?", user.Phone)
	}
	var ret entity.User
	db = db.First(&ret)
	if db.Error != nil {
		return nil, db.Error
	}
	return &ret, nil
}
