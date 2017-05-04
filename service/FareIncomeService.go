package service

import (
	"time"
	. "github.com/galahade/bus_staff_managment/domain"
)

var defaultLineFareIncomeModel = LineFareIncome{}

type LineFareIncomeModel struct {
	ID             string     `json:"id"`
	LineNo         int        `json:"lineNo"`
	CarryingAmount float32    `json:"carryingAmount"`
	ActualAmount   float32    `json:"actualAmount"`
	WornCoinAmount float32    `json:"wornCoinAmount"`
	BusNumbers     int        `json:"busNumbers"`
	CountingDate   string     `json:"countingDate"`
	CountingStaff1 StaffModel      `json:"countingStaff1"`
	CountingStaff2 StaffModel      `json:"countingStaff2"`
	RecordStaff    StaffModel      `json:"recordStaff"`
}

func CreateLineFareIncome(chargeRecordModel *LineFareIncomeModel) error {
	lineFareIncome, err := chargeRecordModel.toDomain()
	err = lineFareIncome.Create()
	chargeRecordModel.ID = lineFareIncome.ID
	return err
}

func GetLineFareIncome(query map[string]interface{}) []LineFareIncomeModel {
	return fillFromLineFareIncomeDomains(defaultLineFareIncomeModel.Query(query))
}

func GetAllLineFareIncome() []LineFareIncomeModel {
	return fillFromLineFareIncomeDomains(defaultLineFareIncomeModel.QueryAll())
}

func ChangeLineFareIncome(lineFareIncomeModel *LineFareIncomeModel) error {
	lineFareIncome, err := lineFareIncomeModel.toDomain()
	err = lineFareIncome.Update()
	return err
}

func fillFromLineFareIncomeDomains(lineFareIncomes []LineFareIncome) (lineFareIncomeModels []LineFareIncomeModel) {
	for _, lineFareIncome := range lineFareIncomes {
		lineFareIncomeModel := new(LineFareIncomeModel)
		lineFareIncomeModel.fillFromDomain(lineFareIncome)
		lineFareIncomeModels = append(lineFareIncomeModels, *lineFareIncomeModel)
	}
	return
}

func (lineFareIncomeModel LineFareIncomeModel) toDomain() (lineFareIncome *LineFareIncome, err error) {
	lineFareIncome = new(LineFareIncome)
	lineFareIncome.ID = lineFareIncomeModel.ID
	lineFareIncome.LineNo = lineFareIncomeModel.LineNo
	lineFareIncome.CarryingAmount = lineFareIncomeModel.CarryingAmount
	lineFareIncome.ActualAmount = lineFareIncomeModel.ActualAmount
	lineFareIncome.WornCoinAmount = lineFareIncomeModel.WornCoinAmount
	lineFareIncome.BusNumbers = lineFareIncomeModel.BusNumbers
	lineFareIncome.CountingDate, err = time.Parse(DateString, lineFareIncomeModel.CountingDate)
	lineFareIncome.CountingStaff1ID = lineFareIncomeModel.CountingStaff1.ID
	lineFareIncome.CountingStaff2ID = lineFareIncomeModel.CountingStaff2.ID
	lineFareIncome.RecordStaffID = lineFareIncomeModel.RecordStaff.ID

	if serviceErr, ok := hasError(err); ok {
		return nil, serviceErr
	}
	return lineFareIncome, nil
}

func (lineFareIncomeModel *LineFareIncomeModel) fillFromDomain(lineFareIncome LineFareIncome) {
	lineFareIncomeModel.ID = lineFareIncome.ID
	lineFareIncomeModel.LineNo = lineFareIncome.LineNo
	lineFareIncomeModel.CarryingAmount = lineFareIncome.CarryingAmount
	lineFareIncomeModel.ActualAmount = lineFareIncome.ActualAmount
	lineFareIncomeModel.WornCoinAmount = lineFareIncome.WornCoinAmount
	lineFareIncomeModel.BusNumbers = lineFareIncome.BusNumbers
	lineFareIncomeModel.CountingDate = lineFareIncome.CountingDate.Format(DateString)
	lineFareIncomeModel.CountingStaff1.fillFromDomain(lineFareIncome.CountingStaff1)
	lineFareIncomeModel.CountingStaff2.fillFromDomain(lineFareIncome.CountingStaff2)
	lineFareIncomeModel.RecordStaff.fillFromDomain(lineFareIncome.RecordStaff)
}

