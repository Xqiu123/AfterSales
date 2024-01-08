package car

import (
	. "as/handler"
	"as/pkg/errno"

	"as/dao"
	"as/service"
	"github.com/gin-gonic/gin"
)

// CreateCar 处理增加汽车信息的请求
func CreateCar(c *gin.Context) {
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
