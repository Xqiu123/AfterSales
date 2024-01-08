package middleware

import (
	"as/handler"
	"as/pkg/auth"
	"as/pkg/errno"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware ... 认证中间件
// limit 为限制的权限等级
func AuthMiddleware(limit int) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		ctx, err := auth.ParseRequest(c)
		if err != nil {
			handler.SendError(c, errno.ErrAuthToken, nil, err.Error(), handler.GetLine())
			c.Abort()
			return
		} else if ctx.Role&limit == 0 {
			handler.SendError(c, errno.ErrPermissionDenied, nil, "", handler.GetLine())
			c.Abort()
			return
		}

		c.Set("userId", ctx.Id)
		c.Set("role", ctx.Role)
		c.Set("expiresAt", ctx.ExpiresAt)

		c.Next()
	}
}
