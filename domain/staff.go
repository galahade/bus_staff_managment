package domain

import (
	"time"
	"github.com/pborman/uuid"
	"github.com/galahade/bus_staff_managment/util"

	"database/sql/driver"
	"strings"
	"errors"
	"database/sql"
	"fmt"
)

type PhoneNumbers []string
type JobType byte

type Staff struct {
	Domain
	Name                     string
	StaffIdentity            string       `gorm:"column:staff_id"`
	JobType                  JobType
	OnboardTime              time.Time
	PersonalID               string
	DriverType               string
	IsInternship             bool
	IsMultiTimeHired         bool
	FirstOnboardTime         time.Time
	Phone                    PhoneNumbers
	Department               string
	EmergencyContact         string
	EmergencyContactPhone    PhoneNumbers
	EmergencyContactRelation string
}

func (staff *Staff) GetJobTypeName() (jobType string) {
	temp := byte(staff.JobType)
	switch temp {
	case byte(1):
		jobType = "司机"
	case byte(2):
		jobType = "维修"
	case byte(4):
		jobType = "技术"
	case byte(8):
		jobType = "保障"
	case byte(128):
		jobType = "管理"
	default:
		jobType = "未知"
	}
	return
}

func (staff *Staff) GetPhoneString() (sPhone string) {
	for _, phone := range staff.Phone {
		if sPhone == "" {
			sPhone = phone
		} else {
			sPhone = sPhone + "," + phone
		}
	}
	return
}

func (staff *Staff) GetECPhoneString() (sPhone string) {
	for _, phone := range staff.EmergencyContactPhone {
		if sPhone == "" {
			sPhone = phone
		} else {
			sPhone = sPhone + "," + phone
		}
	}
	return
}


func (staff Staff) InsertString() string {
	return "INSERT staff SET ID=?, NAME=?, STAFF_ID=?, JOB_TYPE=?, ONBOARD_TIME=?, PERSONAL_ID=?, DRIVER_TYPE=?, " +
		"IS_INTERNSHIP=?, IS_MULTITIME_HIRED=?, FIRST_ONBOARD_TIME=?, PHONE=?, DEPARTMENT=?, EMERGENCY_CONTACT=?, " +
		"EMERGENCY_CONTACT_PHONE=?, EMERGENCY_CONTACT_RELATION=?"
}
func (staff Staff) QueryAllString() string {
	return "SELECT ID, NAME, STAFF_ID, JOB_TYPE, ONBOARD_TIME, PERSONAL_ID, DRIVER_TYPE, " +
		"IS_INTERNSHIP, IS_MULTITIME_HIRED, FIRST_ONBOARD_TIME, PHONE, DEPARTMENT, EMERGENCY_CONTACT, " +
		"EMERGENCY_CONTACT_PHONE, EMERGENCY_CONTACT_RELATION FROM staff"
}
func (staff Staff) QueryByIDString() string {
	return "SELECT ID, NAME, STAFF_ID, JOB_TYPE, ONBOARD_TIME, PERSONAL_ID, DRIVER_TYPE, " +
		"IS_INTERNSHIP, IS_MULTITIME_HIRED, FIRST_ONBOARD_TIME, PHONE, DEPARTMENT, EMERGENCY_CONTACT, " +
		"EMERGENCY_CONTACT_PHONE, EMERGENCY_CONTACT_RELATION FROM staff WHERE ID = ?"
}
func (staff Staff) QueryByStaffIDString() string {
	return "SELECT ID, NAME, STAFF_ID, JOB_TYPE, ONBOARD_TIME, PERSONAL_ID, DRIVER_TYPE, " +
		"IS_INTERNSHIP, IS_MULTITIME_HIRED, FIRST_ONBOARD_TIME, PHONE, DEPARTMENT, EMERGENCY_CONTACT, " +
		"EMERGENCY_CONTACT_PHONE, EMERGENCY_CONTACT_RELATION FROM staff WHERE STAFF_ID = ?"
}
func (staff Staff)  QueryByJobTypeString() string {
	return "SELECT ID, NAME, STAFF_ID, JOB_TYPE, ONBOARD_TIME, PERSONAL_ID, DRIVER_TYPE, " +
		"IS_INTERNSHIP, IS_MULTITIME_HIRED, FIRST_ONBOARD_TIME, PHONE, DEPARTMENT, EMERGENCY_CONTACT, " +
		"EMERGENCY_CONTACT_PHONE, EMERGENCY_CONTACT_RELATION FROM staff WHERE JOB_TYPE = ?"
}

func (staff Staff) DeleteByIdString() string {
	return "DELETE FROM staff WHERE ID = ?"
}
// map go type to db type
func (phoneNumbers PhoneNumbers) Value() (driver.Value, error) {
	var phones string
	for _, phone := range phoneNumbers {
		if phones == "" {
			phones = phone
		} else {
			phones = phones + "," + phone
		}
	}
	return phones, nil
}
func (jobType JobType) Value() (driver.Value, error) {
	var jobTypeDB []byte
	jobTypeDB = append(jobTypeDB, byte(jobType))
	return jobTypeDB, nil
}

