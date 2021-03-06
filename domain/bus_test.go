package domain

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestBusBrand_Create(t *testing.T) {
	setGORMShowSQL()
	busBrand := BusBrand{
		Name:"长安",
		Model:"SC6833BEV",
	}
	(&busBrand).Create()

	assert.NotEmpty(t, busBrand.ID)
	assert.NotEmpty(t, busBrand.CreatedAt)
	assert.Nil(t, busBrand.DeletedAt)

	Gdb.Unscoped().Delete(&busBrand)
}

func TestBus_Create(t *testing.T) {
	setGORMShowSQL()
	bus := Bus{
		BusLicense:"12341",
		CustomId:"123123",
		BrandID: "1c4f1929-298d-4301-b4c8-130fbb3208a6",
		RegisterDate:time.Now(),
		VehicleIDNumber: "testVIN",
		EngineNo: "testEngineNo",
	}
	(&bus).Create()

	assert.NotEmpty(t, bus.ID)
	assert.NotEmpty(t, bus.CreatedAt)
	assert.NotEmpty(t, bus.CustomId)

	Gdb.Unscoped().Delete(&bus)
}

func TestBus_QueryByLicense(t *testing.T) {
	setGORMShowSQL()
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
	setGORMShowSQL()
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

func TestBus_UpdateCreateDate(t *testing.T) {
	setGORMShowSQL()
	bus := &Bus{
		Domain: Domain{
			ID: "0cc2b5b3-ba9d-466b-83cd-e15fcaaf545f",
		},
		BusLicense:"75090",
		CustomId:"75090",
		VehicleIDNumber:"LS4AJB5XXFG002522",
		EngineNo:"IS5005JJ20150S0174",
		PersonsCapacity:57,
		BrandID:"1c4f1929-298d-4301-b4c8-130fbb3208a6",
	}
	UpdateCreateDate(bus)
	assert.NotEmpty(t, bus.ID)
	assert.NotEmpty(t, bus.CreatedAt)
	assert.NotEmpty(t, bus.CustomId)
	assert.NotEmpty(t, bus.VehicleIDNumber)
	assert.NotEmpty(t, bus.EngineNo)
	assert.NotEmpty(t, bus.PersonsCapacity)
	assert.NotEmpty(t, bus.BrandID)
}

