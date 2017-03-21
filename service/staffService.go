package service

import (
	. "github.com/galahade/bus_staff_managment/domain"
	"log"
)

var defaultStaff = Staff{}

func FetchDriverBySID(sid string) (StaffModel, bool) {
	log.Printf("query for staff ID : %s", sid)
	staff := new(Staff)
	staff.StaffId = sid
	err := staff.QueryByStaffID()
	if err == nil {
		staffModel := new(StaffModel)
		staffModel.fetchFromDomain(*staff)
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
		staffModel.fetchFromDomain(staff)
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
			staffModel.fetchFromDomain(staff)
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
			staffModel.fetchFromDomain(staff)
			staffModels = append(staffModels, *staffModel)
		}
	}
	return staffModels
}

type StaffModel struct {
	Id                       string       `json:"-"`
	Name                     string       `json:"name"`
	StaffId                  string       `json:"id"`
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

func (staffModel *StaffModel) fetchFromDomain(staff Staff) {
	staffModel.Id = staff.ID
	staffModel.Name = staff.Name
	staffModel.StaffId = staff.StaffId
	staffModel.JobType = getJobTypeFromByte(byte(staff.JobType))
	staffModel.OnboardTime = staff.OnboardTime.Format("2006-01-02")
	staffModel.PersonalID = staff.PersonalID
	staffModel.DriverType = staff.DriverType
	staffModel.IsInternship = staff.IsInternship
	staffModel.IsMultiTimeHired = staff.IsMultiTimeHired
	staffModel.FirstOnboardTime = staff.FirstOnboardTime.Format("2006-01-02")
	staffModel.Phone = getPhone(staff.Phone)
	staffModel.Department = staff.Department
	staffModel.EmergencyContact = staff.EmergencyContact
	staffModel.EmergencyContactRelation = staff.EmergencyContactRelation
	staffModel.EmergencyContactPhone = getPhone(staff.EmergencyContactPhone)
}

func getPhone(phones PhoneNumbers) (sPhone string) {
	for _, phone := range phones {
		if sPhone == "" {
			sPhone = phone
		} else {
			sPhone = sPhone + "," + phone
		}
	}
	return
}

func getJobTypeFromByte(b byte) (jobtype string) {
	switch b {
	case byte(1):
		jobtype = "司机"
	case byte(2):
		jobtype = "维修"
	case byte(4):
		jobtype = "技术"
	case byte(8):
		jobtype = "保障"
	case byte(128):
		jobtype = "管理"
	default:
		jobtype = "未知"
	}
	return
}
