package model

import (
	"gorm.io/gorm"
	"zk0212/inits"
)

type RoomUser struct {
	gorm.Model
	RoomId       int `gorm:"column:room_id" json:"room_id"`
	UserId       int `gorm:"column:user_id" json:"user_id"`
	Jurisdiction int `gorm:"column:jurisdiction" json:"jurisdiction"` //1 普通成员 2 管理员 3 群主
}

func (r *RoomUser) TableName() string {
	return "room_user"
}
func (r *RoomUser) Create() error {
	return inits.Db.Create(&r).Error
}
func (r *RoomUser) UpdateJurisdiction(RoomId, UserId int) error {
	return inits.Db.Model(&RoomUser{}).Where("room_id = ? and user_id = ?", RoomId, UserId).Update("jurisdiction", 2).Error
}
func (r *RoomUser) GetInfoByAll(roomId, userId int) error {
	return inits.Db.Model(&RoomUser{}).Where("room_id = ? and user_id = ?", roomId, userId).Limit(1).Find(r).Error
}
func (r *RoomUser) DeleteRoomUser(roomId int) error {
	return inits.Db.Model(&RoomUser{}).Where("room_id = ?", roomId).Delete(&RoomUser{}).Error
}
func (r *RoomUser) DeleteRoomUserOne(roomId, userId int) error {
	return inits.Db.Model(&RoomUser{}).Where("room_id = ? and user_id = ?", roomId, userId).Delete(&RoomUser{}).Error
}
