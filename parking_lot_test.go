package parking

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ParkingLot", func() {
	It("parking lot should not be available to park in zero capacity parking lot", func() {
		parking_lot := NewParkingLot(0)
		Expect(parking_lot.IsLotAvailable()).To(Equal(false))
	})
})
