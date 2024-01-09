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

// UpdateInfo ... 修改用户个人信息
// @Summary 修改用户个人信息 api
// @Tags user
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "token 用户令牌"
// @Param object body UpdateInfoRequest  true "update_info_request"
// @Router /user [put]
func UpdateInfo(c *gin.Context) {
	log.Info("User UpdateInfo function called.", zap.String("X-Request-Id", util.GetReqID(c)))

	var req UpdateInfoRequest
	if err := c.BindJSON(&req); err != nil {
		SendError(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	userId := c.MustGet("userId").(int)

	if err := service.UpdateInfo(userId, req.Name); err != nil {
		SendError(c, err, nil, "", GetLine())
		return
	}

	SendResponse(c, nil, nil)
}
