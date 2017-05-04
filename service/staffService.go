package service

import (
	. "github.com/galahade/bus_staff_managment/domain"
	"time"
	"fmt"
	"errors"
)

var defaultStaff = Staff{}

type StaffModel struct {
	ID                       string                `json:"id"`
	Name                     string                `json:"name"`
	JobType                  DictionaryModel       `json:"jobType"`
	OnboardTime              string                `json:"onboardTime"`
	PersonalID               string                `json:"personalID"`
	DriverType               DictionaryModel       `json:"driverType"`
	IsInternship             bool                  `json:"isInternship"`
	IsMultitimeHired         bool                  `json:"isMultitimeHired"`
	IsResign                 bool                  `json:"isResign"`
	FirstOnboardTime         string                `json:"firstOnboardTime"`
	Phone                    string                `json:"phone,omitempty"`
	Department               DictionaryModel       `json:"department"`
	EmergencyContact         string                `json:"emergencyContact"`
	EmergencyContactPhone    string                `json:"emergencyContactPhone,omitempty"`
	EmergencyContactRelation string                `json:"emergencyContactRelation"`
}

type DictionaryModel struct {
	ID       string                `json:"id"`
	Name     string                `json:"name"`
	Type     int                   `json:"type"`
	IsActive bool                  `json:"isActive"`
}

func CreateStaff(staffModel *StaffModel) error {
	staff, err := staffModel.toDomain()
	if (err == nil) {
		err = staff.Create()
		staffModel.ID = staff.ID
	}
	return err
}

func FetchDictionariesByType(dicType string) ([]DictionaryModel, error) {
	switch dicType {
	case "1":
		jobType := JobType{}
		return fillFromJobTypeDomains(jobType.QueryAll()), nil
	case "2":
		department := Department{}
		return fillFromDepartmentDomains(department.QueryAll()), nil
	case "3":
		driverType := DriverType{}
		return fillFromDriverTypeDomains(driverType.QueryAll()), nil
	default:
		return nil, errors.New(fmt.Sprintf("There is not result for the dictType: %s", dicType))

	}
}

func FetchStaffByID(id string) (StaffModel, bool) {
	staff := new(Staff)
	staff.ID = id
	err := staff.QueryUnique()
	if err == nil {
		staffModel := new(StaffModel)
		staffModel.fillFromDomain(*staff)
		return *staffModel, true
	} else {
		return *new(StaffModel), false
	}
}

func GetAllDrivers() []StaffModel {
	query := make(map[string]interface{})
	return fillFromStaffDomains(queryDriver(query))
}

func GetAllStaffs() []StaffModel {
	return fillFromStaffDomains(defaultStaff.Query(nil, false))
}

func GetSupportStaffs() []StaffModel {
	dic1 := Dictionary{
		Name: "后勤",
		Type: 1,
	}
	dic2 := Dictionary{
		Name: "管理",
		Type: 1,
	}
	(&dic1).QueryUnique()
	(&dic2).QueryUnique()
	return fillFromStaffDomains(defaultStaff.QueryByJoin("LEFT JOIN dictionary ON dictionary.id = job_type_id", fmt.Sprintf("dictionary.id = '%s' OR dictionary.id = '%s'", dic1.ID, dic2.ID)))
}

/*this method need correct*/
func GetAllQualifiedDrivers() []StaffModel {
	query := make(map[string]interface{})
	return fillFromStaffDomains(queryDriver(query))
}

func GetAllInternshipDrivers() [] StaffModel {
	query := make(map[string]interface{})
	query["is_internship"] = true
	return fillFromStaffDomains(queryDriver(query))
}

func fillFromStaffDomains(staffs []Staff) (staffModels []StaffModel) {
	for _, staff := range staffs {
		staffModel := new(StaffModel)
		staffModel.fillFromDomain(staff)
		staffModels = append(staffModels, *staffModel)
	}
	return
}

func fillFromJobTypeDomains(dics []JobType) (dicModels []DictionaryModel) {
	for _, dic := range dics {
		dicModel := new(DictionaryModel)
		dicModel.fillFromDomain(Dictionary(dic))
		dicModels = append(dicModels, *dicModel)
	}
	return
}

func fillFromDepartmentDomains(dics []Department) (dicModels []DictionaryModel) {
	for _, dic := range dics {
		dicModel := new(DictionaryModel)
		dicModel.fillFromDomain(Dictionary(dic))
		dicModels = append(dicModels, *dicModel)
	}
	return
}

func fillFromDriverTypeDomains(dics []DriverType) (dicModels []DictionaryModel) {
	for _, dic := range dics {
		dicModel := new(DictionaryModel)
		dicModel.fillFromDomain(Dictionary(dic))
		dicModels = append(dicModels, *dicModel)
	}
	return
}

func (dicModel *DictionaryModel) fillFromDomain(dic Dictionary) {
	dicModel.ID = dic.ID
	dicModel.Name = dic.Name
	dicModel.Type = dic.Type
	dicModel.IsActive = dic.IsActive
}

func (dicModel *DictionaryModel) toDomain() (*Dictionary) {
	dic := new(Dictionary)
	dic.ID = dicModel.ID
	dic.Type = dicModel.Type
	dic.IsActive = dicModel.IsActive
	dic.Name = dicModel.Name
	return dic
}

func (staffModel *StaffModel) fillFromDomain(staff Staff) {
	staffModel.ID = staff.ID
	staffModel.Name = staff.Name
	staffModel.JobType.fillFromDomain(Dictionary(staff.JobType))
	staffModel.OnboardTime = staff.OnboardTime.Format(DateString)
	staffModel.PersonalID = staff.PersonalID
	staffModel.DriverType.fillFromDomain(Dictionary(staff.DriverType))
	staffModel.IsInternship = staff.IsInternship
	staffModel.IsMultitimeHired = staff.IsMultitimeHired
	staffModel.IsResign = staff.IsResign
	staffModel.FirstOnboardTime = staff.FirstOnboardTime.Format(DateString)
	staffModel.Phone = staff.Phone
	staffModel.Department.fillFromDomain(Dictionary(staff.Department))
	staffModel.EmergencyContact = staff.EmergencyContact
	staffModel.EmergencyContactRelation = staff.EmergencyContactRelation
	staffModel.EmergencyContactPhone = staff.EmergencyContactPhone
}

func (staffModel StaffModel) toDomain() (staff *Staff, err error) {
	staff = new(Staff)
	staff.ID = staffModel.ID
	staff.Name = staffModel.Name
	staff.JobTypeID = staffModel.JobType.ID
	staff.OnboardTime, err = time.Parse(DateString, staffModel.OnboardTime)
	staff.PersonalID = staffModel.PersonalID
	staff.DriverTypeID = staffModel.DriverType.ID
	staff.IsInternship = staffModel.IsInternship
	staff.IsResign = staffModel.IsResign
	staff.Phone = staffModel.Phone
	staff.DepartmentID = staffModel.Department.ID
	staff.EmergencyContact = staffModel.EmergencyContact
	staff.EmergencyContactRelation = staffModel.EmergencyContactRelation
	staff.EmergencyContactPhone = staffModel.EmergencyContactPhone
	if (staffModel.FirstOnboardTime == "") {
		staff.FirstOnboardTime = staff.OnboardTime
	}
	if serviceErr, ok := hasError(err); ok {
		return nil, serviceErr
	}
	return staff, nil
}

func queryDriver(query map[string]interface{}) []Staff {
	query["job_type"] = 1
	staffs := defaultStaff.Query(query, false)
	return staffs
}



