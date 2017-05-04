package service

import (
	"testing"
	"github.com/galahade/bus_staff_managment/domain"
	"fmt"
)

func TestGetSupportStaffs(t *testing.T) {
	domain.Gdb.Debug()
	domain.Gdb.LogMode(true)
	staffModels := GetSupportStaffs()
	for _, staffModel := range staffModels {
		fmt.Printf("%#v \n", staffModel)
	}
}
