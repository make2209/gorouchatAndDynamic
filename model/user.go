package model

import (
	"gorm.io/gorm"
	"zk0212/inits"
)

type User struct {
	gorm.Model
	Tel      string `gorm:"column:tel" json:"tel"`
	Password string `gorm:"column:password" json:"password"`
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) GetUserByUsername(tel string) error {
	return inits.Db.Model(&User{}).Where("tel = ? ", tel).Limit(1).Find(&u).Error
}
func (u *User) GetUserById(id int) error {
	return inits.Db.Model(&User{}).Where("id = ?", id).Limit(1).Find(&u).Error
}
func (u *User) Create() error {
	return inits.Db.Create(&u).Error
}
