package service

import (
	. "github.com/galahade/bus_staff_managment/domain"
	"log"
)

var defaultStaff = Staff{}

func FetchDriverBySID(sid string) (StaffModel, bool) {
	log.Printf("query for staff ID : %s", sid)
	staff := new(Staff)
	staff.StaffIdentity = sid
	err := staff.QueryByStaffID()
	if err == nil {
		staffModel := new(StaffModel)
		staffModel.fillFromDomain(*staff)
		return  *staffModel, true
	} else {
		return *new(StaffModel), false
	}
}

func GetAllDrivers() []StaffModel {
	staffs := defaultStaff.QueryByJobType(1)
	var staffModels []StaffModel
	for _, staff := range staffs {
		staffModel := new(StaffModel)
		staffModel.fillFromDomain(staff)
		staffModels = append(staffModels, *staffModel)
	}
	return staffModels
}

func GetAllQualifiedDrivers() []StaffModel {
	staffs := defaultStaff.QueryByJobType(1)
	var staffModels []StaffModel
	for _, staff := range staffs {
		if staff.IsQualified() {
			staffModel := new(StaffModel)
			staffModel.fillFromDomain(staff)
			staffModels = append(staffModels, *staffModel)
		}
	}
	return staffModels
}

func GetAllInternshipDrivers() [] StaffModel {
	staffs := defaultStaff.QueryByJobType(1)
	var staffModels []StaffModel
	for _, staff := range staffs {
		if staff.IsInternship {
			staffModel := new(StaffModel)
			staffModel.fillFromDomain(staff)
			staffModels = append(staffModels, *staffModel)
		}
	}
	return staffModels
}

type StaffModel struct {
	ID                       string       `json:"id"`
	Name                     string       `json:"name"`
	StaffID                  string       `json:"sid"`
	JobType                  string       `json:"jobType"`
	OnboardTime              string       `json:"onboardTime"`
	PersonalID               string       `json:"personalID"`
	DriverType               string       `json:"driverType"`
	IsInternship             bool         `json:"isInternship"`
	IsMultiTimeHired         bool         `json:"isMultitimeHired"`
	FirstOnboardTime         string       `json:"firstOnboardTime"`
	Phone                    string       `json:"phone,omitempty"`
	Department               string       `json:"department"`
	EmergencyContact         string       `json:"emergencyContact"`
	EmergencyContactPhone    string       `json:"emergencyContactPhone,omitempty"`
	EmergencyContactRelation string       `json:"emergencyContactRelation"`
}

func (staffModel *StaffModel) fillFromDomain(staff Staff) {
	staffModel.ID = staff.ID
	staffModel.Name = staff.Name
	staffModel.StaffID = staff.StaffIdentity
	staffModel.JobType = staff.GetJobTypeName()
	staffModel.OnboardTime = staff.OnboardTime.Format(DateString)
	staffModel.PersonalID = staff.PersonalID
	staffModel.DriverType = staff.DriverType
	staffModel.IsInternship = staff.IsInternship
	staffModel.IsMultiTimeHired = staff.IsMultiTimeHired
	staffModel.FirstOnboardTime = staff.FirstOnboardTime.Format(DateString)
	staffModel.Phone = staff.GetPhoneString()
	staffModel.Department = staff.Department
	staffModel.EmergencyContact = staff.EmergencyContact
	staffModel.EmergencyContactRelation = staff.EmergencyContactRelation
	staffModel.EmergencyContactPhone = staff.GetECPhoneString()
}


