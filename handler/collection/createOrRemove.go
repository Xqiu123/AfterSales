package collection

import (
	. "as/handler"
	"as/log"
	"as/pkg/errno"
	"as/service"
	"as/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// CreateOrRemove ... 收藏/取消收藏帖子
// @Summary 收藏/取消收藏帖子 api
// @Tags collection
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "token 用户令牌"
// @Param car_id path int true "car_id"
// @Success 200 {object} handler.Response
// @Router /collection/{car_id} [post]
func CreateOrRemove(c *gin.Context) {
	log.Info("Collection CreateOrRemove function called.", zap.String("X-Request-Id", util.GetReqID(c)))

	carId, err := strconv.Atoi(c.Param("car_id"))
	if err != nil {
		SendError(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	userId := c.MustGet("userId").(int)

	if err := service.CreateOrRemoveCollection(userId, carId); err != nil {
		SendError(c, err, nil, "", GetLine())
		return
	}

	SendResponse(c, err, nil)
}
