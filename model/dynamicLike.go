package model

import (
	"gorm.io/gorm"
	"zk0212/inits"
)

type DynamicLike struct {
	gorm.Model
	UserId    int `gorm:"column:user_id" json:"user_id"`
	DynamicId int `gorm:"column:dynamic_id" json:"dynamic_id"`
}

func (d *DynamicLike) TableName() string {
	return "dynamic_like"
}
func (d *DynamicLike) Create() error {
	return inits.Db.Create(&d).Error
}
func (d *DynamicLike) GetDynamicLikeInfoByAll(userId, dynamicId int) error {
	return inits.Db.Model(&DynamicLike{}).Where("user_id = ? and dynamic_id = ?", userId, dynamicId).Limit(1).Find(&d).Error
}
func (d *DynamicLike) DeleteDynamicLike(userId, dynamicId int) error {
	return inits.Db.Model(&DynamicLike{}).Where("user_id = ? and dynamic_id = ?", userId, dynamicId).Delete(&DynamicLike{}).Error
}
func GetDynamicLikeByUserId(userId int) ([]int, error) {
	var dynamicIds []int
	err := inits.Db.Model(&DynamicLike{}).Where("user_id = ?", userId).Pluck("dynamic_id", &dynamicIds).Error
	if err != nil {
		return nil, err
	}
	return dynamicIds, nil
}
func GetDynamicLikeByDynamicId(dynamicId int) ([]int, error) {
	var userIds []int
	err := inits.Db.Model(&DynamicLike{}).Where("dynamic_id = ?", dynamicId).Pluck("user_id", &userIds).Error
	if err != nil {
		return nil, err
	}
	return userIds, nil
}
