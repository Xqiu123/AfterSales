package application

import (
	. "as/handler"
	"as/pkg/errno"

	"as/dao"
	"as/service"
	"github.com/gin-gonic/gin"
)

// Create 处理增加汽车信息的请求
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