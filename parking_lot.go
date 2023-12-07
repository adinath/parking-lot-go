package parking

type ParkingLot struct {
	capacity int
}

func NewParkingLot(cap int) ParkingLot {
	return ParkingLot{
		capacity: cap,
	}
}

func (parkingLot *ParkingLot) IsLotAvailable() bool {
	return false
}
