package user

import (
	. "as/handler"
	"as/log"
	"as/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetMyProfile ... 获取 myProfile
// @Summary get 我的 profile api
// @Description 获取 my 完整 user 信息（权限 role: Normal-普通学生用户; NormalAdmin-学生管理员; Muxi-团队成员; MuxiAdmin-团队管理员; SuperAdmin-超级管理员）
// @Tags user
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "token 用户令牌"
// @Success 200 {object} UserProfile
// @Router /user/myprofile [get]
func GetMyProfile(c *gin.Context) {
	log.Info("User GetMyProfile function called.", zap.String("X-Request-Id", util.GetReqID(c)))

	userId := c.MustGet("userId").(int)

	user, err := GetUserProfile(userId)

	if err != nil {
		SendError(c, err, nil, "", GetLine())
		return
	}

	SendResponse(c, nil, user)
}
