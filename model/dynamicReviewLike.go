package model

import (
	"gorm.io/gorm"
	"zk0212/inits"
)

type DynamicReviewLike struct {
	gorm.Model
	UserId          int `gorm:"column:user_id" json:"user_id"`
	DynamicReviewId int `gorm:"column:dynamic_review_id" json:"dynamic_review_id"`
}

func (d *DynamicReviewLike) TableName() string {
	return "dynamic_review_like"
}
func (d *DynamicReviewLike) Create() error {
	return inits.Db.Create(&d).Error
}
func (d *DynamicReviewLike) GetDynamicReviewLikeInfoByAll(userid, dynamicReviewId int) error {
	return inits.Db.Model(&DynamicReviewLike{}).Where("user_id = ? AND dynamic_review_id = ?", userid, dynamicReviewId).Limit(1).Find(&d).Error
}
func GetDynamicReviewLikeByUserId(userId int) ([]int, error) {
	var dynamicReviewIds []int
	err := inits.Db.Model(&DynamicReviewLike{}).Where("user_id = ?", userId).Pluck("dynamic_review_id", &dynamicReviewIds).Error
	if err != nil {
		return nil, err
	}
	return dynamicReviewIds, nil
}
