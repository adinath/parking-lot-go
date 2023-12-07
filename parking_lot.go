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
	for _, parkedVehicle := range parkingLot.parkedVehicles {
		if parkedVehicle == vehicle {
			return true
		}
	}
	return false
}

func (parkingLot *ParkingLot) parkedIndex(vehicle Vehicle) int {
	for i, parkedVehicle := range parkingLot.parkedVehicles {
		if parkedVehicle == vehicle {
			return i
		}
	}
	return -1
}

func (parkingLot *ParkingLot) Park(vehicle Vehicle) error {
	if parkingLot.IsParked(vehicle) {
		return errors.New("Vehicle is already parked")
	}
	parkingLot.parkedVehicles = append(parkingLot.parkedVehicles, vehicle)
	return nil
}

func (parkingLot *ParkingLot) UnPark(vehicle Vehicle) error {
	if parkingLot.IsParked(vehicle) {
		parked_index := parkingLot.parkedIndex(vehicle)
		parkingLot.parkedVehicles = append(parkingLot.parkedVehicles[:parked_index], parkingLot.parkedVehicles[parked_index+1:]...)
		return nil
	} else {
		return errors.New("Vehicle is not parked")
	}
}
