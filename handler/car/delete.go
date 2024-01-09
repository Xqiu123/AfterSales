package car

import (
	. "as/handler"
	"as/pkg/errno"
	"as/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

// Delete
// @Summary 删除车辆 api
// @Tags car
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "token 用户令牌"
// @Param id path int true "car_id"
// @Router /car/{id} [delete]
func Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendError(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	if err := service.DeleteCar(id); err != nil {
		SendError(c, err, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, nil, "Car deleted successfully")
}
