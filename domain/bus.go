package domain

import (
	"time"
	"github.com/pborman/uuid"
	"github.com/jinzhu/gorm"
)

type Bus struct {
	Domain
	BusLicense      string
	CustomId        string
	RegisterDate    time.Time
	VehicleIDNumber string      `gorm:"column:VIN"`
	EngineNo        string
	PersonsCapacity int
	BrandID         string
	BusBrand        BusBrand      `gorm:"save_associations:false"`
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

	return insertDomain(busBrand)

}

func (bus *Bus) Create() error {

	return insertDomain(bus)

}

func (bus *Bus) Update() error {
	tempDB := Gdb.Save(bus)
	if tempDB.Error != nil {
		return tempDB.Error
	}
	return nil

}

func (bus *Bus) QueryByLicense() error {
	//gdb.Where("bus_license = ?", bus.BusLicense).First(bus).Related(&bus.Brand)
	//gdb.Preload("Brand").First(bus, "bus_license = ?", bus.BusLicense)//.Model(bus).Related(&bus.Brand)
	Gdb.First(bus, "bus_license = ?", bus.BusLicense).Model(bus).Related(&bus.BusBrand, "BrandID")
	if err := checkQueryFirstNotNil(bus); err != nil {
		return err
	}
	return nil
}

func (bus Bus) QueryAll() []Bus {
	buses := []Bus{}
	Gdb.Find(&buses);
	for i, _ := range buses {
		Gdb.Model(buses[i]).Related(&buses[i].BusBrand, "BrandID")
	}
	return buses
}

func (busBrand BusBrand) QueryAll() []BusBrand {
	busBrands := []BusBrand{}
	Gdb.Find(&busBrands);
	return busBrands
}


