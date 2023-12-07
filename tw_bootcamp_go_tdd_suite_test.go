package parking_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTwBootcampGoTdd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t,
		"ParkingLot Test Suite")
}
