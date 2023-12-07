package parking

import (
	"errors"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ParkingLot", func() {
	It("parking lot should not be available to park in zero capacity parking lot", func() {
		parking_lot := NewParkingLot(0)
		Expect(parking_lot.IsLotAvailable()).To(Equal(false))
	})

	It("parking lot should be available to park with single capacity parking lot", func() {
		parking_lot := NewParkingLot(1)
		Expect(parking_lot.IsLotAvailable()).To(Equal(true))
	})

	It("parking lot should not be available to park after parking a vehicle with single capacity parking lot", func() {
		parking_lot := NewParkingLot(1)
		vehicle := NewVehicle()
		parking_lot.Park(vehicle)
		Expect(parking_lot.IsLotAvailable()).To(Equal(false))
	})

	It("already parked vehicle should not be able to park again", func() {
		parking_lot := NewParkingLot(2)
		vehicle := NewVehicle()
		parking_lot.Park(vehicle)
		err := parking_lot.Park(vehicle)
		Expect(err).To(Equal(errors.New("Vehicle is already parked")))
	})

})
