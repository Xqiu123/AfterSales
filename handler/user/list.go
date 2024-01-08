package user

import (
	. "as/handler"
	"as/log"
	"as/pkg/errno"
	"as/service"
	"as/util"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// List ... 获取 userList
// @Summary list user api
// @Tags user
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "token 用户令牌"
// @Param limit query int false "limit"
// @Param page query int false "page"
// @Param last_id query int false "last_id"
// @Success 200 {object} ListResponse
// @Router /user/list [get]
func List(c *gin.Context) {
	log.Info("User List function called.", zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 limit
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "50"))
	if err != nil {
		SendError(c, errno.BadRequest, nil, err.Error(), GetLine())
		return
	}

	// 获取 page
	page, err := strconv.Atoi(c.DefaultQuery("page", "0"))
	if err != nil {
		SendError(c, errno.BadRequest, nil, err.Error(), GetLine())
		return
	}

	lastId, err := strconv.Atoi(c.DefaultQuery("last_id", "0"))
	if err != nil {
		SendError(c, errno.BadRequest, nil, err.Error(), GetLine())
		return
	}

	listResp, err := service.List("", lastId, page*limit, limit)
	if err != nil {
		SendError(c, err, nil, "", GetLine())
		return
	}

	SendResponse(c, nil, listResp)
}
