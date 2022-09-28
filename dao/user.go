package dao

import (
	"gorm.io/gorm"
	"how-to-init-sql.DB/dao/internal"
	"how-to-init-sql.DB/dao/model"
)

type userCRUD interface { // 可进行 mock，用于调用者的测试中
	Create(user *model.User) error
}

type userDao struct {
	db *gorm.DB
}

var (
	user = &userDao{
		internal.DB(),
	}
)

func User() userCRUD {
	return user
}

func (u *userDao) Create(user *model.User) error {
	return u.db.Create(user).Error
}
