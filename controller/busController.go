package controller

import (
	. "github.com/galahade/bus_staff_managment/service"
	"net/http"
	"gopkg.in/gin-gonic/gin.v1"
	"log"
)

func GetBusByLicense(c *gin.Context) {
	setCORSHeader(c)
	license := c.Param("license")

	staffModel, ok := FetchBusByLicense(license)
	var busModels []BusModel
	if ok {
		busModels = append(busModels, staffModel)
		c.JSON(http.StatusOK, wrapperBus(busModels))
	} else {
		c.JSON(http.StatusNotFound,wrapperBus(busModels))
	}
}

func ShowAllBuses(c *gin.Context) {
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
	err = CreateBus(&busModel)
	if err != nil {
		BadRequestResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, wrapperBus([]BusModel{
		busModel,
	}));
}

func PutBus(c *gin.Context) {
	busModel, err := fillBusModelByRequest(c)
	setCORSHeader(c);
	if err != nil {
		BadRequestResponse(c, err)
		return
	}
	err = ChangeBus(&busModel)
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
	id := c.Param("id")
	var busModel BusModel
	requestWrapper := map[string]*BusModel{
	}
	err := c.Bind(&requestWrapper)

	if err == nil {
		if id != ""  {
			requestWrapper["bus"].ID = id
			log.Printf("%#v", requestWrapper["bus"])
			return *requestWrapper["bus"], nil
		}
	}
	return busModel, err
}

