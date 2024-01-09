package application

import (
	"as/dao"
	. "as/handler"
	"as/service"
	"github.com/gin-gonic/gin"
)

// ListAllApplication
// @Summary list所有试驾和维修application api
// @Tags application
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "token 用户令牌"
// @Success 200 {object} ListAllApplicationResponse
// @Router /application/list/all [get]
func ListAllApplication(c *gin.Context) {
	driveApplications, repairApplications, err := service.GetAllApplications()
	if err != nil {
		SendError(c, err, nil, err.Error(), GetLine())
		return
	}

	response := ListAllApplicationResponse{
		Drive:  driveApplications,
		Repair: repairApplications,
	}

	SendResponse(c, nil, response)
}

type ListAllApplicationResponse struct {
	Drive  []*dao.ApplicationInfo `json:"drive"`
	Repair []*dao.ApplicationInfo `json:"repair"`
}
