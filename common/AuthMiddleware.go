package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//=============================gin的中间件，就是一个函数，返回gin 的HandlerFunc======================================================

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		tokenString := c.Query("token")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足666"})
			c.Abort()
			return
		}

		claims, err := ParsenToken(tokenString)
		if err != nil { //解析失败，或者解析后的token无效
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足...."})
			c.Abort()
			return
		}

		userId := claims.UserId

		//用户存在 将user的信息写入上下文
		c.Set("UserId", userId)
		c.Next()
	}
}
