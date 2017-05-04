package controller

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	. "github.com/galahade/bus_staff_managment/service"

)

func AddLineFareIncome(c *gin.Context) {
	lineFareIncomeModel, err := filLineFareIncomeModelByRequest(c)
	setCORSHeader(c);
	if err != nil {
		BadRequestResponse(c, err)
		return
	}
	err = CreateLineFareIncome(&lineFareIncomeModel)
	if err != nil {
		BadRequestResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, wrapperLineFareIncome([]LineFareIncomeModel{
		lineFareIncomeModel,
	}));
}

func PutLineFareIncome(c *gin.Context) {
	lineFareIncomeModel, err := filLineFareIncomeModelByRequest(c)
	setCORSHeader(c);
	if err != nil {
		BadRequestResponse(c, err)
		return
	}
	err = ChangeLineFareIncome(&lineFareIncomeModel)
	if err != nil {
		BadRequestResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, wrapperLineFareIncome([]LineFareIncomeModel{
		lineFareIncomeModel,
	}));
}

func ShowLineFareIncomes(c *gin.Context) {
	setCORSHeader(c);

	query, ok := assembleQuery(c)
	if ok {
		c.JSON(http.StatusOK, wrapperLineFareIncome(GetLineFareIncome(query)))
		return
	}
	c.JSON(http.StatusOK, wrapperLineFareIncome(GetAllLineFareIncome()));
}

func wrapperLineFareIncome(chargeRecords []LineFareIncomeModel) RESTWrapper {
	wrapper := NewWrapper();
	wrapper.setSelf("api.bus.com/lineFareIncome")
	wrapper.setData("lineFareIncomes", chargeRecords)
	return *wrapper
}

func filLineFareIncomeModelByRequest(c *gin.Context) (LineFareIncomeModel, error) {
	id := c.Param("id")

	requestWrapper := map[string]*LineFareIncomeModel {
	}
	err := c.Bind(&requestWrapper)
	if err == nil {
		if id != ""  {
			requestWrapper["lineFareIncome"].ID = id
		}
	}
	return *requestWrapper["lineFareIncome"], err
}


