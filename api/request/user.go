package request

type UserLoginRequest struct {
	Tel      string `form:"tel" json:"tel" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Code     string `form:"code" json:"code" binding:"required"`
}
type SendSmsRequest struct {
	Tel  string `form:"tel" json:"tel" binding:"required"`
	Come string `form:"come" json:"come" binding:"required"`
}
