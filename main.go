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

	router.GET("/drivers", ShowDrivers)
	router.GET("/drivers/:sid", GetDriverByStaffID)

	router.OPTIONS("/buses", HandleOptionsRequest)
	router.OPTIONS("/buses/:license", HandleOptionsRequest)
	router.OPTIONS("/chargeRecords", HandleOptionsRequest)

	router.GET("/buses", ShowAllBuses)
	router.GET("/buses/:license", GetBusByLicense)
	router.POST("/buses", AddBus)
	router.PUT("/buses/:license", PutBus)

	router.GET("/brands", ShowAllBusBrands)

	router.POST("/chargeRecords", AddChargeRecord)
	router.GET("/chargeRecords", ShowAllChargeRecord)

	log.Fatal(router.Run(":8000"))
}

