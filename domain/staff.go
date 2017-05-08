package domain

import (
	"time"
	"github.com/pborman/uuid"
	"fmt"
	"github.com/jinzhu/gorm"
)

type Staff struct {
	Domain
	Name                     string
	JobTypeID                string
	JobType                  JobType           `gorm:"ForeignKey:JobTypeID"`
	OnboardTime              time.Time
	PersonalID               string
	DriverTypeID             string
	DriverType               DriverType        `gorm:"ForeignKey:DriverTypeID"`
	IsInternship             bool
	IsMultitimeHired         bool
	IsResign                 bool
	FirstOnboardTime         time.Time
	Phone                    string
	DepartmentID             string
	Department               Department        `gorm:"ForeignKey:DepartmentID"`
	EmergencyContact         string
	EmergencyContactPhone    string
	EmergencyContactRelation string
}

func (*Staff) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.New())
	return nil
}

func (staff *Staff) Create() error {
	return insertDomain(staff)
}

func (staff *Staff) Update() error {
	UpdateCreateDate(staff)
	tempDB := Gdb.Save(staff)
	if tempDB.Error != nil {
		return tempDB.Error
	}
	return nil
}

func (staff *Staff) UpdateToResign() error {
	tempDB := Gdb.Model(&staff).Update("is_resign", true)
	if tempDB.Error != nil {
		return tempDB.Error
	}
	return nil
}

func (staff Staff) QueryByJoin(joins, conditions []string) []Staff {
	staffs := []Staff{}
	db := Gdb.Model(staff)
	for i := range joins {
		db = db.Joins(joins[i])
	}

	for i := range conditions {
		db = db.Where(conditions[i])
	}

	db.Find(&staffs)

	for i := range staffs {
		Gdb.Model(staffs[i]).Related(&staffs[i].JobType).Related(&staffs[i].DriverType).Related(&staffs[i].Department)
	}
	return staffs
}

func (Staff) Query(query map[string]interface{}, isNot bool) []Staff {
	staffs := []Staff{}
	db := Gdb.Where("is_resign = false")
	if (query != nil) {
		if(isNot) {
			db = db.Not(query)
		} else {
			db = db.Where(query)
		}
	}
	db.Find(&staffs).Order("staff_id")
	for i := range staffs {
		Gdb.Model(staffs[i]).Related(&staffs[i].JobType).Related(&staffs[i].DriverType).Related(&staffs[i].Department)
	}
	return staffs
}

func (staff *Staff) QueryUnique() error {
	Gdb.First(staff).Related(&staff.JobType).Related(&staff.DriverType).Related(&staff.Department)
	if err := checkQueryFirstNotNil(staff); err != nil {
		return err
	}
	return nil
}


func (staff Staff) String() string {
	return fmt.Sprintf("Staff data are : \n id : %s,\n name : %s,\n jobType : %#v,\n onboardTime : %s,\n PersonalId : %s,\n " +
		"DriverType : %#v,\n IsInternship : %t,\n isMultiTimeHired : %t,\n firstOnboardTime : %s,\n phone : %s,\n department : %#v,\n " +
		"emergencyContact : %s,\n emergencyContactPhone : %s,\n emergencyContactRelation : %s,\n",
		staff.ID, staff.Name, staff.JobType, staff.OnboardTime, staff.PersonalID,
		staff.DriverType, staff.IsInternship, staff.IsMultitimeHired, staff.FirstOnboardTime, staff.Phone,
		staff.Department, staff.EmergencyContact, staff.EmergencyContactPhone,
		staff.EmergencyContactRelation)
}

func (staff Staff) IsQualified() (result bool) {
	result = false
	if staff.IsInternship {
		return
	}
	if (staff.DriverType.Name == "A1" || staff.DriverType.Name == "A3") {
		result = true
	}
	return
}

func (staff Staff) NeedUpgrade() (result bool) {
	result = false
	if !staff.IsQualified() && !staff.IsInternship && staff.JobType.Name == "司机" {
		result = true
	}
	return
}