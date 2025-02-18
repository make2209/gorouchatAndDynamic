package request

type GroupAddRequest struct {
	UserIds []int `form:"userIds" json:"userIds" binding:"required"`
}
type SetAdministrationRequest struct {
	RoomId    int `form:"roomId" json:"roomId" binding:"required"`
	SetUserId int `form:"setUserId" json:"setUserId" binding:"required"`
}
type DissolutionGroupRequest struct {
	RoomId int `form:"roomId" json:"roomId" binding:"required"`
}
type AddMemberRequest struct {
	RoomId    int `form:"roomId" json:"roomId" binding:"required"`
	AddUserId int `form:"addUserId" json:"addUserId" binding:"required"`
}
type DeleteMemberRequest struct {
	RoomId    int `form:"roomId" json:"roomId" binding:"required"`
	DelUserId int `form:"delUserId" json:"delUserId" binding:"required"`
}
