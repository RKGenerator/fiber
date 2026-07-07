package main

import (
	"log"
	"test-fiber/controlers"
	"test-fiber/repositories"
	"test-fiber/services"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()

	//db, err := pgx.Connect(context.Background(), "postgres://postgres:qwerty@localhost:5432/cremiabuildings")
	dsn := "host=localhost user=postgres password=qwerty dbname=cremiabuildings port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	userRep := repositories.NewApartmentsRepository(db)

	apartmentService := services.NewApartmentsService(userRep)
	controllers := controlers.Controllers{
		ApartmentsServices: apartmentService,
	}

	app.Use(logger.New())
	app.Get("/apartments", controllers.GetApartments)
	app.Get("/req", controllers.GetReq)
	app.Listen(":8000")
}
