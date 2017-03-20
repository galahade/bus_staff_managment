package controller

import (
	. "github.com/galahade/bus_staff_managment/service"
	"net/http"
	"gopkg.in/gin-gonic/gin.v1"
)

func ShowAllBuses (c *gin.Context) {
	setCORSHeader(c);
	c.JSON(http.StatusOK, wrapperBus(GetAllBuses()));
}

func AddBus(c *gin.Context) {
	busModel, err := fillBusModelByRequest(c)
	setCORSHeader(c);

	if err != nil {
		BadRequestResponse(c, err)
		return
	}
	err = CreateNewBus(&busModel)
	if err != nil {
		BadRequestResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, wrapperBus([]BusModel{
		busModel,
	}));
}

func wrapperBus(buses []BusModel) RESTWrapper {
	wrapper := NewWrapper();
	wrapper.setSelf("api.bus.com/buses")
	wrapper.setData("buses", buses)
	return *wrapper
}

func fillBusModelByRequest(c *gin.Context) (BusModel, error){

	var busModel BusModel
	requestWrapper := map[string]BusModel{
		"bus": busModel,
	}
	err := c.Bind(&requestWrapper)

	if err == nil {
		return requestWrapper["bus"], nil
	}
	return busModel, err
}

