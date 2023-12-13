package parking

import "errors"

type Attendant struct {
	parkingLots []ParkingLot
}

func (attendant *Attendant) Park(vehicle Vehicle) error {
	availableParkingLot, err := attendant.firstParkingLot()
	if err == nil {
		availableParkingLot.Park(vehicle)
		return nil
	} else {
		return err
	}
}

func (attendant *Attendant) firstParkingLot() (ParkingLot, error) {
	for _, parkingLot := range attendant.parkingLots {
		if parkingLot.IsLotAvailable() {
			return parkingLot, nil
		}
	}
	return ParkingLot{}, errors.New("All parkings are full")
}

func NewAgent(parkingLot ParkingLot) Attendant {
	firstParkingLot := []ParkingLot{parkingLot}
	return Attendant{
		parkingLots: firstParkingLot,
	}
}
