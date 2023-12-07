package parking

type ParkingLot struct {
	capacity       int
	parkedVehicles int
}

func NewParkingLot(cap int) ParkingLot {
	return ParkingLot{
		capacity:       cap,
		parkedVehicles: 0,
	}
}

func (parkingLot *ParkingLot) IsLotAvailable() bool {
	return parkingLot.capacity > parkingLot.parkedVehicles
}

func (parkingLot *ParkingLot) Park(vehicle Vehicle) {
	parkingLot.parkedVehicles = parkingLot.parkedVehicles + 1
}
