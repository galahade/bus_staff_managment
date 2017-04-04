package domain

import (
	"time"
	"github.com/pborman/uuid"
	"github.com/jinzhu/gorm"
)

type ChargeRecord struct {
	Domain
	RecordStaffID string
	RecordStaff   Staff
	BusID         string
	Bus           Bus
	RecordDate    time.Time
	Mileage       float32
	ChargedTWH    float32       `gorm:"column:charged_TWH"`
	RemainPercent float32
	FinalPercent  float32
}

func (*ChargeRecord) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.New())
	return nil
}

func (chargeRecord *ChargeRecord) Create() error {
	return insertDomain(chargeRecord)
}

func (chargeRecord *ChargeRecord) Update() error {
	UpdateCreateDate(chargeRecord)
	tempDB := Gdb.Save(chargeRecord)
	if tempDB.Error != nil {
		return tempDB.Error
	}
	return nil
}

func (chargeRecord *ChargeRecord) QueryByID() error {
	//gdb.Where("bus_license = ?", bus.BusLicense).First(bus).Related(&bus.Brand)
	//gdb.Preload("Brand").First(bus, "bus_license = ?", bus.BusLicense)//.Model(bus).Related(&bus.Brand)
	Gdb.First(chargeRecord).Related(&chargeRecord.Bus)
	if err := checkQueryFirstNotNil(chargeRecord); err != nil {
		return err
	}
	return nil
}

func (ChargeRecord) Query(query map[string]interface{}) []ChargeRecord {
	chargeRecords := []ChargeRecord{}
	Gdb.Where(query).Find(&chargeRecords)
	for i, _ := range chargeRecords {
		Gdb.Model(chargeRecords[i]).Related(&chargeRecords[i].Bus).Related(&chargeRecords[i].RecordStaff, "RecordStaffID")
	}
	return chargeRecords
}

func (ChargeRecord) QueryAll() []ChargeRecord {
	chargeRecords := []ChargeRecord{}
	Gdb.Order("record_date desc").Find(&chargeRecords);
	for i, _ := range chargeRecords {
		Gdb.Model(chargeRecords[i]).Related(&chargeRecords[i].Bus).Related(&chargeRecords[i].RecordStaff, "RecordStaffID")
	}
	return chargeRecords
}

