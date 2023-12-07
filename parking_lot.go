package parking

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

func (parkingLot *ParkingLot) Park(vehicle Vehicle) {
	parkingLot.parkedVehicles = append(parkingLot.parkedVehicles, vehicle)
}