func (phoneNumbers *PhoneNumbers) Scan(value interface{}) error {
	if value == nil {
		*phoneNumbers = make([]string, 0)
		return nil
	}
	if phones, err := driver.String.ConvertValue(value); err == nil {
		if v, ok := phones.(string); ok {
			results := strings.Split(v, ",")
			*phoneNumbers = results
		} else if v, ok := phones.([]byte); ok {
			temp := string(v)
			results := strings.Split(temp, ",")
			*phoneNumbers = results
		}
		return nil
	}
	return errors.New("failed to scan PhoneNumbers.")
}
func (jobType *JobType) Scan(value interface{}) error {
	if value == nil {
		*jobType = 0
		return nil
	}
	if v, ok := value.([]byte); ok {
		*jobType = JobType(v[0])
		return nil
	}
	return errors.New("failed to scan JobType.")
}

func (staff *Staff) Insert() {

	stmtP, err := db.Prepare(staff.InsertString())
	util.CheckErr(err)

	res, err := stmtP.Exec(uuid.NewUUID(), staff.Name, staff.StaffIdentity, staff.JobType, staff.OnboardTime,
		staff.PersonalID, staff.DriverType, staff.IsInternship, staff.IsMultiTimeHired, staff.FirstOnboardTime,
		staff.Phone, staff.Department, staff.EmergencyContact, staff.EmergencyContactPhone, staff.EmergencyContactRelation)
	util.CheckErr(err)

	if affected, _ := res.RowsAffected(); affected == 0 {
		fmt.Errorf("Fail to insert staff data into db.\n")
	}
}

func (staff *Staff) QueryByID() {
	stmtP, err := db.Prepare(staff.QueryByIDString())
	util.CheckErr(err)

	err = stmtP.QueryRow(staff.ID).Scan(&staff.ID, &staff.Name, &staff.StaffIdentity, &staff.JobType, &staff.OnboardTime, &staff.PersonalID, &staff.DriverType,
		&staff.IsInternship, &staff.IsMultiTimeHired, &staff.FirstOnboardTime, &staff.Phone, &staff.Department, &staff.EmergencyContact,
		&staff.EmergencyContactPhone, &staff.EmergencyContactRelation)
	util.CheckErr(err)

}

func (staff *Staff) QueryByStaffID() error {
	stmtP, err := db.Prepare(staff.QueryByStaffIDString())
	util.CheckErr(err)

	err = stmtP.QueryRow(staff.StaffIdentity).Scan(&staff.ID, &staff.Name, &staff.StaffIdentity, &staff.JobType, &staff.OnboardTime, &staff.PersonalID, &staff.DriverType,
		&staff.IsInternship, &staff.IsMultiTimeHired, &staff.FirstOnboardTime, &staff.Phone, &staff.Department, &staff.EmergencyContact,
		&staff.EmergencyContactPhone, &staff.EmergencyContactRelation)
	if(err == sql.ErrNoRows) {
		return err
	} else {
		util.CheckErr(err)
	}
	return nil
}

func (staff *Staff) QueryAll() []Staff {
	rows, err := db.Query(staff.QueryAllString())
	util.CheckErr(err)
	return scanQueryResult(rows)

}

func (staff Staff) QueryByJobType(jobType byte) []Staff {
	stmtP, err := db.Prepare(staff.QueryByJobTypeString())
	util.CheckErr(err)

	rows, err := stmtP.Query(jobType)

	return scanQueryResult(rows)
}

func (staff *Staff) DeleteById() {
	stmtP, err := db.Prepare(staff.DeleteByIdString())
	util.CheckErr(err)

	res, err := stmtP.Exec(staff.ID)
	checkChangeDBFailed(res, err, "")
}

func (staff Staff) String() string {
	return fmt.Sprintf("Staff data are : \n id : %s,\n name : %s,\n staffId : %s,\n jobType : %d,\n onboardTime : %s,\n PersonalId : %s,\n " +
		"DriverType : %s,\n IsInternship : %t,\n isMultiTimeHired : %t,\n firstOnboardTime : %s,\n phone : %s,\n department : %s,\n " +
		"emergencyContact : %s,\n emergencyContactPhone : %s,\n emergencyContactRelation : %s,\n",
		staff.ID, staff.Name, staff.StaffIdentity, staff.JobType, staff.OnboardTime, staff.PersonalID, staff.DriverType, staff.IsInternship,
		staff.IsMultiTimeHired, staff.FirstOnboardTime, staff.Phone, staff.Department, staff.EmergencyContact, staff.EmergencyContactPhone,
		staff.EmergencyContactRelation)
}

func (staff Staff) IsQualified() (result bool) {
	result = false
	if staff.IsInternship {
		return
	}
	if strings.Contains(staff.DriverType, "A1") || strings.Contains(staff.DriverType, "A3") {
		result = true
	}
	return
}

func (staff Staff) NeedUpgrade() (result bool) {
	result = false
	if !staff.IsQualified() && !staff.IsInternship && staff.Department == "运营" {
		result = true
	}
	return
}

func scanQueryResult(rows *sql.Rows) []Staff {

	var staffs []Staff
	for rows.Next() {
		staffP := new(Staff)

		err = rows.Scan(&staffP.ID, &staffP.Name, &staffP.StaffIdentity, &staffP.JobType, &staffP.OnboardTime, &staffP.PersonalID,
			&staffP.DriverType, &staffP.IsInternship, &staffP.IsMultiTimeHired, &staffP.FirstOnboardTime,
			&staffP.Phone, &staffP.Department, &staffP.EmergencyContact, &staffP.EmergencyContactPhone, &staffP.EmergencyContactRelation)
		util.CheckErr(err)

		staffs = append(staffs, *staffP)
	}

	return staffs
}