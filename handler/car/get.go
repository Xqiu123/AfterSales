package car

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

// Get
// @Summary 获取车辆 api
// @Tags car
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "token 用户令牌"
// @Param id path int true "car_id"
// @Router /car/info/{id} [get]
func Get(c *gin.Context) {
	log.Info("GetCar function called.", zap.String("X-Request-Id", util.GetReqID(c)))

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendError(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	car, err := service.GetCarByID(id)
	if err != nil {
		SendError(c, err, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, nil, car)
}
