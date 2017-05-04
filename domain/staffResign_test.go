package domain

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestStaffResign_Create(t *testing.T) {
	setGORMShowSQL()
	staffResign := StaffResign{
		ResignDate: time.Now(),
		StaffID: "3cd0c63a-f433-11e6-8fa1-60f81dacbf60",
		ResignReason: "个人原因辞职。",
	}
	(&staffResign).Create()
	assert.NotEmpty(t, staffResign.ID)
	assert.NotEmpty(t, staffResign.CreatedAt)

	Gdb.Unscoped().Delete(&staffResign)
}
