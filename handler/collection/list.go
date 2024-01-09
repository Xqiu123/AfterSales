package collection

import (
	. "as/handler"
	"as/log"
	"as/service"
	"as/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// List ... list收藏
// @Summary list收藏 api
// @Tags collection
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "token 用户令牌"
// @Success 200 {object} []dao.CarModel
// @Router /collection/list [get]
func List(c *gin.Context) {
	log.Info("Collection List function called.", zap.String("X-Request-Id", util.GetReqID(c)))

	userId := c.MustGet("userId").(int)

	listResp, err := service.ListCollection(userId)
	if err != nil {
		SendError(c, err, listResp, "", GetLine())
		return
	}

	SendResponse(c, nil, listResp)
}
