package car

import (
	. "as/handler"
	"as/pkg/errno"
	"as/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func DeleteCar(c *gin.Context) {
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
