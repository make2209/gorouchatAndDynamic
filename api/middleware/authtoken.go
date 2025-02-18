package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zk0212/pkg"
)

func AuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "请登录",
				"data": nil,
			})
			return
		}
		jwtToken, err := pkg.ParseJwtToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "token解析失败" + err.Error(),
				"data": nil,
			})
			return
		}
		c.Set("userid", jwtToken)
		c.Next()
	}
}
