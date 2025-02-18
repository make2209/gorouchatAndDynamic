package request

type DynamicAddRequest struct {
	Title   string `form:"title" json:"title" binding:"required"`
	Content string `form:"content" json:"content" binding:"required"`
	Types   int    `form:"types" json:"types" binding:"required"`
}
type FollowRequest struct {
	FollowUserId int `form:"follow_user_id" json:"follow_user_id" binding:"required"`
}
type DynamicLikeRequest struct {
	DynamicId int `form:"dynamic_id" json:"dynamic_id" binding:"required"`
}
type CancelDynamicLikeRequest struct {
	DynamicId int `form:"dynamic_id" json:"dynamic_id" binding:"required"`
}
type DynamicLikeListByDynamicRequest struct {
	DynamicId int `form:"dynamic_id" json:"dynamic_id" binding:"required"`
}
type DynamicReviewRequest struct {
	DynamicId     int    `form:"dynamic_id" json:"dynamic_id" binding:"required"`
	ReviewContent string `form:"review_content" json:"review_content" binding:"required"`
	FatherId      int    `form:"father_id" json:"father_id" binding:"required"`
}
type DynamicReviewLikeRequest struct {
	DynamicReviewId int `form:"dynamic_review_id" json:"dynamic_review_id" binding:"required"`
}
