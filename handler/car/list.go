package car

import (
	. "as/handler"
	"as/service"
	"github.com/gin-gonic/gin"
)

// List
// @Summary list车辆 api
// @Tags car
// @Accept application/json
// @Produce application/json
// @Param brand query int false "brand"
// @Router /car/list [get]
func List(c *gin.Context) {
	brand := c.DefaultQuery("brand", "")

	cars, err := service.GetAllCarsByBrand(brand)
	if err != nil {
		SendError(c, err, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, nil, cars)
}
