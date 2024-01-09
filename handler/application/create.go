package application

import (
	. "as/handler"
	"as/pkg/errno"

	"as/dao"
	"as/service"
	"github.com/gin-gonic/gin"
)

// Create
// @Summary 申请试驾或维修 api
// @Tags application
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "token 用户令牌"
// @Param object body dao.ApplicationModel true "application"
// @Router /application [post]
func Create(c *gin.Context) {
	var req dao.ApplicationModel
	if err := c.BindJSON(&req); err != nil {
		SendError(c, errno.BadRequest, nil, err.Error(), GetLine())
		return
	}

	if err := service.CreateApplication(&req); err != nil {
		SendError(c, err, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, nil, "Car created successfully")
}
