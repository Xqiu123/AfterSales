package car

import (
	"as/dao"
	. "as/handler"
	"as/pkg/errno"
	"as/service"
	"github.com/gin-gonic/gin"
)

// Update 处理更新汽车信息的请求
func Update(c *gin.Context) {
	var req dao.CarModel
	if err := c.BindJSON(&req); err != nil {
		SendError(c, errno.BadRequest, nil, err.Error(), GetLine())
		return
	}

	if err := service.UpdateCar(&req); err != nil {
		SendError(c, errno.BadRequest, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, nil, "Car updated successfully")
}
