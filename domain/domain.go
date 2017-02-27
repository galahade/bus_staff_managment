package domain

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	."github.com/galahade/bus_staff_managment/util"
)


var db *sql.DB
var err error

type Domain interface {
	InsertString() string
	QueryAllString() string
	QueryByIdString() string
	DeleteByIdString() string
}

func init()  {
	db, err = sql.Open(DriverName, DSN)
	CheckErr(err)
}


func checkChangeDBFailed(result sql.Result, err error, errMessage string)  {
	CheckErr(err)
	if affected, _ := result.RowsAffected(); affected == 0 {
		panic(errMessage)
	}
}
