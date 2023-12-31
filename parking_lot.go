package parking

import "errors"

type ParkingLot struct {
	capacity       int
	parkedVehicles []Vehicle
	observers      []ParkingLotObserver
}

func NewParkingLot(cap int, passedOwner *Owner) ParkingLot {
	firstObservers := []ParkingLotObserver{passedOwner}
	return ParkingLot{
		capacity:  cap,
		observers: firstObservers,
	}
}

func (parkingLot *ParkingLot) IsLotAvailable() bool {
	return parkingLot.capacity > len(parkingLot.parkedVehicles)
}

func (parkingLot *ParkingLot) IsFull() bool {
	return parkingLot.capacity == len(parkingLot.parkedVehicles)
}
func (parkingLot *ParkingLot) IsParked(vehicle Vehicle) bool {
	return parkingLot.indexOf(vehicle) >= 0
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
	if len(parkingLot.parkedVehicles) == parkingLot.capacity {
		return errors.New("Parking lot is full")
	}
	parkingLot.parkedVehicles = append(parkingLot.parkedVehicles, vehicle)
	if parkingLot.IsFull() {
		parkingLot.notifyParkingFullToAllObservers()
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
		parkingLot.removeVehicle(vehicle)
		parkingLot.notifyParkingAvailableToAllObservers()
		return nil
	} else {
		return errors.New("Vehicle is not parked")
	}
}

func (parkingLot *ParkingLot) removeVehicle(vehicle Vehicle) {
	parkedIndex := parkingLot.indexOf(vehicle)
	parkingLot.parkedVehicles = append(parkingLot.parkedVehicles[:parkedIndex], parkingLot.parkedVehicles[parkedIndex+1:]...)
}

func (parkingLot *ParkingLot) register(observer ParkingLotObserver) {
	parkingLot.observers = append(parkingLot.observers, observer)
}
