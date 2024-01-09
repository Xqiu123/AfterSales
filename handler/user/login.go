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

// Login ... 登录
// @Summary 登录 api
// @Description login the student-as
// @Tags auth
// @Accept application/json
// @Produce application/json
// @Param object body LoginRequest true "login_request"
// @Success 200 {object} LoginResponse
// @Router /auth/login [post]
func Login(c *gin.Context) {
	log.Info("User Login function called.", zap.String("X-Request-Id", util.GetReqID(c)))

	var req LoginRequest
	if err := c.BindJSON(&req); err != nil {
		SendError(c, errno.BadRequest, nil, err.Error(), GetLine())
		return
	}

	token, err := service.Login(req.Telephone, req.Password)
	if err != nil {
		SendError(c, err, nil, "", GetLine())
		return
	}

	SendResponse(c, nil, LoginResponse{
		Token: token,
	})
}
