package main

import (
	"fmt"
	"log"
	"test-fiber/controlers"
	"test-fiber/model"
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
	dsn := "host=localhost user=postgres password=qwerty123 dbname=table port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	userRep := repositories.NewApartmentsRepository(db)

	apartmentService := services.NewApartmentsService(userRep)
	controllers := controlers.Controllers{
		ApartmentsServices: apartmentService,
	}

	var apartments []*model.Apartment

	err = db.Preload("PropertyImages").Find(&apartments).Error
	if err != nil {
		log.Println(err)
	}

	apart := apartments[0]
	fmt.Println(apart)

	app.Use(logger.New())
	app.Get("/hello", controllers.HelloHandler)
	app.Get("/apartments", controllers.GetApartments)
	app.Get("/apartmentsE", controllers.GetApartmentsExpirence)
	app.Get("/bedrooms", controllers.GetBedrooms)
	app.Listen(":8000")
}
