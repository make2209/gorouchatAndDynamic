package model

import (
	"gorm.io/gorm"
	"zk0212/inits"
)

type Follow struct {
	gorm.Model
	UserId       int `gorm:"column:user_id" json:"user_id"`
	FollowUserId int `gorm:"column:follow_user_id" json:"follow_user_id"`
}

func (f *Follow) TableName() string {
	return "follow"
}
func (f *Follow) Create() error {
	return inits.Db.Create(&f).Error
}
func (f *Follow) GetFollowInfo(userId, followUserId int) error {
	return inits.Db.Where("user_id = ? and follow_user_id = ?", userId, followUserId).Limit(1).Find(&f).Error
}
func GetFollowUserIds(userId int) ([]int, error) {
	var ids []int
	err := inits.Db.Model(&Follow{}).Where("user_id = ?", userId).Pluck("follow_user_id", &ids).Error
	if err != nil {
		return nil, err
	}
	return ids, nil
}
