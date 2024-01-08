package application

import (
	. "as/handler"
	"as/service"
	"github.com/gin-gonic/gin"
)

func GetMyOrders(c *gin.Context) {
	userId := c.MustGet("userId").(int)

	orders, err := service.GetMyApplications(userId)
	if err != nil {
		SendError(c, err, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, nil, orders)
}
