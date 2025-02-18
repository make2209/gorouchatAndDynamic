package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"groupChatAndDynamic/api/request"
	"groupChatAndDynamic/api/response"
	"groupChatAndDynamic/inits"
	"groupChatAndDynamic/model"
	"strconv"
)

func DynamicAdd(c *gin.Context) {
	var data request.DynamicAddRequest
	if err := c.ShouldBind(&data); err != nil {
		response.CurrencyErrResponse(c, err.Error())
		return
	}
	if data.Types != 1 && data.Types != 2 && data.Types != 3 {
		response.CurrencyErrResponse(c, "请输入合理的动态类型")
		return
	}
	userid, _ := strconv.Atoi(c.GetString("userid"))
	if userid == 0 {
		response.CurrencyErrResponse(c, "请登录后发布动态")
		return
	}
	dynamicCreate := model.Dynamic{
		UserId:    userid,
		Title:     data.Title,
		Content:   data.Content,
		LikeCount: 0,
		Types:     data.Types,
	}
	err := dynamicCreate.Crete()
	if err != nil {
		response.CurrencyErrResponse(c, err.Error())
		return
	}
	if dynamicCreate.ID == 0 {
		response.CurrencyErrResponse(c, "文章发布失败")
		return
	}
	response.CurrencySuccessResponse(c, "文章发布成功", map[string]interface{}{"dynamicId": dynamicCreate.ID})
}

func Follow(c *gin.Context) {
	var data request.FollowRequest
	if err := c.ShouldBind(&data); err != nil {
		response.CurrencyErrResponse(c, err.Error())
		return
	}
	userid, _ := strconv.Atoi(c.GetString("userid"))
	if userid == 0 {
		response.CurrencyErrResponse(c, "请登录后关注他人")
		return
	}
	if userid == data.FollowUserId {
		response.CurrencyErrResponse(c, "不可关注自己")
		return
	}
	var user model.User
	err := user.GetUserById(data.FollowUserId)
	if err != nil {
		response.CurrencyErrResponse(c, err.Error())
		return
	}
	if user.ID == 0 {
		response.CurrencyErrResponse(c, "你要关注的人不存在")
		return
	}
	var validFollow model.Follow
	err = validFollow.GetFollowInfo(userid, data.FollowUserId)
	if err != nil {
		response.CurrencyErrResponse(c, err.Error())
		return
	}
	if validFollow.ID != 0 {
		response.CurrencyErrResponse(c, "你已关注他,无需重复关注")
		return
	}
	followAdd := model.Follow{
		UserId:       userid,
		FollowUserId: data.FollowUserId,
	}
	err = followAdd.Create()
	if err != nil {
		response.CurrencyErrResponse(c, err.Error())
		return
	}
	if followAdd.ID == 0 {
		response.CurrencyErrResponse(c, "关注失败")
		return
	}
	response.CurrencySuccessResponse(c, "关注成功", nil)
}

