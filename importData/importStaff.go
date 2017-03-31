package importData

import (
	"database/sql"
	"github.com/tealeg/xlsx"
	"fmt"
	_"github.com/go-sql-driver/mysql"
	"time"
	"strings"
	. "github.com/galahade/bus_staff_managment/domain"
	. "github.com/galahade/bus_staff_managment/util"
)

var timeFormat = "2006.01.02"

func connectToDB() *sql.DB {
	db, err := sql.Open(DriverName, DSN)
	CheckErr(err)
	return db
}

func ImportStaff() {
	excelFileName := "/Users/young_mac/Downloads/入职表.xlsx"

	xlFile, err := xlsx.OpenFile(excelFileName)
	CheckErr(err)

	fmt.Printf("xlFile sheets number is %d.\n", len(xlFile.Sheets))

	sheet := xlFile.Sheets[0]
	var staffs []Staff

	for i, row := range sheet.Rows {
		if (i > 1 ) {
			var staff = new(Staff)
			for j, cell := range row.Cells {
				if (j > 0 && j < 11) {
					text, _ := cell.String()
					if text != "" {
						switch j {
						case 1 :
							staff.Name = strings.TrimSpace(text)
						case 2 :
							staff.PersonalID = strings.TrimSpace(text)
						case 3 :
							staff.Phone = append(staff.Phone, strings.TrimSpace(text))
						case 4 :
							staff.EmergencyContact = strings.TrimSpace(text)
						case 5 :
							staff.EmergencyContactRelation = strings.TrimSpace(text)
						case 6 :
							staff.EmergencyContactPhone = append(staff.EmergencyContactPhone, strings.TrimSpace(text))
						case 7 :
							onboardTime, err := time.Parse(timeFormat, text)
							if err != nil {
								onboardTime = time.Time{}
							}
							staff.OnboardTime = onboardTime
							staff.FirstOnboardTime = onboardTime
						case 8:
							driverType := strings.TrimSpace(text)
							if strings.Contains(driverType, "a3") {
								staff.IsInternship = true
								driverType = strings.ToUpper(strings.Replace(driverType, "a3", "", 0))
							}

							staff.DriverType = driverType
						case 9:
							staff.JobType = jobTypeConv(strings.TrimSpace(text))
						case 10:
							staff.Department = strings.TrimSpace(text)
						case 11:
							staff.StaffIdentity = strings.TrimSpace(text)
						}
					}
				}
			}
			staffs = append(staffs, *staff)
		}
	}
	dbP := connectToDB()

	for _, staff := range staffs {
		fmt.Println(staff)
		if dbP != nil {

			staff.Insert()
		}
	}
}

func ImportDriver() {
	excelFileName := "/Users/young_mac/Downloads/入职表.xlsx"

	xlFile, err := xlsx.OpenFile(excelFileName)
	CheckErr(err)

	fmt.Printf("xlFile sheets number is %d.\n", len(xlFile.Sheets))

	sheet := xlFile.Sheets[1]
	var staffs []Staff

	for i, row := range sheet.Rows {
		if (i > 1 ) {
			var staff = new(Staff)
			for j, cell := range row.Cells {
				if (j > 0 && j < 12) {
					text, _ := cell.String()
					if text != "" {
						switch j {
						case 1 :
							staff.Name = strings.TrimSpace(text)
						case 2 :
							staff.PersonalID = strings.TrimSpace(text)
						case 3 :
							staff.Phone = append(staff.Phone, strings.TrimSpace(text))
						case 4 :
							staff.EmergencyContact = strings.TrimSpace(text)
						case 5 :
							staff.EmergencyContactRelation = strings.TrimSpace(text)
						case 6 :
							staff.EmergencyContactPhone = append(staff.EmergencyContactPhone, strings.TrimSpace(text))
						case 7 :
							onboardTime, err := time.Parse(timeFormat, text)
							if err != nil {
								onboardTime = time.Time{}
							}
							staff.OnboardTime = onboardTime
							staff.FirstOnboardTime = onboardTime
						case 8:
							driverType := strings.TrimSpace(text)
							if strings.Contains(driverType, "实习") {
								staff.IsInternship = true
								driverType = strings.TrimSpace(strings.ToUpper(strings.Replace(driverType, "实习", "", -1)))
							}

							staff.DriverType = driverType
						case 9:
							staff.JobType = jobTypeConv(strings.TrimSpace(text))
						case 10:
							staff.Department = strings.TrimSpace(text)
						case 11:
							staff.StaffIdentity = strings.TrimSpace(text)
						}
					}
				}
			}
			staffs = append(staffs, *staff)
		}
	}
	dbP := connectToDB()

	for _, staff := range staffs {
		fmt.Println(staff)

		if dbP != nil {

			staff.Insert()
		}

	}
}

func jobTypeConv(text string) byte {
	var val byte
	switch text {
	case "司机":
		val = byte(1)
	case "维修":
		val = byte(2)
	case "技术":
		val = byte(4)
	case "保障":
		val = byte(8)
	case "管理":
		val = byte(128)
	default:
		val = byte(255)
	}
	return val
}