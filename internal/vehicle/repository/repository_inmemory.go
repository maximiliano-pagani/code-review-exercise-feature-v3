package repository

import "app/internal/domain"

func NewRepositoryVehicleInMemory(db map[int]*domain.VehicleAttributes) *RepositoryVehicleInMemory {
	return &RepositoryVehicleInMemory{db: db}
}

// RepositoryVehicleInMemory is an struct that represents a vehicle storage in memory.
type RepositoryVehicleInMemory struct {
	// db is the database of vehicles.
	db map[int]*domain.VehicleAttributes
}

// GetAll returns all vehicles
func (s *RepositoryVehicleInMemory) GetAll() (v []*domain.Vehicle, err error) {
	// check if the database is empty
	if len(s.db) == 0 {
		err = ErrRepositoryVehicleNotFound
		return
	}

	// get all vehicles from the database
	v = make([]*domain.Vehicle, 0, len(s.db))
	for key, value := range s.db {
		v = append(v, &domain.Vehicle{
			Id:         key,
			Attributes: *value,
		})
	}

	return
}

func (s *RepositoryVehicleInMemory) GetByBrand(brand string) (v []*domain.Vehicle, err error) {
	// check if the database is empty
	if len(s.db) == 0 {
		err = ErrRepositoryVehicleNotFound
		return
	}

	for key, value := range s.db {
		if value.Brand == brand {
			v = append(v, &domain.Vehicle{
				Id:         key,
				Attributes: *value,
			})
		}
	}

	if len(v) == 0 {
		err = ErrRepositoryVehicleNotFound
		return
	}

	return
}
