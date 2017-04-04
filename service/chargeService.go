package service

import (
	"time"
	. "github.com/galahade/bus_staff_managment/domain"
)

var defaultChargeRecord = ChargeRecord{}

type ChargeRecordModel struct {
	ID            string       `json:"id"`
	Bus           BusModel     `json:"bus"`
	RecordDate    string       `json:"recordDate"`
	RecordStaff   StaffModel   `json:"recordStaff"`
	Mileage       float32      `json:"mileage"`
	ChargedTWH    float32      `json:"chargedTWH"`
	RemainPercent float32      `json:"remainPercent"`
	FinalPercent  float32      `json:"finalPercent"`
}

func CreateChargeRecord(chargeRecordModel *ChargeRecordModel) error {
	chargeRecord, err := chargeRecordModel.toDomain()
	err = chargeRecord.Create()
	chargeRecordModel.ID = chargeRecord.ID
	return err
}

func GetChargeRecord(query map[string]interface{}) []ChargeRecordModel {
	return fillFromChargeRecordDomains(defaultChargeRecord.Query(query))
}

func GetAllChargeRecord() []ChargeRecordModel {
	return fillFromChargeRecordDomains(defaultChargeRecord.QueryAll())
}

func ChangeChargeRecord(chargeRecordModel *ChargeRecordModel) error {
	chargeRecord, err := chargeRecordModel.toDomain()
	err = chargeRecord.Update()
	return err
}

func fillFromChargeRecordDomains(chargeRecords []ChargeRecord) (chargeRecordModels []ChargeRecordModel) {
	for _, chargeRecord := range chargeRecords {
		chargeRecordModel := new(ChargeRecordModel)
		chargeRecordModel.fillFromDomain(chargeRecord)
		chargeRecordModels = append(chargeRecordModels, *chargeRecordModel)
	}
	return
}

func (chargeRecordModel ChargeRecordModel) toDomain() (chargeRecord *ChargeRecord, err error) {
	chargeRecord = new(ChargeRecord)
	chargeRecord.ID = chargeRecordModel.ID
	chargeRecord.BusID = chargeRecordModel.Bus.ID
	chargeRecord.RecordDate, err = time.Parse(DateString, chargeRecordModel.RecordDate)
	chargeRecord.RecordStaffID = chargeRecordModel.RecordStaff.ID
	chargeRecord.Mileage = chargeRecordModel.Mileage
	chargeRecord.ChargedTWH = chargeRecordModel.ChargedTWH
	chargeRecord.RemainPercent = chargeRecordModel.RemainPercent
	chargeRecord.FinalPercent = chargeRecordModel.FinalPercent

	if serviceErr, ok := hasError(err); ok {
		return nil, serviceErr
	}
	return chargeRecord, nil
}

func (chargeRecordModel *ChargeRecordModel) fillFromDomain(chargeRecord ChargeRecord) {
	chargeRecordModel.ID = chargeRecord.ID
	chargeRecordModel.Bus.fillFromDomain(chargeRecord.Bus)
	chargeRecordModel.RecordDate = chargeRecord.RecordDate.Format(DateString)
	chargeRecordModel.RecordStaff.fillFromDomain(chargeRecord.RecordStaff)
	chargeRecordModel.Mileage = chargeRecord.Mileage
	chargeRecordModel.ChargedTWH = chargeRecord.ChargedTWH
	chargeRecordModel.RemainPercent = chargeRecord.RemainPercent
	chargeRecordModel.FinalPercent = chargeRecord.FinalPercent
}
