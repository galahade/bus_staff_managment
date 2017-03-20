package domain

import (
	"time"
	"github.com/pborman/uuid"
	"github.com/jinzhu/gorm"
)

type Bus struct {
	Domain
	BusLicense   string
	CustomId     string
	RegisterDate time.Time
	VIN          string
	EngineNo     string
	BrandID      string
	BusBrand     `gorm:"ForeignKey:BrandID"`
}

type BusBrand struct {
	Domain
	Name  string
	Model string
	Alias string
}

func (*Bus) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.New())
	return nil
}

func (Bus) TableName() string {
	return "bus_basic"
}

func (*BusBrand) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.New())
	return nil
}

func (busBrand *BusBrand) Create() error {
	if !gdb.NewRecord(*busBrand) {
		return RecordAlreadyExistError
	}
	tempDB := gdb.Create(busBrand)
	if tempDB.Error != nil {
		return tempDB.Error
	}
	return nil

}

func (bus *Bus) Create() error {
	if !gdb.NewRecord(*bus) {
		return RecordAlreadyExistError
	}
	tempDB := gdb.Create(bus)
	if tempDB.Error != nil {
		return tempDB.Error
	}
	return nil

}

func (bus *Bus) QueryByLicense() error {
	//gdb.Where("bus_license = ?", bus.BusLicense).First(bus).Related(&bus.Brand)
	//gdb.Preload("Brand").First(bus, "bus_license = ?", bus.BusLicense)//.Model(bus).Related(&bus.Brand)
	gdb.First(bus, "bus_license = ?", bus.BusLicense).Model(bus).Related(&bus.BusBrand, "BrandID")
	if err := checkQueryFirstNotNil(bus); err != nil {
		return err
	}
	return nil
}

func (bus Bus) QueryAll() []Bus {
	buses := []Bus{}
	gdb.Find(&buses);
	for i, _ := range buses {
		gdb.Model(buses[i]).Related(&buses[i].BusBrand, "BrandID")
	}
	return buses
}


