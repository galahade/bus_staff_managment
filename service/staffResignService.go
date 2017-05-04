package service

import (
	"time"
	. "github.com/galahade/bus_staff_managment/domain"
)

type StaffResignModel struct {
	ID           string     `json:"id"`
	ResignDate   string     `json:"resignDate"`
	ResignReason string     `json:"resignReason"`
	Staff        StaffModel `json:"staff"`
}

func Resign(model *StaffResignModel) error {
	staffResign, err := model.toDomain()
	staff,err := model.Staff.toDomain()
	err = staff.UpdateToResign()
	err = staffResign.Create()
	model.ID = staffResign.ID
	model.Staff.IsResign = true
	return err
}

func (staffResignModel StaffResignModel) toDomain() (staffResign *StaffResign, err error) {
	staffResign = new(StaffResign)
	staffResign.ID = staffResignModel.ID
	staffResign.ResignDate, err = time.Parse(DateString, staffResignModel.ResignDate)
	staffResign.ResignReason = staffResignModel.ResignReason
	staffResign.StaffID = staffResignModel.Staff.ID
	if serviceErr, ok := hasError(err); ok {
		return nil, serviceErr
	}
	return staffResign, nil
}

func (staffResignModel *StaffResignModel) fillFromDomain(staffResign StaffResign) {
	staffResignModel.ID = staffResign.ID
	staffResignModel.ResignDate = staffResign.ResignDate.Format(DateString)
	staffResignModel.ResignReason = staffResign.ResignReason
	staffResignModel.Staff.fillFromDomain(staffResign.Staff)
}