func DynamicList(c *gin.Context) {
	userid, _ := strconv.Atoi(c.GetString("userid"))
	if userid == 0 {
		response.CurrencyErrResponse(c, "请登录后查看列表")
		return
	}
	res, err := model.GetFollowUserIds(userid)
	if err != nil {
		response.CurrencyErrResponse(c, err.Error())
		fmt.Println(111)
		return
	}
	if len(res) == 0 {
		all, err := model.GetDynamicAll()
		if err != nil {
			response.CurrencyErrResponse(c, err.Error())
			fmt.Println(222)
			return
		}
		var dynamicList []*response.Dynamic
		for _, v := range all {
			dynamicList = append(dynamicList, &response.Dynamic{
				DynamicId:    int(v.ID),
				UserId:       v.UserId,
				DynamicTitle: v.Title,
				LikeCount:    v.LikeCount,
				ReviewCount:  v.ReviewCount,
			})
		}
		response.CurrencySuccessResponse(c, "您未关注任何人，已按点赞量排序展示", map[string]interface{}{"dynamicList": dynamicList})
		return
	}
	var dynamicList []*response.Dynamic
	followDynamic, err := model.GetDynamicByUserid(res)
	if err != nil {
		response.CurrencyErrResponse(c, err.Error())
		return
	}
	for _, v := range followDynamic {
		dynamicList = append(dynamicList, &response.Dynamic{
			DynamicId:    int(v.ID),
			UserId:       v.UserId,
			DynamicTitle: v.Title,
			LikeCount:    v.LikeCount,
			ReviewCount:  v.ReviewCount,
		})
	}
	notFollowDynamic, err := model.GetDynamicByNotUserid(res)
	if err != nil {
		response.CurrencyErrResponse(c, err.Error())
		return
	}
	for _, v := range notFollowDynamic {
		dynamicList = append(dynamicList, &response.Dynamic{
			DynamicId:    int(v.ID),
			UserId:       v.UserId,
			DynamicTitle: v.Title,
			LikeCount:    v.LikeCount,
			ReviewCount:  v.ReviewCount,
		})
	}
	response.CurrencySuccessResponse(c, "已优先展示你的关注，其次按照点赞量排序", map[string]interface{}{"dynamicList": dynamicList})
	return
}
func DynamicReview(c *gin.Context) {
	var data request.DynamicReviewRequest
	if err := c.ShouldBind(&data); err != nil {
		response.CurrencyErrResponse(c, err.Error())
		return
	}
	if len(data.ReviewContent) > 100 {
		response.CurrencyErrResponse(c, "评论的字数不能超过100字")
		return
	}
	userid, _ := strconv.Atoi(c.GetString("userid"))
	if userid == 0 {
		response.CurrencyErrResponse(c, "登录后才可以评论")
		return
	}
	var dynamic model.Dynamic
	err := dynamic.GetDynamicById(data.DynamicId)
	if err != nil {
		response.CurrencyErrResponse(c, err.Error())
		return
	}
	if dynamic.ID == 0 {
		response.CurrencyErrResponse(c, "该动态不存在")
		return
	}
	if data.FatherId != 0 {
		var dynamicReview model.DynamicReview
		err = dynamicReview.GetReviewInfoById(data.FatherId)
		if err != nil {
			response.CurrencyErrResponse(c, err.Error())
			return
		}
		if dynamicReview.ID == 0 {
			response.CurrencyErrResponse(c, "你要回复的评论不存在")
			return
		}
	}
	tx := inits.Db.Begin()
	reviewCreate := model.DynamicReview{
		UserId:        userid,
		DynamicId:     int(dynamic.ID),
		ReviewContent: data.ReviewContent,
		FatherId:      data.FatherId,
	}
	err = reviewCreate.Create()
	if err != nil {
		response.CurrencyErrResponse(c, err.Error())
		tx.Rollback()
		return
	}
	if reviewCreate.ID == 0 {
		response.CurrencyErrResponse(c, "评论失败")
		tx.Rollback()
		return
	}
	var updateDynamicReviewCount model.Dynamic
	err = updateDynamicReviewCount.DynamicReviewCountAdd(data.DynamicId, dynamic.ReviewCount)
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
	response.CurrencySuccessResponse(c, "评论成功", nil)
}

func DynamicReviewLike(c *gin.Context) {
	var data request.DynamicReviewLikeRequest
	if err := c.ShouldBind(&data); err != nil {
		response.CurrencyErrResponse(c, err.Error())
		return
	}
	userid, _ := strconv.Atoi(c.GetString("userid"))
	if userid == 0 {
		response.CurrencyErrResponse(c, "登录后才可以点赞")
		return
	}
	var dynamicReview model.DynamicReview
	err := dynamicReview.GetReviewInfoById(data.DynamicReviewId)
	if err != nil {
		response.CurrencyErrResponse(c, err.Error())
		return
	}
	if dynamicReview.ID == 0 {
		response.CurrencyErrResponse(c, "你要点赞的评论不存在")
		return
	}
	var dynamicReviewLike model.DynamicReviewLike
	err = dynamicReviewLike.GetDynamicReviewLikeInfoByAll(userid, data.DynamicReviewId)
	if err != nil {
		response.CurrencyErrResponse(c, err.Error())
		return
	}
	if dynamicReview.ID != 0 {
		response.CurrencyErrResponse(c, "你已给该评论点过赞")
		return
	}
	tx := inits.Db.Begin()
	dynamicReviewLikeCreate := model.DynamicReviewLike{
		UserId:          userid,
		DynamicReviewId: data.DynamicReviewId,
	}
	err = dynamicReviewLikeCreate.Create()
	if err != nil {
		response.CurrencyErrResponse(c, err.Error())
		tx.Rollback()
		return
	}
	if dynamicReviewLikeCreate.ID == 0 {
		response.CurrencyErrResponse(c, "评论点赞失败")
		tx.Rollback()
		return
	}
	var updateDynamicReview model.DynamicReview
	err = updateDynamicReview.DynamicReviewLikeCountAdd(data.DynamicReviewId, dynamicReview.LikeCount)
	if err != nil {
		response.CurrencyErrResponse(c, err.Error())
		tx.Rollback()
		return
	}
	response.CurrencySuccessResponse(c, "评论点赞成功", nil)
}

