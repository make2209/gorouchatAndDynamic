package response

type Dynamic struct {
	DynamicId    int    `json:"dynamic_id"`
	UserId       int    `json:"user_id"`
	DynamicTitle string `json:"dynamic_title"`
	LikeCount    int    `json:"like_count"`
	ReviewCount  int    `json:"review_count"`
}
type DynamicReview struct {
	DynamicReviewId int    `json:"dynamic_review_id"`
	DynamicId       int    `gorm:"column:dynamic_id" json:"dynamic_id"`
	ReviewContent   string `gorm:"column:review_content" json:"review_content"`
	FatherId        int    `form:"father_id" json:"father_id" binding:"required"`
	LikeCount       int    `gorm:"column:like_count" json:"like_count"`
}
