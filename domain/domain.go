package domain

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	."github.com/galahade/bus_staff_managment/util"
	"github.com/jinzhu/gorm"
	"time"
	"errors"
)

var RecordAlreadyExistError error  = errors.New("domain: create new record failed, this record already exist.")
var RecordNotFoundError error  = errors.New("domain: record not found in DB.")


var (
	db *sql.DB
	err error
	Gdb *gorm.DB
)

type Domain struct {
	ID        string     `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func init()  {
	db, err = sql.Open(DriverName, DSN)
	CheckErr(err)
	//gorm
	Gdb, err = gorm.Open(DriverName, DSN)
	Gdb.SingularTable(true)
}


func checkChangeDBFailed(result sql.Result, err error, errMessage string)  {
	CheckErr(err)
	if affected, _ := result.RowsAffected(); affected == 0 {
		panic(errMessage)
	}
}

func checkQueryFirstNotNil(domain interface{}) (err error) {
	if Gdb.NewRecord(domain) {
		return RecordNotFoundError
	}
	return nil
}