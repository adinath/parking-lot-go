package parking

import "github.com/google/uuid"

type Vehicle struct {
	id uuid.UUID
}

func NewVehicle() Vehicle {
	return Vehicle{
		id: uuid.New(),
	}
}
