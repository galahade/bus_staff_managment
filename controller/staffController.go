package controller

import (
	"net/http"
	. "github.com/galahade/bus_staff_managment/service"
	"gopkg.in/gin-gonic/gin.v1"
)

func AddStaff(c *gin.Context)  {
	staff, err := fillStaffModelByRequest(c)
	setCORSHeader(c);
	if err != nil {
		BadRequestResponse(c, err)
		return
	}
	err = CreateStaff(&staff)
	if err != nil {
		BadRequestResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, wrapperResponseJson("api.bus.com/staffs", "staffs",[]StaffModel{
		staff,
	}));

}

func ShowDictionaryByType(c *gin.Context) {
	setCORSHeader(c)
	dicType := c.Query("type")

	dicModels, err := FetchDictionariesByType(dicType)
	if (err != nil) {
		BadRequestResponse(c, err)
	}
	c.JSON(http.StatusOK, wrapperResponseJson("api.bus.com/dictionary", "dictionaries", dicModels))

}

func ShowStaffs(c *gin.Context) {
	setCORSHeader(c)
	staffType := c.Query("type")

	switch staffType {
	case "driver":
		getDriverByType(c)
	case "support":
		getSupportStaffs(c)
	default:
		c.JSON(http.StatusOK,wrapperResponseJson("api.bus.com/staffs", "staffs", GetAllStaffs()));
	}

}

func GetStaffByID (c *gin.Context) {
	setCORSHeader(c)
	id := c.Param("id")

	staffModel, ok := FetchStaffByID(id)
	var staffModels []StaffModel
	if ok {
		staffModels = append(staffModels, staffModel)
		c.JSON(http.StatusOK,wrapperResponseJson("api.bus.com/staffs", "staffs", staffModels));
	} else {
		c.JSON(http.StatusNotFound,wrapperResponseJson("api.bus.com/staffs", "staffs", staffModels));
	}
}

func getDriverByType(c *gin.Context) {
	switch c.Query("driverType") {
	case "qualified":
		c.JSON(http.StatusOK,wrapperResponseJson("api.bus.com/staffs", "staffs", GetAllQualifiedDrivers()));
	case "internship":
		c.JSON(http.StatusOK,wrapperResponseJson("api.bus.com/staffs", "staffs", GetAllInternshipDrivers()));
	default:
		c.JSON(http.StatusOK,wrapperResponseJson("api.bus.com/staffs", "staffs", GetAllDrivers()));
	}
}

func getSupportStaffs(c *gin.Context)  {
	c.JSON(http.StatusOK,wrapperResponseJson("api.bus.com/staffs", "staffs", GetSupportStaffs()));
}

func fillStaffModelByRequest(c *gin.Context) (StaffModel, error) {
	id := c.Param("id")

	requestWrapper := map[string]*StaffModel {
	}
	err := c.Bind(&requestWrapper)
	if err == nil {
		if id != ""  {
			requestWrapper["staff"].ID = id
		}
	}
	return *requestWrapper["staff"], err
}