func DynamicLike(c *gin.Context) {
	var data request.DynamicLikeRequest
	if err := c.ShouldBind(&data); err != nil {
		response.CurrencyErrResponse(c, err.Error())
		return
	}
	userid, _ := strconv.Atoi(c.GetString("userid"))
	if userid == 0 {
		response.CurrencyErrResponse(c, "登录后才可以给他人点赞")
		return
	}
	var dynamic model.Dynamic
	err := dynamic.GetDynamicById(data.DynamicId)
	if err != nil {
		response.CurrencyErrResponse(c, err.Error())
		return
	}
	if dynamic.ID == 0 {
		response.CurrencyErrResponse(c, "该动态不存在")
		return
	}
	var dynamicLike model.DynamicLike
	err = dynamicLike.GetDynamicLikeInfoByAll(userid, data.DynamicId)
	if err != nil {
		response.CurrencyErrResponse(c, err.Error())
		return
	}
	if dynamicLike.ID != 0 {
		response.CurrencyErrResponse(c, "你已给该动态点赞了")
		return
	}
	tx := inits.Db.Begin()
	dynamicLikeCreate := model.DynamicLike{
		UserId:    userid,
		DynamicId: data.DynamicId,
	}
	err = dynamicLikeCreate.Create()
	if err != nil {
		response.CurrencyErrResponse(c, err.Error())
		tx.Rollback()
		return
	}
	if dynamicLikeCreate.ID == 0 {
		response.CurrencyErrResponse(c, "点赞失败")
		tx.Rollback()
		return
	}
	var updateDynamicLike model.Dynamic
	err = updateDynamicLike.DynamicLikeCountAdd(data.DynamicId, dynamic.LikeCount)
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
	response.CurrencySuccessResponse(c, "点赞成功", nil)
}

