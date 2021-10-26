package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"realness/auth"
	"realness/model"
	"strings"
)

// AuthMiddleware 用户 token 认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取 Authorization, header
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "token无效"})
			// 抛弃请求
			ctx.Abort()
			return
		}
		tokenString = tokenString[7:]
		token, claims, err := auth.ParseUserToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "Token解析失败"})
			// 抛弃请求
			ctx.Abort()
			return
		}
		// 通过 token 获取userID
		userID := claims.UserID
		var user model.User
		model.DB.First(&user, userID)
		// 验证用户是否存在
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "token解析成功，用户不存在"})
			ctx.Abort()
			return
		}
		// 如果用户存在，将用户信息写入上下文
		ctx.Set("id", user.ID)
		ctx.Set("user", user)
		ctx.Next()
	}
}
