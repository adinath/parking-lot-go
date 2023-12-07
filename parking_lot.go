package parking

import "errors"

type ParkingLot struct {
	capacity       int
	parkedVehicles []Vehicle
}

func NewParkingLot(cap int) ParkingLot {
	return ParkingLot{
		capacity: cap,
	}
}

func (parkingLot *ParkingLot) IsLotAvailable() bool {
	return parkingLot.capacity > len(parkingLot.parkedVehicles)
}

func (parkingLot *ParkingLot) IsParked(vehicle Vehicle) bool {
	for _, v := range parkingLot.parkedVehicles {
		if v == vehicle {
			return true
		}
	}
	return false
}

func (parkingLot *ParkingLot) Park(vehicle Vehicle) error {
	if parkingLot.IsParked(vehicle) {
		return errors.New("Vehicle is already parked")
	}
	parkingLot.parkedVehicles = append(parkingLot.parkedVehicles, vehicle)
	return nil
}
