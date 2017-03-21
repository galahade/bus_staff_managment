package domain

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestBusBrand_Create(t *testing.T) {
	busBrand := BusBrand{
		Name:"长安",
		Model:"SC6833BEV",
	}
	(&busBrand).Create()

}

func TestBus_Create(t *testing.T) {
	Gdb.Debug()
	Gdb.LogMode(true)
	bus1 := Bus{
		BusLicense:"12341",
		CustomId:"123123",
		BrandID: "1c4f1929-298d-4301-b4c8-130fbb3208a6",
		RegisterDate:time.Now(),
		VehicleIDNumber: "testVIN",
		EngineNo: "testEngineNo",
	}
	(&bus1).Create()

	bus2 := Bus{BusLicense:"12341"}
	(&bus2).QueryByLicense();
	assert.NotEmpty(t, bus2.ID)
	assert.NotEmpty(t, bus2.CreatedAt)
	assert.NotEmpty(t, bus2.CustomId)

	Gdb.Unscoped().Delete(&bus2)


}

func TestBus_QueryByLicense(t *testing.T) {
	Gdb.Debug()
	Gdb.LogMode(true)
	bus := &Bus{BusLicense:"322D8"}

	bus.QueryByLicense()
	assert.NotEmpty(t, bus.ID)
	assert.NotEmpty(t, bus.CreatedAt)
	assert.NotEmpty(t, bus.CustomId)
	assert.NotEmpty(t, bus.BusBrand.ID)
	assert.NotEmpty(t, bus.BusBrand.Name)
	assert.NotEmpty(t, bus.BusBrand.Model)
	assert.NotEmpty(t, bus.BusBrand.CreatedAt)
}

func TestBus_QueryAll(t *testing.T) {
	Gdb.LogMode(true)
	bus := Bus{}
	buses := bus.QueryAll()
	for _, bus := range buses {
		assert.NotEmpty(t, bus.ID)
		assert.NotEmpty(t, bus.BusBrand.ID)
		assert.NotEmpty(t, bus.BusBrand.Name)
		assert.NotEmpty(t, bus.BusBrand.Model)
		assert.NotEmpty(t, bus.BusBrand.CreatedAt)
	}
}

