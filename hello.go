package main

import (
	"context"
	"log"
	"test-fiber/controlers"
	"test-fiber/repositories"
	"test-fiber/services"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/jackc/pgx/v5"
)

func main() {
	app := fiber.New()

	db, err := pgx.Connect(context.Background(), "postgres://postgres:qwerty123@localhost:5432/table")
	if err != nil {
		log.Fatal(err)
	}
	userRep := repositories.NewApartmentsRepository(db)

	apartmentService := services.NewUserService(userRep)
	controllers := controlers.Controllers{
		ApartmentsServices: apartmentService,
	}
	app.Use(logger.New())
	app.Get("/hello", controllers.HelloHandler)
	app.Get("/user", controllers.GetUser)
	app.Get("/bedrooms", controllers.GetBedrooms)
	app.Listen(":8000")
}
