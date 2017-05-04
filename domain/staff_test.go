package domain

import (
	_ "github.com/stretchr/testify/assert"
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestStaff_Create(t *testing.T) {
	setGORMShowSQL()
	staff := Staff{
		Name: "test",
		JobTypeID: "6624e490-29ca-11e7-8e57-df067f6c9ed3",
		OnboardTime: time.Now(),
		PersonalID: "1234567890",
		DriverTypeID: "04a78b8a-2a30-11e7-8e57-df067f6c9ed3",
		IsInternship: false,
		Phone: "12332124",
		DepartmentID: "6624e774-29ca-11e7-8e57-df067f6c9ed3",
		EmergencyContact: "你好",
		EmergencyContactPhone: "12312321312",
		EmergencyContactRelation: "fuqi",
	}
	(&staff).Create()

	assert.NotEmpty(t, staff.ID)
	assert.NotEmpty(t, staff.CreatedAt)
	Gdb.Unscoped().Delete(&staff)
}

func TestStaff_IsQualified(t *testing.T) {
	staff := Staff{}
	staffs := staff.Query(nil, false)

	count := 0
	for _, s := range staffs {
		if s.IsQualified() {
			t.Logf("Driver %s is qualified.\n", s.Name)
			count++
		}
	}
	t.Logf("There are %d drivers are qualified.\n", count)
}

func TestStaff_NeedUpgrade(t *testing.T) {
	staff := Staff{}
	staffs := staff.Query(nil, false)

	count := 0

	for _, s := range staffs {
		if s.NeedUpgrade() {
			t.Logf("Driver %s need update.\n", s.Name)
			count ++
		}
	}
	t.Logf("There are %d drivers need update.", count)
}


func TestStaff_Query(t *testing.T) {
	setGORMShowSQL()
	staff := Staff{}
	//count := 0

	staffs := staff.Query(nil, false)
	count := 0
	for _, s := range staffs {
		t.Logf("Driver is %s .\n", s)
		count ++
	}
	t.Logf("There are %d drivers.", count)
	/*
	query := make(map[string]interface{})
	query["job_type"] = 1

	staffs := staff.Query(query, true)
	for _, s := range staffs {
		t.Logf("Driver name is %s .\n", s.Name)
		count ++
	}
	t.Logf("There are %d drivers.", count)
	*/
}

func TestStaff_QueryUnique(t *testing.T) {
	setGORMShowSQL()
	//staff := Staff{StaffIdentity: "1001"}

	staff := Staff{Domain: Domain{
		ID: "3ca31c2a-f433-11e6-8fa1-60f81dacbf60",
	} }
	(&staff).QueryUnique()
	t.Logf("Driver is %s .\n", staff)
}

