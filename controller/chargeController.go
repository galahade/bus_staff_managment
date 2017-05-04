package controller

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	. "github.com/galahade/bus_staff_managment/service"

)

func AddChargeRecord(c *gin.Context) {
	chargeRecordModel, err := fillChargeRecordModelByRequest(c)
	setCORSHeader(c);
	if err != nil {
		BadRequestResponse(c, err)
		return
	}
	err = CreateChargeRecord(&chargeRecordModel)
	if err != nil {
		BadRequestResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, wrapperChargeRecord([]ChargeRecordModel{
		chargeRecordModel,
	}));
}

func PutChargeRecord(c *gin.Context) {
	chargeRecordModel, err := fillChargeRecordModelByRequest(c)
	setCORSHeader(c);
	if err != nil {
		BadRequestResponse(c, err)
		return
	}
	err = ChangeChargeRecord(&chargeRecordModel)
	if err != nil {
		BadRequestResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, wrapperChargeRecord([]ChargeRecordModel{
		chargeRecordModel,
	}));
}

func ShowChargeRecords(c *gin.Context) {
	setCORSHeader(c);

	query, ok := assembleQuery(c)
	if ok {
		c.JSON(http.StatusOK, wrapperChargeRecord(GetChargeRecord(query)))
		return
	}
	c.JSON(http.StatusOK, wrapperChargeRecord(GetAllChargeRecord()));
}

func wrapperChargeRecord(chargeRecords []ChargeRecordModel) RESTWrapper {
	wrapper := NewWrapper();
	wrapper.setSelf("api.bus.com/chargeRecord")
	wrapper.setData("chargeRecords", chargeRecords)
	return *wrapper
}

func fillChargeRecordModelByRequest(c *gin.Context) (ChargeRecordModel, error) {
	id := c.Param("id")

	requestWrapper := map[string]*ChargeRecordModel {
	}
	err := c.Bind(&requestWrapper)
	if err == nil {
		if id != ""  {
			requestWrapper["chargeRecord"].ID = id
		}
	}
	return *requestWrapper["chargeRecord"], err
}

