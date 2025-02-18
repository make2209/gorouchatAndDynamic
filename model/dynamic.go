package model

import (
	"gorm.io/gorm"
	"zk0212/inits"
)

type Dynamic struct {
	gorm.Model
	UserId      int    `gorm:"column:user_id" json:"user_id"`
	Title       string `gorm:"column:title" json:"title"`
	Content     string `gorm:"column:content" json:"content"`
	LikeCount   int    `gorm:"column:like_count" json:"like_count"`
	ReviewCount int    `gorm:"column:review_count" json:"review_count"`
	Types       int    `gorm:"types type:int default:1" json:"types"` //动态类型  1 文本 2图片 3视频 默认为1
}

func (d *Dynamic) TableName() string {
	return "dynamic"
}
func (d *Dynamic) Crete() error {
	return inits.Db.Create(&d).Error
}
func GetDynamicByUserid(userid []int) ([]Dynamic, error) {
	var dynamics []Dynamic
	err := inits.Db.Where("user_id in (?)", userid).Order("like_count desc").Find(&dynamics).Error
	if err != nil {
		return nil, err
	}
	return dynamics, nil
}
func (d *Dynamic) GetDynamicById(id int) error {
	return inits.Db.Model(&Dynamic{}).Where("id = ?", id).Limit(1).Find(&d).Error
}
func (d *Dynamic) DynamicLikeCountAdd(id, likeCount int) error {
	return inits.Db.Model(&Dynamic{}).Where("id = ?", id).Update("like_count", likeCount+1).Error
}
func (d *Dynamic) DynamicLikeCountReduce(id, likeCount int) error {
	return inits.Db.Model(&Dynamic{}).Where("id = ?", id).Update("like_count", likeCount-1).Error
}
func (d *Dynamic) DynamicReviewCountAdd(id, reviewCount int) error {
	return inits.Db.Model(&Dynamic{}).Where("id = ?", id).Update("review_count", reviewCount+1).Error
}
func GetDynamicByNotUserid(userid []int) ([]Dynamic, error) {
	var dynamics []Dynamic
	err := inits.Db.Where("user_id not in (?)", userid).Order("like_count desc").Find(&dynamics).Error
	if err != nil {
		return nil, err
	}
	return dynamics, nil
}
func GetDynamicAll() ([]Dynamic, error) {
	var dynamics []Dynamic
	err := inits.Db.Order("like_count desc").Find(&dynamics).Error
	if err != nil {
		return nil, err
	}
	return dynamics, nil
}
