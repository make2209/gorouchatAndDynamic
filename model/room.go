package model

import (
	"gorm.io/gorm"
	"zk0212/inits"
)

type Room struct {
	gorm.Model
	UserId int `gorm:"column:user_id" json:"user_id"`
}

func (r *Room) TableName() string {
	return "room"
}
func (r *Room) Create() error {
	return inits.Db.Create(&r).Error
}
func (r *Room) GetRoomById(id int) error {
	return inits.Db.Model(&Room{}).Where("id = ?", id).Limit(1).Find(&r).Error
}
func (r *Room) GetRoomByRoomIdAndUserId(roomId, userId int) error {
	return inits.Db.Model(&Room{}).Where("id = ? and user_id = ?", roomId, userId).Limit(1).Find(&r).Error
}
func (r *Room) DeleteRoom(roomId int) error {
	return inits.Db.Model(&Room{}).Where("id = ?", roomId).Delete(&Room{}).Error
}
