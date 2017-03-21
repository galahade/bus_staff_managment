package service

import (
	. "github.com/galahade/bus_staff_managment/domain"
	"time"
	"log"
)

type BusModel struct {
	ID              string       `form:"id" json:"id"`
	License         string       `form:"license" json:"license" binding:"required"`
	CustomID        string       `form:"customID" json:"customID"`
	BrandID         string       `form:"brandID" json:"brandID"`
	BrandAlias      string       `form:"brandAlias" json:"brandAlias"`
	RegisterDate    string       `form:"registerDate" json:"registerDate" binding:"required"`
	VehicleIDNumber string       `form:"vin" json:"vin"`
	EngineNo        string       `form:"engineNo" json:"engineNo"`
	PersonsCapacity int          `form:"personsCapacity" json:"personsCapacity"`
}

type BusBrandModel struct {
	ID    string `form:"id" json:"id"`
	Name  string `form:"name" json:"name"`
	Model string `form:"model" json:"model"`
	Alias string `form:"alias" json:"alias"`
}

var defaultBus = Bus{}
var defaultBusBrand = BusBrand{}

func CreateNewBus(busModel *BusModel) error {
	bus, err := busModel.toDomain()
	log.Printf("bus is : %#v", bus)

	err = bus.Create()
	busModel.ID = bus.BusLicense
	return err
}

func GetAllBuses() ([]BusModel) {
	return fillFromBusDomains(defaultBus.QueryAll())
}

func GetAllBusBrands() ([]BusBrandModel) {
	return fetchFromBusBrandDomains(defaultBusBrand.QueryAll())
}

func fetchFromBusBrandDomains(busBrands []BusBrand) (busBrandModels []BusBrandModel) {
	for _, busBrand := range busBrands {
		busBrandModel := new(BusBrandModel)
		busBrandModel.fillFromDomain(busBrand)
		busBrandModels = append(busBrandModels, *busBrandModel)
	}
	return
}

func fillFromBusDomains(buses []Bus) (busModels []BusModel) {
	for _, bus := range buses {
		busModel := new(BusModel)
		busModel.fillFromDomain(bus)
		busModels = append(busModels, *busModel)
	}
	return
}

func (busBrandModel *BusBrandModel) fillFromDomain(busBrand BusBrand) {
	busBrandModel.ID = busBrand.ID
	busBrandModel.Model = busBrand.Model
	busBrandModel.Name = busBrand.Name
	busBrandModel.Alias = busBrand.Alias
}

func (busModel *BusModel) fillFromDomain(bus Bus) {
	busModel.ID = bus.BusLicense
	busModel.License = bus.BusLicense
	busModel.BrandID = bus.BrandID
	busModel.BrandAlias = bus.BusBrand.Alias
	busModel.CustomID = bus.CustomId
	busModel.RegisterDate = bus.RegisterDate.Format(DateString)
}

func (busModel BusModel) toDomain() (bus *Bus, err error) {
	bus = new(Bus)
	bus.BusLicense = busModel.License
	bus.BrandID = busModel.BrandID
	bus.CustomId = busModel.CustomID
	bus.RegisterDate, err = time.Parse(DateString, busModel.RegisterDate)
	bus.VehicleIDNumber = busModel.VehicleIDNumber
	bus.PersonsCapacity = busModel.PersonsCapacity
	bus.EngineNo = busModel.EngineNo
	if serviceErr, ok := hasError(err); ok {
		return nil, serviceErr
	}
	return bus, nil
}


