package parking

import (
	"errors"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ParkingLot", func() {
	owner := NewOwner()
	It("parking lot should not be available to park in zero capacity parking lot", func() {
		parkingLot := NewParkingLot(0, &owner)
		Expect(parkingLot.IsLotAvailable()).To(Equal(false))
	})

	It("parking lot should be available to park with single capacity parking lot", func() {
		parkingLot := NewParkingLot(1, &owner)
		Expect(parkingLot.IsLotAvailable()).To(Equal(true))
	})

	It("parking lot should not be available to park after parking a vehicle with single capacity parking lot", func() {
		parkingLot := NewParkingLot(1, &owner)
		vehicle := NewVehicle()
		parkingLot.Park(vehicle)
		Expect(parkingLot.IsLotAvailable()).To(Equal(false))
	})

	It("should not be able to park once parking lot is full", func() {
		parkingLot := NewParkingLot(1, &owner)
		vehicle := NewVehicle()
		vehicle2 := NewVehicle()
		parkingLot.Park(vehicle)
		err := parkingLot.Park(vehicle2)
		Expect(err).To(Equal(errors.New("Parking lot is full")))
	})

	It("already parked vehicle should not be able to park again", func() {
		parkingLot := NewParkingLot(2, &owner)
		vehicle := NewVehicle()
		parkingLot.Park(vehicle)
		err := parkingLot.Park(vehicle)
		Expect(err).To(Equal(errors.New("Vehicle is already parked")))
	})

	It("can park more than 1 vehicle", func() {
		parkingLot := NewParkingLot(3, &owner)
		vehicle := NewVehicle()
		vehicle2 := NewVehicle()
		parkingLot.Park(vehicle)
		parkingLot.Park(vehicle2)
		Expect(parkingLot.IsLotAvailable()).To(Equal(true))
	})

	It("can check if vehicle is parked in parking lot", func() {
		parkingLot := NewParkingLot(3, &owner)
		vehicle := NewVehicle()
		vehicle2 := NewVehicle()
		parkingLot.Park(vehicle)
		Expect(parkingLot.IsParked(vehicle)).To(Equal(true))
		Expect(parkingLot.IsParked(vehicle2)).To(Equal(false))
	})

	It("should not be able to unpark the vehicle which is not parked", func() {
		parkingLot := NewParkingLot(1, &owner)
		notParkedVehicle := NewVehicle()
		err := parkingLot.UnPark(notParkedVehicle)
		Expect(err).To(Equal(errors.New("Vehicle is not parked")))
	})

	It("should be able to unpark the parked vehicle", func() {
		parkingLot := NewParkingLot(1, &owner)
		vehicle := NewVehicle()
		parkingLot.Park(vehicle)
		err := parkingLot.UnPark(vehicle)
		Expect(err).To(BeNil())
	})

	It("should notify observers when parking lot becomes full", func() {
		owner := NewOwner()
		parkingLot := NewParkingLot(1, &owner)
		vehicle := NewVehicle()
		err := parkingLot.Park(vehicle)
		Expect(err).To(BeNil())
		Expect(owner.IsParkingFull()).To(Equal(true))
	})

	It("should notify observers when parking becomes available", func() {
		owner := NewOwner()
		parkingLot := NewParkingLot(1, &owner)
		vehicle := NewVehicle()
		parkingLot.Park(vehicle)
		Expect(owner.IsParkingFull()).To(Equal(true))
		parkingLot.UnPark(vehicle)
		Expect(owner.IsParkingFull()).To(Equal(false))
	})

	It("should be able to register for parking lot notifications", func() {
		owner := NewOwner()
		owner2 := NewOwner()
		parkingLot := NewParkingLot(1, &owner)
		parkingLot.register(&owner2)
		vehicle := NewVehicle()
		parkingLot.Park(vehicle)
		Expect(owner2.IsParkingFull()).To(Equal(true))
		parkingLot.UnPark(vehicle)
		Expect(owner2.IsParkingFull()).To(Equal(false))
	})

})
