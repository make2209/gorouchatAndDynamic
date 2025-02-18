package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"groupChatAndDynamic/api/request"
	"groupChatAndDynamic/api/response"
	"groupChatAndDynamic/inits"
	"groupChatAndDynamic/model"
	"groupChatAndDynamic/pkg"
	"groupChatAndDynamic/utils"
	"math/rand"
	"strconv"
	"time"
)

func UserLoginAndCreate(c *gin.Context) {
	var data request.UserLoginRequest
	if err := c.ShouldBind(&data); err != nil {
		response.CurrencyErrResponse(c, err.Error())
		return
	}
	key := "code:Login:" + data.Tel
	lockKey := "lock:" + data.Tel
	lockKeyRes, _ := inits.Client.Get(context.Background(), lockKey).Int()
	if lockKeyRes == 3 {
		response.CurrencyErrResponse(c, "验证码错误三次,已锁定，请五分钟后重试")
		return
	}
	code, err := inits.Client.Get(context.Background(), key).Result()
	if err != nil || code == "" {
		response.CurrencyErrResponse(c, "请先发送验证码")
		return
	}
	if code != data.Code {
		response.CurrencyErrResponse(c, "验证码错误")
		res := inits.Client.Incr(context.Background(), lockKey).Val()
		if res == 3 {
			inits.Client.Expire(context.Background(), lockKey, 5*time.Minute)
		}
		return
	}
	var user model.User
	err = user.GetUserByUsername(data.Tel)
	if err != nil {
		response.CurrencyErrResponse(c, err.Error())
		return
	}
	if user.ID == 0 {
		userCreate := model.User{
			Tel:      data.Tel,
			Password: utils.MD5(data.Password),
		}
		err = userCreate.Create()
		if err != nil {
			response.CurrencyErrResponse(c, err.Error())
			return
		}
		if userCreate.ID == 0 {
			response.CurrencyErrResponse(c, "用户登录失败")
			return
		}
		token, err := pkg.GetJwtToken(strconv.Itoa(int(userCreate.ID)))
		if err != nil {
			response.CurrencyErrResponse(c, err.Error())
			return
		}
		response.CurrencySuccessResponse(c, "用户登录成功", map[string]interface{}{"token": token})
	}
	if user.Password != utils.MD5(data.Password) {
		response.CurrencyErrResponse(c, "密码错误")
		return
	}
	token, err := pkg.GetJwtToken(strconv.Itoa(int(user.ID)))
	if err != nil {
		response.CurrencyErrResponse(c, err.Error())
		return
	}
	response.CurrencySuccessResponse(c, "用户登录成功", map[string]interface{}{"token": token})
}

func SendSms(c *gin.Context) {
	var data request.SendSmsRequest
	if err := c.ShouldBind(&data); err != nil {
		response.CurrencyErrResponse(c, err.Error())
		return
	}
	if len(data.Tel) != 11 {
		response.CurrencyErrResponse(c, "请输入合理的手机号")
		return
	}
	code := rand.Intn(9000) + 1000
	/*sms, err := pkg.AliYunSendSms(data.Tel, strconv.Itoa(code))
	if err != nil {
		response.CurrencyErrResponse(c, err.Error())
		return
	}
	if *sms.Body.Code != "OK" {
		response.CurrencyErrResponse(c, *sms.Body.Message)
		return
	}*/
	key := "code:" + data.Come + ":" + data.Tel
	err := inits.Client.Set(context.Background(), key, code, time.Minute*5).Err()
	if err != nil {
		response.CurrencyErrResponse(c, err.Error())
		return
	}
	response.CurrencySuccessResponse(c, "短信验证码发送成功", nil)
}

func Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.CurrencyErrResponse(c, err.Error())
		return
	}
	if file == nil {
		response.CurrencyErrResponse(c, "请上传文件")
		return
	}
	if file.Size > 500*1024*1024 {
		response.CurrencyErrResponse(c, "只允许上传500M以内的文件")
		return
	}
	dst := "D:\\GoWork\\src\\zk0212\\static\\" + file.Filename
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		response.CurrencyErrResponse(c, err.Error())
		return
	}
	response.CurrencySuccessResponse(c, "文件上传成功", nil)

}
