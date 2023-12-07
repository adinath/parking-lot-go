package parking

import "errors"

type Attendant struct {
	parkingLot *ParkingLot
}

func (agent *Attendant) Park(vehicle Vehicle) error {
	if agent.parkingLot.IsLotAvailable() {
		agent.parkingLot.Park(vehicle)
		return nil
	} else {
		return errors.New("All parking is full")
	}
}

func NewAgent(parkingLot ParkingLot) Attendant {
	return Attendant{
		parkingLot: &parkingLot,
	}
}
