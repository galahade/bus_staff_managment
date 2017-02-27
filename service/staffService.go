package service

import . "github.com/galahade/bus_staff_managment/domain"

var defaultStaff = Staff{}

func GetAllDrivers() []Staff {
	return defaultStaff.QueryByJobType(1)
}
