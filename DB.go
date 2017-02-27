package main

import (
	"database/sql"
	_"github.com/pborman/uuid"
	_"github.com/go-sql-driver/mysql"
	"time"
	_"fmt"
	"fmt"
	"github.com/pborman/uuid"
)


func dfd() {

	db, err := sql.Open("mysql", "admin:1234@tcp(192.168.1.7:3306)/bus_system?parseTime=true")
	defer db.Close();

	checkErr(err)

	insertData(db)

	queryData(db)

	deleteData(db)



}

func insertData(db *sql.DB)  {
																									stmt, err := db.Prepare("INSERT STAFF SET ID=?, NAME=?, JOB_TYPE=?, ONBOARD_TIME=?, PERSONAL_ID=?, DRIVER_TYPE=?, IS_INTERNSHIP=?, IS_MULTITIME_HIRED=?, FIRST_ONBOARD_TIME=?")
	checkErr(err)

	uuid := uuid.NewUUID()
	name := "testname"
	jobType := byte(1)
	onboardTime := time.Now()
	personalId := "132930198109194713"
	driverType := "A3"
	isInternship := false
	isMultitimeHired := false
	firstOnboardTime := time.Now();

	res, err := stmt.Exec(uuid, name, jobType, onboardTime, personalId, driverType, isInternship, isMultitimeHired, firstOnboardTime)
	checkErr(err)

	_, err = res.RowsAffected();
	checkErr(err)

	//fmt.Printf("affect rows are : %d\n",affect)
}

func queryData(db *sql.DB) {
	rows, err := db.Query("SELECT ID, NAME, JOB_TYPE, ONBOARD_TIME, PERSONAL_ID, DRIVER_TYPE, IS_INTERNSHIP, IS_MULTITIME_HIRED, FIRST_ONBOARD_TIME FROM STAFF")
	checkErr(err)

	for rows.Next()  {
		var id string
		var name string
		var jobType []byte
		var onboardTime time.Time
		var personalId string
		var driverType string
		var isInternship bool
		var isMultitimeHired bool
		var firstOnboardTime time.Time

		err = rows.Scan(&id, &name, &jobType, &onboardTime, &personalId, &driverType, &isInternship, &isMultitimeHired, &firstOnboardTime)
		checkErr(err)

		fmt.Println(id)
		fmt.Println(name)
		fmt.Println(getJobTitle(jobType))
		fmt.Println(onboardTime.String())
		fmt.Println(personalId)
		fmt.Println(driverType)
		fmt.Println(isInternship)
		fmt.Println(isMultitimeHired)
		fmt.Println(firstOnboardTime.String())
	}
}

func getJobTitle(jobType []byte) string {
	var jobName string

	switch jobType[0] {
	case byte(1):
		jobName = "Driver"
	case byte(3):
		jobName = "Maintenance Operator"
	case byte(5):
		jobName = "Technician"
	case byte(7):
		jobName = "Supporter"
	case byte(255):
		jobName = "Manager"
	}

	return jobName
}

func deleteData(db *sql.DB)  {
	stmt, err := db.Prepare("delete from staff")
	checkErr(err)

	res, err := stmt.Exec()
	checkErr(err)

	_, err = res.RowsAffected()
	checkErr(err)

	//fmt.Printf(affect)
}


func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}