package domain

import (
	"testing"
	"github.com/pborman/uuid"
	"time"
)

func TestBus_Insert(t *testing.T) {
	bus := Bus{
		Id:uuid.NewUUID().String(),
		BusLicense:"冀J233D8",
		CustomId:"33D8",
		Brand:"长安",
		ProductionDate:time.Now()}
	bus.Insert()

}

func TestBus_QueryById(t *testing.T) {
/*
	bus := Bus{
		Id:"5bb41446-f8da-11e6-82ec-60f81dacbf60"}

	bus.QueryById()
	bus.DeleteById()
*/
}

