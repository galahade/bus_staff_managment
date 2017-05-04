package controller

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	. "github.com/galahade/bus_staff_managment/service"
)

func AddStaffResign(c *gin.Context) {
	staffResignModel, err := fillStaffResignModelByRequest(c)
	setCORSHeader(c);
	if err != nil {
		BadRequestResponse(c, err)
		return
	}
	err = Resign(&staffResignModel)
	if err != nil {
		BadRequestResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, wrapperResponseJson("api.bus.com/staffResign", "staffResigns",[]StaffResignModel{
		staffResignModel,
	}));
}

func fillStaffResignModelByRequest(c *gin.Context) (StaffResignModel, error) {
	id := c.Param("id")

	requestWrapper := map[string]*StaffResignModel {
	}
	err := c.Bind(&requestWrapper)
	if err == nil {
		if id != ""  {
			requestWrapper["staffResign"].ID = id
		}
	}
	return *requestWrapper["staffResign"], err
}
