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
	/*
	busBrand := BusBrand{
		//Domain.ID:"239e7938-9f76-49f9-a290-3cbacc52ab62",
		Name:"长安",
		Model:"SC6833BEV",
	}
	*/
	bus1 := Bus{
		BusLicense:"233D8",
		CustomId:"338",
	//	Brand: busBrand,
		RegisterDate:time.Now(),
		VIN:"testVIN",
		EngineNo:"testEngineNo",
	}
	(&bus1).Create()

	bus2 := Bus{BusLicense:"322D8"}
	(&bus2).QueryByLicense();
	assert.NotEmpty(t, bus2.ID)
	assert.NotEmpty(t, bus2.CreatedAt)
	assert.NotEmpty(t, bus2.CustomId)

	gdb.Unscoped().Delete(&bus2)


}

func TestBus_QueryByLicense(t *testing.T) {
	gdb.Debug()
	gdb.LogMode(true)
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
	gdb.LogMode(true)
	bus := Bus{}
	buses := bus.QueryAll()
	for _, bus := range buses {
		assert.NotEmpty(t, bus.ID)
		assert.NotEmpty(t, bus.BusBrand.ID)
	}
}

