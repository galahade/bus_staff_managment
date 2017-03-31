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

func ShowAllChargeRecord(c *gin.Context) {
	setCORSHeader(c);
	c.JSON(http.StatusOK, wrapperChargeRecord(GetAllChargeRecord()));
}

func wrapperChargeRecord(chargeRecords []ChargeRecordModel) RESTWrapper {
	wrapper := NewWrapper();
	wrapper.setSelf("api.bus.com/chargeRecord")
	wrapper.setData("chargeRecords", chargeRecords)
	return *wrapper
}

func fillChargeRecordModelByRequest(c *gin.Context) (ChargeRecordModel, error) {
	var chargeRecordModel ChargeRecordModel
	requestWrapper := map[string]ChargeRecordModel {
		"chargeRecord": chargeRecordModel,
	}
	err := c.Bind(&requestWrapper)
	return requestWrapper["chargeRecord"], err
}
