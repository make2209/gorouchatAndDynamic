package model

import (
	"gorm.io/gorm"
	"zk0212/inits"
)

type DynamicReview struct {
	gorm.Model
	UserId        int    `gorm:"column:user_id" json:"user_id"`
	DynamicId     int    `gorm:"column:dynamic_id" json:"dynamic_id"`
	ReviewContent string `gorm:"column:review_content" json:"review_content"`
	FatherId      int    `form:"father_id" json:"father_id" binding:"required"`
	LikeCount     int    `gorm:"column:like_count" json:"like_count"`
}

func (d *DynamicReview) TableName() string {
	return "dynamic_review"
}

func (d *DynamicReview) Create() error {
	return inits.Db.Create(&d).Error
}
func (d *DynamicReview) GetReviewInfoById(id int) error {
	return inits.Db.Model(&DynamicReview{}).Where("id = ?", id).Limit(1).Find(&d).Error
}
func (d *DynamicReview) DynamicReviewLikeCountAdd(id, likeCount int) error {
	return inits.Db.Model(&DynamicReview{}).Where("id = ?", id).Update("like_count", likeCount+1).Error
}
func GetDynamicReviewByUserId(userId int) ([]DynamicReview, error) {
	var DynamicReviews []DynamicReview
	err := inits.Db.Model(&DynamicReview{}).Where("user_id = ?", userId).Find(&DynamicReviews).Error
	if err != nil {
		return nil, err
	}
	return DynamicReviews, nil
}
