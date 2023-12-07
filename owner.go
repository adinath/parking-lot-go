package parking

import "github.com/google/uuid"

type ParkingLotObserver interface {
	NotifyParkingFull()
	NotifyParkingAvailable()
}

type Owner struct {
	id          uuid.UUID
	parkingFull bool
}

func (o *Owner) IsParkingFull() bool {
	return o.parkingFull
}

func (o *Owner) NotifyParkingFull() {
	o.parkingFull = true
}

func (o *Owner) NotifyParkingAvailable() {
	o.parkingFull = false
}

func NewOwner() Owner {
	return Owner{
		id:          uuid.New(),
		parkingFull: false,
	}
}
