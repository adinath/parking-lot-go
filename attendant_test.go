package parking

import (
	"errors"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Attendant", func() {
	mockObserver := NewOwner()
	parkingLotSelector := NewFirstAvailableParkingLotSelector()
	It("agent should not be able to park vehicle in parking lot", func() {
		parkingLot := NewParkingLot(0, &mockObserver)
		vehicle := NewVehicle()
		agent := NewAgent(&parkingLot, parkingLotSelector)
		err := agent.Park(vehicle)
		Expect(err).To(Equal(errors.New("All parkings are full")))
	})

	It("agent should be able to park vehicle in parking lot", func() {
		parkingLot := NewParkingLot(1, &mockObserver)
		vehicle := NewVehicle()
		agent := NewAgent(&parkingLot, parkingLotSelector)
		err := agent.Park(vehicle)
		Expect(err).To(BeNil())
	})

	It("agent should be able to unPark vehicle in parking lot", func() {
		parkingLot := NewParkingLot(1, &mockObserver)
		vehicle := NewVehicle()
		parkingLot.Park(vehicle)
		agent := NewAgent(&parkingLot, parkingLotSelector)
		err := agent.UnPark(vehicle)
		Expect(err).To(BeNil())
	})

	It("agent should not be able to unPark vehicle if it is not parked", func() {
		parkingLot := NewParkingLot(1, &mockObserver)
		vehicle := NewVehicle()
		agent := NewAgent(&parkingLot, parkingLotSelector)
		err := agent.UnPark(vehicle)
		Expect(err).To(Equal(errors.New("Vehicle is not parked")))
	})

})
