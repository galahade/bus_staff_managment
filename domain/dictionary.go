package domain

import (
	"github.com/jinzhu/gorm"
	"github.com/pborman/uuid"
)

var dictionary = new(Dictionary)

type Dictionary struct {
	Domain
	Name                     string
	Type                     int
	IsActive                 bool
}

type JobType Dictionary

type Department Dictionary

type DriverType Dictionary

func (JobType) TableName() string {
	return "dictionary"
}

func (Department) TableName() string {
	return "dictionary"
}

func (DriverType) TableName() string {
	return "dictionary"
}


func (*Dictionary) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.New())
	return nil
}

func (jobType *JobType) Create() error {
	return insertDomain(jobType)
}

func (department *Department) Create() error {
	return insertDomain(department)
}

func (driverType *DriverType) Create() error {
	return insertDomain(driverType)
}

func (jobType *JobType) Update() error {
	UpdateCreateDate(jobType)
	tempDB := Gdb.Save(jobType)
	if tempDB.Error != nil {
		return tempDB.Error
	}
	return nil
}

func (department *Department) Update() error {
	UpdateCreateDate(department)
	tempDB := Gdb.Save(department)
	if tempDB.Error != nil {
		return tempDB.Error
	}
	return nil
}

func (driverType *DriverType) Update() error {
	UpdateCreateDate(driverType)
	tempDB := Gdb.Save(driverType)
	if tempDB.Error != nil {
		return tempDB.Error
	}
	return nil
}


func (JobType) QueryAll() []JobType {
	query := make(map[string]interface{})
	query["type"] = 1
	dictionaries := dictionary.Query(query)
	jobTypes := make([]JobType, len(dictionaries))
	for i, dic := range dictionaries {
		jobTypes[i] = JobType(dic)
	}
	return jobTypes
}

func (Department) QueryAll() []Department {
	query := make(map[string]interface{})
	query["type"] = 2
	dictionaries := dictionary.Query(query)
	departments := make([]Department, len(dictionaries))
	for i, dic := range dictionaries {
		departments[i] = Department(dic)
	}
	return departments
}

func (DriverType) QueryAll() []DriverType {
	query := make(map[string]interface{})
	query["type"] = 3
	dictionaries := dictionary.Query(query)
	driverTypes := make([]DriverType, len(dictionaries))
	for i, dic := range dictionaries {
		driverTypes[i] = DriverType(dic)
	}
	return driverTypes
}

func (Dictionary) Query(query map[string]interface{}) []Dictionary {
	query["is_active"] = 1
	dictionaries := []Dictionary{}
	Gdb.Order("name").Where(query).Find(&dictionaries);
	return dictionaries
}

func (dic *Dictionary) QueryUnique() error {
	Gdb.Where(dic).First(dic)
	if err := checkQueryFirstNotNil(dic); err != nil {
		return err
	}
	return nil
}


