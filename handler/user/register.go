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

// Register
// @Summary 注册 api
// @Tags auth
// @Accept application/json
// @Produce application/json
// @Param object body RegisterRequest true "register_request"
// @Router /auth/register [post]
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
