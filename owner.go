package parking

import "github.com/google/uuid"

type Owner struct {
	id          uuid.UUID
	parkingFull bool
}

func (o *Owner) IsParkingFull() bool {
	return o.parkingFull
}

func (o *Owner) notifyParkingFull() {
	o.parkingFull = true
}

func (o *Owner) notifyParkingAvailable() {
	o.parkingFull = false
}

func NewOwner() Owner {
	return Owner{
		id:          uuid.New(),
		parkingFull: false,
	}
}
