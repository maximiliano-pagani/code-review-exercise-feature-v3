package repository

import (
	"app/internal/domain"
	"errors"
)

// RepositoryVehicle is the interface that wraps the basic methods for a vehicle repository.
type RepositoryVehicle interface {
	// GetAll returns all vehicles
	GetAll() (v []*domain.Vehicle, err error)
	GetByBrand(brand string) (v []*domain.Vehicle, err error)
}

var (
	// ErrRepositoryVehicleInternal is returned when an internal error occurs.
	ErrRepositoryVehicleInternal = errors.New("repository: internal error")

	// ErrRepositoryVehicleNotFound is returned when a vehicle is not found.
	ErrRepositoryVehicleNotFound = errors.New("repository: vehicle not found")
)
