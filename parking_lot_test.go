package parking

import (
	"errors"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ParkingLot", func() {
	owner := NewOwner()
	It("parking lot should not be available to park in zero capacity parking lot", func() {
		parking_lot := NewParkingLot(0, &owner)
		Expect(parking_lot.IsLotAvailable()).To(Equal(false))
	})

	It("parking lot should be available to park with single capacity parking lot", func() {
		parking_lot := NewParkingLot(1, &owner)
		Expect(parking_lot.IsLotAvailable()).To(Equal(true))
	})

	It("parking lot should not be available to park after parking a vehicle with single capacity parking lot", func() {
		parking_lot := NewParkingLot(1, &owner)
		vehicle := NewVehicle()
		parking_lot.Park(vehicle)
		Expect(parking_lot.IsLotAvailable()).To(Equal(false))
	})

	It("should not be able to park once parking lot is full", func() {
		parking_lot := NewParkingLot(1, &owner)
		vehicle := NewVehicle()
		vehicle2 := NewVehicle()
		parking_lot.Park(vehicle)
		err := parking_lot.Park(vehicle2)
		Expect(err).To(Equal(errors.New("Parking lot is full")))
	})

	It("already parked vehicle should not be able to park again", func() {
		parking_lot := NewParkingLot(2, &owner)
		vehicle := NewVehicle()
		parking_lot.Park(vehicle)
		err := parking_lot.Park(vehicle)
		Expect(err).To(Equal(errors.New("Vehicle is already parked")))
	})

	It("can park more than 1 vehicle", func() {
		parking_lot := NewParkingLot(3, &owner)
		vehicle := NewVehicle()
		vehicle2 := NewVehicle()
		parking_lot.Park(vehicle)
		parking_lot.Park(vehicle2)
		Expect(parking_lot.IsLotAvailable()).To(Equal(true))
	})

	It("can check if vehicle is parked in parking lot", func() {
		parking_lot := NewParkingLot(3, &owner)
		vehicle := NewVehicle()
		vehicle2 := NewVehicle()
		parking_lot.Park(vehicle)
		Expect(parking_lot.IsParked(vehicle)).To(Equal(true))
		Expect(parking_lot.IsParked(vehicle2)).To(Equal(false))
	})

	It("should not be able to unpark the vehicle which is not parked", func() {
		parking_lot := NewParkingLot(1, &owner)
		not_parked_vehicle := NewVehicle()
		err := parking_lot.UnPark(not_parked_vehicle)
		Expect(err).To(Equal(errors.New("Vehicle is not parked")))
	})

	It("should be able to unpark the parked vehicle", func() {
		parking_lot := NewParkingLot(1, &owner)
		vehicle := NewVehicle()
		parking_lot.Park(vehicle)
		err := parking_lot.UnPark(vehicle)
		Expect(err).To(BeNil())
	})

	It("should notify observers when parking lot becomes full", func() {
		owner := NewOwner()
		parking_lot := NewParkingLot(1, &owner)
		vehicle := NewVehicle()
		err := parking_lot.Park(vehicle)
		Expect(err).To(BeNil())
		Expect(owner.IsParkingFull()).To(Equal(true))
	})

	It("should notify observers when parking becomes available", func() {
		owner := NewOwner()
		parking_lot := NewParkingLot(1, &owner)
		vehicle := NewVehicle()
		parking_lot.Park(vehicle)
		Expect(owner.IsParkingFull()).To(Equal(true))
		parking_lot.UnPark(vehicle)
		Expect(owner.IsParkingFull()).To(Equal(false))
	})

	It("should be able to register for parking lot notifications", func() {
		owner := NewOwner()
		owner2 := NewOwner()
		parking_lot := NewParkingLot(1, &owner)
		parking_lot.register(&owner2)
		vehicle := NewVehicle()
		parking_lot.Park(vehicle)
		Expect(owner2.IsParkingFull()).To(Equal(true))
		parking_lot.UnPark(vehicle)
		Expect(owner2.IsParkingFull()).To(Equal(false))
	})

})
