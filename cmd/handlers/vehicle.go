package handlers

import (
	"app/internal/vehicle/service"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// NewControllerVehicle returns a new instance of a vehicle controller.
func NewControllerVehicle(st service.ServiceVehicle) *ControllerVehicle {
	return &ControllerVehicle{st: st}
}

// ControllerVehicle is an struct that represents a vehicle controller.
type ControllerVehicle struct {
	// StorageVehicle is the storage of vehicles.
	st service.ServiceVehicle
}

// GetAll returns all vehicles.
type VehicleHandlerGetAll struct {
	Id           int     `json:"id"`
	Brand        string  `json:"brand"`
	Model        string  `json:"model"`
	Registration string  `json:"registration"`
	Year         int     `json:"year"`
	Color        string  `json:"color"`
	MaxSpeed     int     `json:"max_speed"`
	FuelType     string  `json:"fuel_type"`
	Transmission string  `json:"transmission"`
	Passengers   int     `json:"passengers"`
	Height       float64 `json:"height"`
	Width        float64 `json:"width"`
	Weight       float64 `json:"weight"`
}
type ResponseBodyGetAll struct {
	Message string                  `json:"message"`
	Data    []*VehicleHandlerGetAll `json:"vehicles"`
	Error   bool                    `json:"error"`
}

func (c *ControllerVehicle) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		vehicles, err := c.st.GetAll()
		if err != nil {
			var code int
			var body ResponseBodyGetAll
			switch {
			case errors.Is(err, service.ErrServiceVehicleNotFound):
				code = http.StatusNotFound
				body = ResponseBodyGetAll{Message: "Not found", Error: true}
			default:
				code = http.StatusInternalServerError
				body = ResponseBodyGetAll{Message: "Internal server error", Error: true}
			}

			ctx.JSON(code, body)
			return
		}

		// response
		code := http.StatusOK
		body := ResponseBodyGetAll{Message: "Success", Data: make([]*VehicleHandlerGetAll, 0, len(vehicles)), Error: false}
		for _, vehicle := range vehicles {
			body.Data = append(body.Data, &VehicleHandlerGetAll{
				Id:           vehicle.Id,
				Brand:        vehicle.Attributes.Brand,
				Model:        vehicle.Attributes.Model,
				Registration: vehicle.Attributes.Registration,
				Year:         vehicle.Attributes.Year,
				Color:        vehicle.Attributes.Color,
				MaxSpeed:     vehicle.Attributes.MaxSpeed,
				FuelType:     vehicle.Attributes.FuelType,
				Transmission: vehicle.Attributes.Transmission,
				Passengers:   vehicle.Attributes.Passengers,
				Height:       vehicle.Attributes.Height,
				Width:        vehicle.Attributes.Width,
				Weight:       vehicle.Attributes.Weight,
			})
		}

		ctx.JSON(code, body)
		return
	}
}

func (c *ControllerVehicle) GetByBrandInTimeRange() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		brand := ctx.Param("brand")
		startYear, startYearErr := strconv.Atoi(ctx.Param("start_year"))
		endYear, endYearErr := strconv.Atoi(ctx.Param("end_year"))

		if brand == "" || startYearErr != nil || endYearErr != nil {
			code := http.StatusBadRequest
			body := ResponseBodyGetAll{Message: http.StatusText(code), Error: true}
			ctx.AbortWithStatusJSON(code, body)
			return
		}

		vehicles, err := c.st.GetByBrandInTimeRange(brand, startYear, endYear)

		if err != nil {
			var code int

			switch {
			case errors.Is(err, service.ErrServiceVehicleInvalidTimeRange):
				code = http.StatusBadRequest
			case errors.Is(err, service.ErrServiceVehicleNotFound):
				code = http.StatusNotFound
			default:
				code = http.StatusInternalServerError
			}

			body := ResponseBodyGetAll{Message: http.StatusText(code), Error: true}
			ctx.AbortWithStatusJSON(code, body)
			return
		}

		code := http.StatusOK
		body := ResponseBodyGetAll{
			Message: http.StatusText(code),
			Data:    make([]*VehicleHandlerGetAll, 0, len(vehicles)),
			Error:   false,
		}

		for _, vehicle := range vehicles {
			body.Data = append(body.Data, &VehicleHandlerGetAll{
				Id:           vehicle.Id,
				Brand:        vehicle.Attributes.Brand,
				Model:        vehicle.Attributes.Model,
				Registration: vehicle.Attributes.Registration,
				Year:         vehicle.Attributes.Year,
				Color:        vehicle.Attributes.Color,
				MaxSpeed:     vehicle.Attributes.MaxSpeed,
				FuelType:     vehicle.Attributes.FuelType,
				Transmission: vehicle.Attributes.Transmission,
				Passengers:   vehicle.Attributes.Passengers,
				Height:       vehicle.Attributes.Height,
				Width:        vehicle.Attributes.Width,
				Weight:       vehicle.Attributes.Weight,
			})
		}

		ctx.JSON(code, body)
		return
	}
}
