package service

import (
	"testing"
	"github.com/stretchr/testify/assert"
	. "github.com/galahade/bus_staff_managment/domain"
)

func TestChangeBus(t *testing.T) {
	Gdb.Debug()
	Gdb.LogMode(true)
	busModel := BusModel{
		License:"test45",
		CustomID:"dfdfdf",
		BrandID: "239e7938-9f76-49f9-a290-3cbacc52ab62",
		RegisterDate: "2017-03-29",
		VehicleIDNumber: "testVIN",
		EngineNo: "testEngineNo",
	}
	err := ChangeBus(&busModel)
	assert.Empty(t, err)


}
