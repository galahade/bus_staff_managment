package service

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/galahade/bus_staff_managment/domain"
)

func TestResign(t *testing.T) {
	domain.Gdb.Debug()
	domain.Gdb.LogMode(true)
	staffResignModel := StaffResignModel{
		ResignDate: "2017-04-27",
		ResignReason: "test test",
		Staff: StaffModel{
			ID: "1e8ff860-1dd2-11b2-9660-60f81dacbf60",
		},
	}
	err := Resign(&staffResignModel)
	assert.Empty(t, err)
}