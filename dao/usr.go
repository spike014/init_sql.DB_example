package dao

import (
	"gorm.io/gorm"
	"how-to-init-sql.DB/dao/internal"
	"how-to-init-sql.DB/dao/model"
)

type userDao struct {
	db *gorm.DB
}

var (
	User = &userDao{
		internal.DB(),
	}
)

func (u *userDao) Create(user *model.User) error {
	return u.db.Create(user).Error
}
