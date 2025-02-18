package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"zk0212/api/request"
	"zk0212/api/response"
	"zk0212/inits"
	"zk0212/model"
)

func AddGroup(c *gin.Context) {
	var data request.GroupAddRequest
	if err := c.ShouldBind(&data); err != nil {
		response.CurrencyErrResponse(c, err.Error())
		return
	}
	userid, _ := strconv.Atoi(c.GetString("userid"))
	if userid == 0 {
		response.CurrencyErrResponse(c, "请登录后创建群聊")
		return
	}
	tx := inits.Db.Begin()
	room := model.Room{
		UserId: userid,
	}
	err := room.Create()
	if err != nil {
		response.CurrencyErrResponse(c, err.Error())
		tx.Rollback()
		return
	}
	if room.ID == 0 {
		response.CurrencyErrResponse(c, "群聊创建失败")
		tx.Rollback()
		return
	}
	user := model.RoomUser{
		RoomId:       int(room.ID),
		UserId:       userid,
		Jurisdiction: 3,
	}
	err = user.Create()
	if err != nil {
		response.CurrencyErrResponse(c, err.Error())
		tx.Rollback()
		return
	}
	if user.ID == 0 {
		response.CurrencyErrResponse(c, "群聊创建失败")
		tx.Rollback()
		return
	}
	tx.Commit()
	for _, id := range data.UserIds {
		roomUser := model.RoomUser{
			RoomId:       int(room.ID),
			UserId:       id,
			Jurisdiction: 1,
		}
		err := roomUser.Create()
		if err != nil {
			continue
		}
	}
	response.CurrencySuccessResponse(c, "群聊创建成功", map[string]interface{}{"roomId": room.ID})

}

func AddMember(c *gin.Context) {
	var data request.AddMemberRequest
	if err := c.ShouldBind(&data); err != nil {
		response.CurrencyErrResponse(c, err.Error())
		return
	}
	userid, _ := strconv.Atoi(c.GetString("userid"))
	if userid == 0 {
		response.CurrencyErrResponse(c, "请登录后邀请人进入群聊")
		return
	}
	var room model.Room
	err := room.GetRoomById(data.RoomId)
	if err != nil {
		response.CurrencyErrResponse(c, err.Error())
		return
	}
	if room.ID == 0 {
		response.CurrencyErrResponse(c, "该房间不存在")
		return
	}
	var roomUser model.RoomUser
	err = roomUser.GetInfoByAll(data.RoomId, userid)
	if err != nil {
		response.CurrencyErrResponse(c, err.Error())
		return
	}
	if roomUser.ID == 0 {
		response.CurrencyErrResponse(c, "你不是该群里的成员")
		return
	}
	if roomUser.Jurisdiction != 1 && roomUser.Jurisdiction != 3 {
		response.CurrencyErrResponse(c, "只有群主和管理员有权限邀请人进群")
		return
	}
	var user model.User
	err = user.GetUserById(data.AddUserId)
	if err != nil {
		response.CurrencyErrResponse(c, err.Error())
		return
	}
	if user.ID == 0 {
		response.CurrencyErrResponse(c, "你要邀请进群的人不存在")
		return
	}
	var roomUserValid model.RoomUser
	err = roomUserValid.GetInfoByAll(data.RoomId, data.AddUserId)
	if err != nil {
		response.CurrencyErrResponse(c, err.Error())
		return
	}
	if roomUserValid.ID != 0 {
		response.CurrencyErrResponse(c, "用户已在群内")
		return
	}
	roomUserCreate := model.RoomUser{
		RoomId:       data.RoomId,
		UserId:       data.AddUserId,
		Jurisdiction: 1,
	}
	err = roomUserCreate.Create()
	if err != nil {
		response.CurrencyErrResponse(c, err.Error())
		return
	}
	if roomUserCreate.ID == 0 {
		response.CurrencyErrResponse(c, "邀请好友进群失败")
		return
	}
	response.CurrencySuccessResponse(c, "邀请好友进群成功", nil)
}

