package main

import (
	"app/cmd/handlers"
	"app/internal/vehicle/loader"
	"app/internal/vehicle/repository"
	"app/internal/vehicle/service"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// env
	godotenv.Load(".env")

	// dependencies
	ldVh := loader.NewLoaderVehicleJSON(os.Getenv("FILE_PATH_VEHICLES_JSON"))
	dbVh, err := ldVh.Load()
	if err != nil {
		panic(err)
	}

	rpVh := repository.NewRepositoryVehicleInMemory(dbVh)
	svVh := service.NewServiceVehicleDefault(rpVh)
	ctVh := handlers.NewControllerVehicle(svVh)

	// server
	rt := gin.New()
	// -> middlewares
	rt.Use(gin.Recovery())
	rt.Use(gin.Logger())
	// -> handlers
	api := rt.Group("/api/v1")
	grVh := api.Group("/vehicles")
	grVh.GET("", ctVh.GetAll())
	grVh.GET("/brand/:brand/between/:start_year/:end_year", ctVh.GetByBrandInTimeRange())

	// run
	if err := rt.Run(os.Getenv("SERVER_ADDR")); err != nil {
		panic(err)
	}
}
