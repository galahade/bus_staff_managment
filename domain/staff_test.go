package domain

import "testing"

func TestStaff_IsQualified(t *testing.T) {
	staff := Staff{}
	staffs := staff.QueryAll()

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
	staffs := staff.QueryAll()

	count := 0

	for _, s := range staffs {
		if s.NeedUpgrade() {
			t.Logf("Driver %s need update.\n", s.Name)
			count ++
		}
	}
	t.Logf("There are %d drivers need update.", count)
}

func TestStaff_QueryByJobType(t *testing.T) {
	staff := Staff{}
	staffs := staff.QueryByJobType(1)
	count := 0

	for _, s := range staffs {
		t.Logf("Driver name is %s .\n", s.Name)
		count ++
	}
	t.Logf("There are %d drivers.", count)
}
