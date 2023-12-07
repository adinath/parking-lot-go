package parking

import (
	"errors"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Attendant", func() {
	mockObserver := NewOwner()
	It("agent should not be able to park vehicle in parking lot", func() {
		parkingLot := NewParkingLot(0, &mockObserver)
		vehicle := NewVehicle()
		agent := NewAgent(parkingLot)
		err := agent.Park(vehicle)
		Expect(err).To(Equal(errors.New("All parkings are full")))
	})

	It("agent should be able to park vehicle in parking lot", func() {
		parkingLot := NewParkingLot(1, &mockObserver)
		vehicle := NewVehicle()
		agent := NewAgent(parkingLot)
		err := agent.Park(vehicle)
		Expect(err).To(BeNil())
	})

})
