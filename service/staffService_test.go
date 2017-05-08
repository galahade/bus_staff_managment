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

func TestGetAllQualifiedDrivers(t *testing.T) {
	domain.Gdb.Debug()
	domain.Gdb.LogMode(true)

	staffModels := GetAllQualifiedDrivers()

	for _, staffModel := range staffModels {

		fmt.Printf("%#v \n", staffModel)

	}
}

func TestGetAllInternshipDrivers(t *testing.T)  {
	domain.Gdb.Debug()
	domain.Gdb.LogMode(true)
	staffModels := GetAllInternshipDrivers()
	for _, staffModel := range staffModels {

		fmt.Printf("%#v \n", staffModel)

	}
}

func TestGetAllDrivers(t *testing.T)  {
	domain.Gdb.Debug()
	domain.Gdb.LogMode(true)
	staffModels := GetAllDrivers()
	for _, staffModel := range staffModels {

		fmt.Printf("%#v \n", staffModel)

	}
}