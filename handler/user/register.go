package user

import (
	. "as/handler"
	"as/log"
	"as/pkg/errno"
	"as/service"
	"as/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Register(c *gin.Context) {
	log.Info("User Register function called.", zap.String("X-Request-Id", util.GetReqID(c)))

	var req RegisterRequest
	if err := c.BindJSON(&req); err != nil {
		SendError(c, errno.BadRequest, nil, err.Error(), GetLine())
		return
	}

	if err := service.Register(req.Telephone, req.Name, req.Password, req.Sex); err != nil {
		SendError(c, err, nil, "", GetLine())
		return
	}

	SendResponse(c, nil, "注册成功")
}
