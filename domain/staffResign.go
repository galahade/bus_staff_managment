package domain

import (
	"github.com/jinzhu/gorm"
	"github.com/pborman/uuid"
	"time"
)

type StaffResign struct {
	Domain
	StaffID          string   `gorm:"ForeignKey:staff_id"`
	Staff            Staff
	ResignDate       time.Time
	ResignReason     string
}

func (*StaffResign) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.New())
	return nil
}

func (staffResign *StaffResign) Create() error {
	return insertDomain(staffResign)
}

func (staffResign *StaffResign) Update() error {
	UpdateCreateDate(staffResign)
	tempDB := Gdb.Save(staffResign)
	if tempDB.Error != nil {
		return tempDB.Error
	}
	return nil
}

func (staffResign StaffResign) QueryAll() []StaffResign {
	resignStaffs := []StaffResign{}
	Gdb.Order("counting_date").Find(&resignStaffs);
	for i, _ := range resignStaffs {
		Gdb.Model(resignStaffs[i]).Related(&resignStaffs[i].Staff)
	}
	return resignStaffs
}

func (StaffResign) Query(query map[string]interface{}) []StaffResign {
	resignStaffs := []StaffResign{}
	Gdb.Where(query).Find(&resignStaffs)
	for i, _ := range resignStaffs {
		Gdb.Model(resignStaffs[i]).Related(&resignStaffs[i].Staff)
	}
	return resignStaffs
}
