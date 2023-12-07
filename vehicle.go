package parking

import "github.com/google/uuid"

type Vehicle struct {
	id string
}

func NewVehicle() Vehicle {
	return Vehicle{
		id: uuid.New().String(),
	}
}
