package parking

import "errors"

type Attendant struct {
	parkingLots        []ParkingLot
	parkingLotSelector ParkingLotSelector
}

func (attendant *Attendant) Park(vehicle Vehicle) error {
	availableParkingLot, err := attendant.parkingLotSelector.Select(attendant.parkingLots)
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

func (attendant *Attendant) UnPark(vehicle Vehicle) error {

	for _, parkingLot := range attendant.parkingLots {
		if parkingLot.IsParked(vehicle) {
			parkingLot.UnPark(vehicle)
			return nil
		}
	}
	return errors.New("Vehicle is not parked")
}

func NewAgent(parkingLot *ParkingLot, selector ParkingLotSelector) Attendant {
	firstParkingLot := []ParkingLot{*parkingLot}
	return Attendant{
		parkingLots:        firstParkingLot,
		parkingLotSelector: selector,
	}
}
