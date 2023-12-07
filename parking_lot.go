package parking

import "errors"

type ParkingLot struct {
	capacity       int
	parkedVehicles []Vehicle
	observers      []ParkingLotObserver
}

func NewParkingLot(cap int, passed_owner *Owner) ParkingLot {
	first_observers := []ParkingLotObserver{passed_owner}
	return ParkingLot{
		capacity:  cap,
		observers: first_observers,
	}
}

func (parkingLot *ParkingLot) IsLotAvailable() bool {
	return parkingLot.capacity > len(parkingLot.parkedVehicles)
}

func (parkingLot *ParkingLot) IsFull() bool {
	return parkingLot.capacity == len(parkingLot.parkedVehicles)
}
func (parkingLot *ParkingLot) IsParked(vehicle Vehicle) bool {
	for _, parkedVehicle := range parkingLot.parkedVehicles {
		if parkedVehicle == vehicle {
			return true
		}
	}
	return false
}

func (parkingLot *ParkingLot) indexOf(vehicle Vehicle) int {
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
	if len(parkingLot.parkedVehicles) < parkingLot.capacity {
		parkingLot.parkedVehicles = append(parkingLot.parkedVehicles, vehicle)
		if parkingLot.IsFull() {
			parkingLot.notifyParkingFullToAllObservers()
		}
	} else {
		return errors.New("Parking lot of full")
	}
	return nil
}

func (parkingLot *ParkingLot) notifyParkingFullToAllObservers() {
	for _, observer := range parkingLot.observers {
		observer.NotifyParkingFull()
	}
}

func (parkingLot *ParkingLot) notifyParkingAvailableToAllObservers() {
	for _, observer := range parkingLot.observers {
		observer.NotifyParkingAvailable()
	}
}

func (parkingLot *ParkingLot) UnPark(vehicle Vehicle) error {
	if parkingLot.IsParked(vehicle) {
		parked_index := parkingLot.indexOf(vehicle)
		parkingLot.parkedVehicles = append(parkingLot.parkedVehicles[:parked_index], parkingLot.parkedVehicles[parked_index+1:]...)
		parkingLot.notifyParkingAvailableToAllObservers()
		return nil
	} else {
		return errors.New("Vehicle is not parked")
	}
}
