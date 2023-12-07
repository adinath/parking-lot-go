package parking

import "errors"

type Agent struct {
	parkingLot *ParkingLot
}

func (agent *Agent) Park(vehicle Vehicle) error {
	if agent.parkingLot.IsLotAvailable() {
		agent.parkingLot.Park(vehicle)
		return nil
	} else {
		return errors.New("All parking is full")
	}
}

func NewAgent(parkingLot ParkingLot) Agent {
	return Agent{
		parkingLot: &parkingLot,
	}
}
