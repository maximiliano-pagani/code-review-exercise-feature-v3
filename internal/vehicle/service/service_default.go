package service

import (
	"app/internal/domain"
	"app/internal/vehicle/repository"
	"app/pkg/utils"
	"errors"
	"fmt"
)

// ServiceVehicleDefault is an struct that represents a vehicle service.
type ServiceVehicleDefault struct {
	rp repository.RepositoryVehicle
}

// NewServiceVehicleDefault returns a new instance of a vehicle service.
func NewServiceVehicleDefault(rp repository.RepositoryVehicle) *ServiceVehicleDefault {
	return &ServiceVehicleDefault{rp: rp}
}

// GetAll returns all vehicles.
func (s *ServiceVehicleDefault) GetAll() (v []*domain.Vehicle, err error) {
	v, err = s.rp.GetAll()
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrRepositoryVehicleNotFound):
			err = fmt.Errorf("%w. %v", ErrServiceVehicleNotFound, err)
		default:
			err = fmt.Errorf("%w. %v", ErrServiceVehicleInternal, err)
		}

		return
	}

	return
}

// GetByBrandInTimeRange returns all vehicles of the specified brand found between startYear and endYear interval
func (s *ServiceVehicleDefault) GetByBrandInTimeRange(brand string, startYear int, endYear int) (vehicles []*domain.Vehicle, err error) {
	if !utils.IsValidYearInterval(startYear, endYear) {
		err = ErrServiceVehicleInvalidTimeRange
		return
	}

	brandVehicles, err := s.rp.GetByBrand(brand)

	if err != nil {
		switch {
		case errors.Is(err, repository.ErrRepositoryVehicleNotFound):
			err = fmt.Errorf("%w. %v", ErrServiceVehicleNotFound, err)
		default:
			err = fmt.Errorf("%w. %v", ErrServiceVehicleInternal, err)
		}
		return
	}

	for _, v := range brandVehicles {
		if v.Attributes.Year >= startYear && v.Attributes.Year <= endYear {
			vehicles = append(vehicles, v)
		}
	}

	if len(vehicles) == 0 {
		err = ErrServiceVehicleNotFound
		return
	}

	return
}
