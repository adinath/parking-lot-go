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
