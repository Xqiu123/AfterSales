package car

import (
	. "as/handler"
	"as/service"
	"github.com/gin-gonic/gin"
)

// List 处理列出汽车信息的请求
func List(c *gin.Context) {
	brand := c.Param("brand")

	cars, err := service.GetAllCarsByBrand(brand)
	if err != nil {
		SendError(c, err, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, nil, cars)
}
