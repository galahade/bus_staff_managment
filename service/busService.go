package service

import (
	. "github.com/galahade/bus_staff_managment/domain"
	"time"
)

type BusModel struct {
	ID             string       `form:"id" json:"id"`
	License        string       `form:"license" json:"license" binding:"required"`
	CustomID       string       `form:"customID" json:"customID"`
	Alias          string       `form:"alias" json:"alias"`
	BrandID        string       `form:"brandID" json:"brandID"`
	RegisterDate   string       `form:"registerDate" json:"registerDate" binding:"required"`
}

var defaultBus = Bus{}

func CreateNewBus(busModel *BusModel) error {
	bus, err := busModel.toDomain()
	err = bus.Create()
	busModel.ID = bus.BusLicense
	return err
}

func GetAllBuses() ([]BusModel) {
	return fetchFromDomain(defaultBus.QueryAll())
}

func fetchFromDomain(buses []Bus) (busModels []BusModel) {
	for _, bus := range buses {
		busModel := new(BusModel)
		busModel.fetchFromDomain(bus)
		busModels = append(busModels, *busModel)
	}
	return
}

func (busModel *BusModel) fetchFromDomain(bus Bus) {
	busModel.ID = bus.BusLicense
	busModel.License = bus.BusLicense
	busModel.Alias = bus.BusBrand
	busModel.CustomID = bus.CustomId
	busModel.RegisterDate = bus.RegisterDate.Format(DateString)
}

func (busModel BusModel) toDomain() (bus *Bus, err error) {
	bus = new(Bus)
	bus.BusLicense = busModel.License
	bus.BusBrand = busModel.Alias
	bus.CustomId = busModel.CustomID
	bus.RegisterDate, err = time.Parse(DateString, busModel.RegisterDate)
	if serviceErr, ok := hasError(err); ok {
		return nil, serviceErr
	}
	return bus, nil
}


