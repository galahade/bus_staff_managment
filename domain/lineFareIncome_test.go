package domain

import (
	"time"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLineFareIncome_Create(t *testing.T) {
	setGORMShowSQL()
	lineFareIncome := LineFareIncome{
		LineNo: 1,
		CarryingAmount: 200,
		BusNumbers: 3,
		CountingDate: time.Now(),
		CountingStaff1ID: "3cd0c63a-f433-11e6-8fa1-60f81dacbf60",
		CountingStaff2ID: "3cd0c63a-f433-11e6-8fa1-60f81dacbf60",
		RecordStaffID: "3cd0c63a-f433-11e6-8fa1-60f81dacbf60",
	}
	(&lineFareIncome).Create()

	assert.NotEmpty(t, lineFareIncome.ID)
	assert.NotEmpty(t, lineFareIncome.CreatedAt)

	Gdb.Unscoped().Delete(&lineFareIncome)
}