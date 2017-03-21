package controller

import (
	. "github.com/galahade/bus_staff_managment/service"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

func ShowAllBusBrands (c *gin.Context) {
	setCORSHeader(c);
	c.JSON(http.StatusOK, wrapperBusBrand(GetAllBusBrands()));
}


func wrapperBusBrand(busBrands []BusBrandModel) RESTWrapper {
	wrapper := NewWrapper();
	wrapper.setSelf("api.bus.com/busBrand")
	wrapper.setData("brands", busBrands)
	return *wrapper
}