func CancelDynamicLike(c *gin.Context) {
	var data request.CancelDynamicLikeRequest
	if err := c.ShouldBind(&data); err != nil {
		response.CurrencyErrResponse(c, err.Error())
		return
	}
	userid, _ := strconv.Atoi(c.GetString("userid"))
	if userid == 0 {
		response.CurrencyErrResponse(c, "请先登录")
		return
	}
	var dynamic model.Dynamic
	err := dynamic.GetDynamicById(data.DynamicId)
	if err != nil {
		response.CurrencyErrResponse(c, err.Error())
		return
	}
	if dynamic.ID == 0 {
		response.CurrencyErrResponse(c, "该动态不存在")
		return
	}
	var dynamicLike model.DynamicLike
	err = dynamicLike.GetDynamicLikeInfoByAll(userid, data.DynamicId)
	if err != nil {
		response.CurrencyErrResponse(c, err.Error())
		return
	}
	if dynamicLike.ID == 0 {
		response.CurrencyErrResponse(c, "你未给该动态点过赞")
		return
	}
	tx := inits.Db.Begin()
	var deleteDynamicLike model.DynamicLike
	err = deleteDynamicLike.DeleteDynamicLike(userid, data.DynamicId)
	if err != nil {
		response.CurrencyErrResponse(c, err.Error())
		tx.Rollback()
		return
	}
	var updateDynamicLike model.Dynamic
	err = updateDynamicLike.DynamicLikeCountReduce(data.DynamicId, dynamic.LikeCount)
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
	response.CurrencySuccessResponse(c, "取消点赞成功", nil)
}
func DynamicReviewLikeListByUser(c *gin.Context) {
	userid, _ := strconv.Atoi(c.GetString("userid"))
	if userid == 0 {
		response.CurrencyErrResponse(c, "请先登录")
		return
	}
	DynamicReviewIds, err := model.GetDynamicReviewLikeByUserId(userid)
	if err != nil {
		response.CurrencyErrResponse(c, err.Error())
		return
	}
	if len(DynamicReviewIds) == 0 {
		response.CurrencySuccessResponse(c, "你未给任何评论点过赞", nil)
		return
	}
	response.CurrencySuccessResponse(c, "已查询到你的点赞评论列表", map[string]interface{}{"dynamicReviewIds": DynamicReviewIds})
}
func DynamicLikeListByUser(c *gin.Context) {
	userid, _ := strconv.Atoi(c.GetString("userid"))
	if userid == 0 {
		response.CurrencyErrResponse(c, "请先登录")
		return
	}
	DynamicIds, err := model.GetDynamicLikeByUserId(userid)
	if err != nil {
		response.CurrencyErrResponse(c, err.Error())
		return
	}
	if len(DynamicIds) == 0 {
		response.CurrencySuccessResponse(c, "你未给任何动态点过赞", nil)
		return
	}
	response.CurrencySuccessResponse(c, "已查询到你的点赞动态列表", map[string]interface{}{"dynamicIds": DynamicIds})
}
func DynamicReviewListByUser(c *gin.Context) {
	userid, _ := strconv.Atoi(c.GetString("userid"))
	if userid == 0 {
		response.CurrencyErrResponse(c, "请先登录")
		return
	}
	res, err := model.GetDynamicReviewByUserId(userid)
	if err != nil {
		response.CurrencyErrResponse(c, err.Error())
		return
	}
	if len(res) == 0 {
		response.CurrencyErrResponse(c, "您未进行过任何评论")
		return
	}
	var DynamicReviewList []*response.DynamicReview
	for _, v := range res {
		DynamicReviewList = append(DynamicReviewList, &response.DynamicReview{
			DynamicReviewId: int(v.ID),
			DynamicId:       v.DynamicId,
			ReviewContent:   v.ReviewContent,
			FatherId:        v.FatherId,
			LikeCount:       v.LikeCount,
		})
	}
	response.CurrencySuccessResponse(c, "已展示您的评论列表", map[string]interface{}{"DynamicReviewList": DynamicReviewList})
}
func DynamicLikeListByDynamic(c *gin.Context) {
	var data request.DynamicLikeListByDynamicRequest
	if err := c.ShouldBind(&data); err != nil {
		response.CurrencyErrResponse(c, err.Error())
		return
	}
	userid, _ := strconv.Atoi(c.GetString("userid"))
	if userid == 0 {
		response.CurrencyErrResponse(c, "请先登录")
		return
	}
	var dynamic model.Dynamic
	err := dynamic.GetDynamicById(data.DynamicId)
	if err != nil {
		response.CurrencyErrResponse(c, err.Error())
		return
	}
	if dynamic.ID == 0 {
		response.CurrencyErrResponse(c, "该篇动态不存在")
		return
	}
	if dynamic.UserId != userid {
		response.CurrencyErrResponse(c, "你不是该篇动态的作者，无权查看谁给你点赞了")
		return
	}
	if dynamic.LikeCount == 0 {
		response.CurrencySuccessResponse(c, "暂未有人给你这篇动态点过赞", nil)
		return
	}
	userIds, err := model.GetDynamicLikeByDynamicId(data.DynamicId)
	if err != nil {
		response.CurrencyErrResponse(c, err.Error())
		return
	}
	if len(userIds) == 0 {
		response.CurrencySuccessResponse(c, "暂未有人给你这篇动态点过赞", nil)
		return
	}
	response.CurrencySuccessResponse(c, "已查询到这篇动态有那些用户点赞了", map[string]interface{}{"userIds": userIds})
}
