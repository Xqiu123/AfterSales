package car

import (
	. "as/handler"
	"as/pkg/errno"

	"as/dao"
	"as/service"
	"github.com/gin-gonic/gin"
)

// Create
// @Summary 新增车辆 api
// @Tags car
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "token 用户令牌"
// @Param object body dao.CarModel true "car"
// @Router /car [post]
func Create(c *gin.Context) {
	var req dao.CarModel
	if err := c.BindJSON(&req); err != nil {
		SendError(c, errno.BadRequest, nil, err.Error(), GetLine())
		return
	}

	if err := service.CreateCar(&req); err != nil {
		SendError(c, err, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, nil, "Car created successfully")
}
