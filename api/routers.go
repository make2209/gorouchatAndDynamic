package api

import (
	"github.com/gin-gonic/gin"
	"zk0212/api/middleware"
	"zk0212/controller"
)

func LoadRouters(r *gin.Engine) {

	user := r.Group("/user")
	{
		user.POST("/sendSms", controller.SendSms)                   //短信发送
		user.POST("/loginAndCreate", controller.UserLoginAndCreate) //用户注册登录一体化（验证码）
	}
	community := r.Group("/community")
	community.Use(middleware.AuthToken())
	{
		community.POST("/dynamic/add", controller.DynamicAdd)                                 //动态发布
		community.POST("/upload", controller.Upload)                                          //文件上传
		community.GET("/chat", controller.Chat)                                               //一对一聊天
		community.POST("/follow", controller.Follow)                                          //关注作者
		community.GET("/dynamic/list", controller.DynamicList)                                //动态列表展示（优先展示关注的 再按点赞量排序）
		community.POST("/dynamic/like", controller.DynamicLike)                               //动态点赞
		community.POST("/dynamic/like/cancel", controller.CancelDynamicLike)                  //动态点赞取消
		community.POST("/dynamic/review", controller.DynamicReview)                           //动态评论
		community.POST("/dynamic/review/like", controller.DynamicReviewLike)                  //动态评论点赞
		community.GET("/dynamicReviewLikeListByUser", controller.DynamicReviewLikeListByUser) //获取用户给哪些评论点赞了
		community.GET("/dynamicReviewListByUser", controller.DynamicReviewListByUser)         //获取用户给哪写动态评论了
		community.GET("/dynamicLikeListByUser", controller.DynamicLikeListByUser)             //获取用户给哪些动态点赞了
		community.GET("/dynamicLikeListByDynamic", controller.DynamicLikeListByDynamic)       //获取动态有哪些用户点赞了
	}
	groupChat := r.Group("/groupChat")
	groupChat.Use(middleware.AuthToken())
	{
		groupChat.POST("/addGroup", controller.AddGroup)                  //创建群聊
		groupChat.POST("/addMember", controller.AddMember)                //邀请群成员
		groupChat.POST("/deleteMember", controller.DeleteMember)          //删除群成员
		groupChat.POST("setAdministration", controller.SetAdministration) //设置管理员
		groupChat.POST("/dissolutionGroup", controller.DissolutionGroup)  //解散群聊
	}
}
