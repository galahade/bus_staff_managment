package domain

import (
	"time"
	"github.com/jinzhu/gorm"
	"github.com/pborman/uuid"
)

type LineFareIncome struct {
	Domain
	LineNo           int
	CarryingAmount   float32
	ActualAmount     float32
	WornCoinAmount   float32
	BusNumbers       int
	CountingDate     time.Time
	CountingStaff1ID string  `gorm:"ForeignKey:CountingStaff1ID"`
	CountingStaff1   Staff
	CountingStaff2ID string  `gorm:"ForeignKey:CountingStaff2ID"`
	CountingStaff2   Staff
	RecordStaffID    string
	RecordStaff      Staff   `gorm:"ForeignKey:RecordStaffID"`
}

func (*LineFareIncome) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.New())
	return nil
}

func (fareIncome *LineFareIncome) Create() error {
	return insertDomain(fareIncome)
}

func (fareIncome *LineFareIncome) Update() error {
	UpdateCreateDate(fareIncome)
	tempDB := Gdb.Save(fareIncome)
	if tempDB.Error != nil {
		return tempDB.Error
	}
	return nil
}

func (fareIncome LineFareIncome) QueryAll() []LineFareIncome {
	incomes := []LineFareIncome{}
	Gdb.Order("counting_date").Find(&incomes);
	for i, _ := range incomes {
		Gdb.Model(incomes[i]).Related(&incomes[i].CountingStaff1)
		Gdb.Model(incomes[i]).Related(&incomes[i].CountingStaff2)
		Gdb.Model(incomes[i]).Related(&incomes[i].RecordStaff)
	}
	return incomes
}

func (LineFareIncome) Query(query map[string]interface{}) []LineFareIncome {
	lineFareIncomes := []LineFareIncome{}
	Gdb.Where(query).Find(&lineFareIncomes)
	for i, _ := range lineFareIncomes {
		Gdb.Model(lineFareIncomes[i]).Related(&lineFareIncomes[i].CountingStaff1).Related(&lineFareIncomes[i].CountingStaff2).Related(&lineFareIncomes[i].RecordStaff)
	}
	return lineFareIncomes
}