func DeleteMember(c *gin.Context) {
	var data request.DeleteMemberRequest
	if err := c.ShouldBind(&data); err != nil {
		response.CurrencyErrResponse(c, err.Error())
		return
	}
	userid, _ := strconv.Atoi(c.GetString("userid"))
	if userid == 0 {
		response.CurrencyErrResponse(c, "请登录后邀请人进入群聊")
		return
	}
	var room model.Room
	err := room.GetRoomById(data.RoomId)
	if err != nil {
		response.CurrencyErrResponse(c, err.Error())
		return
	}
	if room.ID == 0 {
		response.CurrencyErrResponse(c, "该房间不存在")
		return
	}
	var roomUser model.RoomUser
	err = roomUser.GetInfoByAll(data.RoomId, userid)
	if err != nil {
		response.CurrencyErrResponse(c, err.Error())
		return
	}
	if roomUser.ID == 0 {
		response.CurrencyErrResponse(c, "你不是该群里的成员")
		return
	}
	if roomUser.Jurisdiction != 1 && roomUser.Jurisdiction != 3 {
		response.CurrencyErrResponse(c, "只有群主和管理员有权限移除群成员")
		return
	}
	var roomUserValid model.RoomUser
	err = roomUserValid.GetInfoByAll(data.RoomId, data.DelUserId)
	if err != nil {
		response.CurrencyErrResponse(c, err.Error())
		return
	}
	if roomUserValid.ID == 0 {
		response.CurrencyErrResponse(c, "该用户已不在群内")
		return
	}
	var roomUserDel model.RoomUser
	err = roomUserDel.DeleteRoomUserOne(data.RoomId, data.DelUserId)
	if err != nil {
		response.CurrencyErrResponse(c, err.Error())
		return
	}
	response.CurrencySuccessResponse(c, "群成员移除成功", nil)
}

func SetAdministration(c *gin.Context) {
	var data request.SetAdministrationRequest
	if err := c.ShouldBind(&data); err != nil {
		response.CurrencyErrResponse(c, err.Error())
		return
	}
	userid, _ := strconv.Atoi(c.GetString("userid"))
	if userid == 0 {
		response.CurrencyErrResponse(c, "请登录后设置管理员")
		return
	}
	var room model.Room
	err := room.GetRoomById(data.RoomId)
	if err != nil {
		response.CurrencyErrResponse(c, err.Error())
		return
	}
	if room.ID == 0 {
		response.CurrencyErrResponse(c, "该房间不存在")
		return
	}
	if room.UserId != userid {
		response.CurrencyErrResponse(c, "你不是该房间的群主，无法设置管理员")
		return
	}
	var roomer model.RoomUser
	err = roomer.GetInfoByAll(data.RoomId, data.SetUserId)
	if err != nil {
		response.CurrencyErrResponse(c, err.Error())
		return
	}
	if roomer.ID == 0 || roomer.Jurisdiction == 2 {
		response.CurrencyErrResponse(c, "该用户已在群聊内，或该用户已是管理员")
		return
	}
	err = roomer.UpdateJurisdiction(data.RoomId, data.SetUserId)
	if err != nil {
		response.CurrencyErrResponse(c, err.Error())
		return
	}
	response.CurrencySuccessResponse(c, "设置管理员成功", map[string]interface{}{"setId": roomer.UserId})
}

func DissolutionGroup(c *gin.Context) {
	var data request.DissolutionGroupRequest
	if err := c.ShouldBind(&data); err != nil {
		response.CurrencyErrResponse(c, err.Error())
		return
	}
	userid, _ := strconv.Atoi(c.GetString("userid"))
	if userid == 0 {
		response.CurrencyErrResponse(c, "请登录后设置管理员")
		return
	}
	var room model.Room
	err := room.GetRoomById(data.RoomId)
	if err != nil {
		response.CurrencyErrResponse(c, err.Error())
		return
	}
	if room.ID == 0 {
		response.CurrencyErrResponse(c, "该房间不存在")
		return
	}
	if room.UserId != userid {
		response.CurrencyErrResponse(c, "你不是该房间的群主，无法设置解散群聊")
		return
	}
	tx := inits.Db.Begin()
	var roomer model.Room
	err = roomer.DeleteRoom(data.RoomId)
	if err != nil {
		response.CurrencyErrResponse(c, err.Error())
		tx.Rollback()
		return
	}
	var roomUser model.RoomUser
	err = roomUser.DeleteRoomUser(data.RoomId)
	if err != nil {
		response.CurrencyErrResponse(c, err.Error())
		tx.Rollback()
		return
	}
	err = tx.Commit().Error
	if err != nil {
		response.CurrencyErrResponse(c, err.Error())
		return
	}
	response.CurrencySuccessResponse(c, "解散群聊成功", map[string]interface{}{"roomId": room.ID})
}
