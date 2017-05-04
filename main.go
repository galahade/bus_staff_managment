package main

import (
	"log"
	"github.com/galahade/bus_staff_managment/session"
	. "github.com/galahade/bus_staff_managment/controller"
	"gopkg.in/gin-gonic/gin.v1"
)

var globalSessions *session.Manager

func init() {
	globalSessions, _ = session.NewManager("memory", "gosessionid", 3600)
	//	go globalSessions.GC()

}

func main() {

	router := gin.Default()

	router.GET("/staffs", ShowStaffs)
	router.GET("/staffs/:id", GetStaffByID)
	router.POST("/staffs", AddStaff)
	router.GET("/dictionaries", ShowDictionaryByType)

	router.OPTIONS("/staffs", HandleOptionsRequest)
	router.OPTIONS("/staffs/:id", HandleOptionsRequest)
	router.OPTIONS("/dictionaries/:type", ShowDictionaryByType)
	router.OPTIONS("/buses", HandleOptionsRequest)
	router.OPTIONS("/buses/:id", HandleOptionsRequest)
	router.OPTIONS("/chargeRecords", HandleOptionsRequest)
	router.OPTIONS("/chargeRecords/:id", HandleOptionsRequest)
	router.OPTIONS("/lineFareIncomes/:id", HandleOptionsRequest)
	router.OPTIONS("/staffResigns", HandleOptionsRequest)

	router.GET("/buses", ShowAllBuses)
	router.GET("/buses/:license", GetBusByLicense)
	router.POST("/buses", AddBus)
	router.PUT("/buses/:id", PutBus)

	router.GET("/brands", ShowAllBusBrands)

	router.POST("/chargeRecords", AddChargeRecord)
	router.GET("/chargeRecords", ShowChargeRecords)
	router.PUT("/chargeRecords/:id", PutChargeRecord)

	router.POST("/lineFareIncomes", AddLineFareIncome)
	router.GET("/lineFareIncomes", ShowLineFareIncomes)
	router.PUT("/lineFareIncomes/:id", PutLineFareIncome)

	router.POST("/staffResigns", AddStaffResign)


	log.Fatal(router.Run(":8000"))
}

