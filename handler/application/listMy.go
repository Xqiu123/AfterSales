package application

import (
	"as/dao"
	. "as/handler"
	"as/service"
	"github.com/gin-gonic/gin"
)

// ListMyApplication
// @Summary list我的试驾和维修application api
// @Tags application
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "token 用户令牌"
// @Success 200 {object} ListMyApplicationResponse
// @Router /application/list [get]
func ListMyApplication(c *gin.Context) {
	userId := c.MustGet("userId").(int)

	applications, err := service.GetApplicationsByUserId(userId)
	if err != nil {
		SendError(c, err, nil, err.Error(), GetLine())
		return
	}
	driveApplications := make([]*dao.ApplicationModel, 0)
	repairApplications := make([]*dao.ApplicationModel, 0)
	for _, application := range applications {
		if application.Class == "drive" {
			driveApplications = append(driveApplications, application)
		} else {
			repairApplications = append(repairApplications, application)
		}
	}
	response := ListMyApplicationResponse{
		Drive:  driveApplications,
		Repair: repairApplications,
	}

	SendResponse(c, nil, response)
}

type ListMyApplicationResponse struct {
	Drive  []*dao.ApplicationModel `json:"drive"`
	Repair []*dao.ApplicationModel `json:"repair"`
}
