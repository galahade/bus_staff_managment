package domain

import (
	"time"
	"github.com/pborman/uuid"
	"github.com/galahade/bus_staff_managment/util"
	"fmt"
)

type Bus struct {
	Id             string
	BusLicense     string
	CustomId       string
	Brand          string
	ProductionDate time.Time
}

func (bus Bus) InsertString() string {
	return "INSERT INTO BUS (ID, BUS_LICENSE, CUSTOM_ID, BRAND, PRODUCTION_DATE) VALUES(?,?,?,?,?)"
}
func (bus Bus) QueryAllString() string {
	return "SELECT ID, BUS_LICENSE, CUSTOM_ID, BRAND, PRODUCTION_DATE FROM BUS"
}
func (bus Bus) QueryByIdString() string {
	return "SELECT ID, BUS_LICENSE, CUSTOM_ID, BRAND, PRODUCTION_DATE FROM BUS WHERE ID = ?"
}

func (bus Bus) DeleteByIdString() string {
	return "DELETE FROM BUS WHERE ID = ?"
}

func (bus *Bus) Insert() {
	tx, err := db.Begin()
	util.CheckErr(err)
	defer tx.Rollback()

	stmtP, err := tx.Prepare(bus.InsertString())
	util.CheckErr(err)
	defer stmtP.Close()

	res, err := stmtP.Exec(uuid.NewUUID(), bus.BusLicense, bus.CustomId, bus.Brand, bus.ProductionDate)
	checkChangeDBFailed(res, err, "Fail to insert staff data into db.")

	err = tx.Commit()
	util.CheckErr(err)
}

func (bus *Bus) QueryById() {
	fmt.Printf("Bus Id %s will be query.", bus.Id)
	stmtP, err := db.Prepare(bus.QueryByIdString())
	util.CheckErr(err)
	defer stmtP.Close()

	err = stmtP.QueryRow(bus.Id).Scan(&bus.Id, &bus.BusLicense, &bus.CustomId, &bus.Brand, &bus.ProductionDate)
	util.CheckErr(err)
}

func (bus *Bus) DeleteById() {
	fmt.Printf("Bus Id %s will be delete.", bus.Id)
	stmtP, err := db.Prepare(bus.DeleteByIdString())
	util.CheckErr(err)
	defer stmtP.Close()

	res, err := stmtP.Exec(bus.Id)
	checkChangeDBFailed(res, err, "")
}

