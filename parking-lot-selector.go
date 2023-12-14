package parking

import "errors"

type ParkingLotSelector interface {
	Select(parkingLots []ParkingLot) (ParkingLot, error)
}

type FirstAvailableParkingLotSelector struct {
}

func (parkingLotSelector FirstAvailableParkingLotSelector) Select(parkingLots []ParkingLot) (ParkingLot, error) {
	for _, parkingLot := range parkingLots {
		if parkingLot.IsLotAvailable() {
			return parkingLot, nil
		}
	}
	return ParkingLot{}, errors.New("All parkings are full")
}

func NewFirstAvailableParkingLotSelector() FirstAvailableParkingLotSelector {
	return FirstAvailableParkingLotSelector{}
}
