package controller

import (
	"net/http"
	"github.com/galahade/bus_staff_managment/service"
	"gopkg.in/gin-gonic/gin.v1"
)

func ShowDrivers(c *gin.Context) {
	setCORSHeader(c)
	switch c.Query("driverType") {
	case "qualified":
		c.JSON(http.StatusOK, wrapperStaff(service.GetAllQualifiedDrivers()))
	case "internship":
		c.JSON(http.StatusOK, wrapperStaff(service.GetAllInternshipDrivers()))
	default:
		c.JSON(http.StatusOK, wrapperStaff(service.GetAllDrivers()))
	}
}

func GetDriverByStaffID (c *gin.Context) {
	setCORSHeader(c)
	sid := c.Param("sid")

	staffModel, ok := service.FetchDriverBySID(sid)
	var staffModels []service.StaffModel
	if ok {

		staffModels = append(staffModels, staffModel)
		c.JSON(http.StatusOK, wrapperStaff(staffModels))
	} else {
		c.JSON(http.StatusNotFound,wrapperStaff(staffModels))
	}
}

func wrapperStaff(staffs []service.StaffModel) RESTWrapper {
	wrapper := NewWrapper();
	wrapper.setSelf("api.bus.com/drivers")
	wrapper.setData("drivers", staffs)
	return *wrapper
}




