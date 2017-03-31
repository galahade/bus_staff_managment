package domain

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestChargeRecord_Create(t *testing.T) {
	setGORMShowSQL()
	chargeRecord := ChargeRecord{
		BusID: "0cc2b5b3-ba9d-466b-83cd-e15fcaaf545f",
		RecordDate: time.Now(),
		RecordStaffID: "1e8ff860-1dd2-11b2-9660-60f81dacbf60",
		Mileage: 30000,
		ChargedTWH: 70,
		RemainPercent: 40,
		FinalPercent: 100,
	}
	(&chargeRecord).Create()

	assert.NotEmpty(t, chargeRecord.ID)
	assert.NotEmpty(t, chargeRecord.CreatedAt)
	assert.Nil(t, chargeRecord.DeletedAt)

	Gdb.Unscoped().Delete(&chargeRecord)
}

func TestChargeRecord_QueryAll(t *testing.T) {
	setGORMShowSQL()
	chargeRecord := ChargeRecord{}

	chargeRecords := chargeRecord.QueryAll()

	for _, chargeRecord := range chargeRecords {
		assert.NotEmpty(t, chargeRecord.ID)
		assert.NotEmpty(t, chargeRecord.ChargedTWH)
		assert.NotEmpty(t, chargeRecord.BusID)
		assert.NotEmpty(t, chargeRecord.RecordStaffID)
		assert.NotEmpty(t, chargeRecord.FinalPercent)
		assert.NotEmpty(t, chargeRecord.Bus)
		assert.NotEmpty(t, chargeRecord.Bus.ID)
		assert.NotEmpty(t, chargeRecord.Bus.CreatedAt)
		assert.NotEmpty(t, chargeRecord.Bus.CustomId)
		assert.NotEmpty(t, chargeRecord.Bus.EngineNo)
		assert.NotEmpty(t, chargeRecord.Bus.PersonsCapacity)
		assert.NotEmpty(t, chargeRecord.Bus.RegisterDate)
		assert.NotEmpty(t, chargeRecord.RecordStaff)
		assert.NotEmpty(t, chargeRecord.RecordStaff.ID)
		assert.NotEmpty(t, chargeRecord.RecordStaff.CreatedAt)
		assert.NotEmpty(t, chargeRecord.RecordStaff.DriverType)
		assert.NotEmpty(t, chargeRecord.RecordStaff.Department)
		assert.NotEmpty(t, chargeRecord.RecordStaff.EmergencyContact)
		assert.NotEmpty(t, chargeRecord.RecordStaff.EmergencyContactPhone)
		assert.NotEmpty(t, chargeRecord.RecordStaff.EmergencyContactRelation)
		assert.NotEmpty(t, chargeRecord.RecordStaff.FirstOnboardTime)
		assert.NotEmpty(t, chargeRecord.RecordStaff.JobType)
		assert.NotEmpty(t, chargeRecord.RecordStaff.Name)
	}
}
