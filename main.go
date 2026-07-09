package main

import (
	"database/sql"
	"embed"
	"fmt"
	"log"
	"test-fiber/controlers"
	"test-fiber/repositories"
	"test-fiber/services"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"

	"github.com/golang-migrate/migrate/v4"
	migratepostgres "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

//go:embed db/migrations/*.sql
var migrationFiles embed.FS

func runMigrations(db *sql.DB) error {
	driver, err := migratepostgres.WithInstance(db, &migratepostgres.Config{})
	if err != nil {
		return fmt.Errorf("could not create drover: %w", err)
	}

	src, err := iofs.New(migrationFiles, "db/migrations")
	if err != nil {
		return fmt.Errorf("could not create iofs: %w", err)
	}

	m, err := migrate.NewWithInstance("iofs", src, "postgres", driver)
	if err != nil {
		return fmt.Errorf("could not create migrate: %w", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("couLd not run migrate: %w", err)
	}

	log.Println("Migrations applied succesfully")

	return nil
}

func main() {
	app := fiber.New(fiber.Config{
		EnableSplittingOnParsers: true,
	})

	//db, err := pgx.Connect(context.Background(), "postgres://postgres:qwerty@localhost:5432/cremiabuildings")
	dsn := "host=localhost user=postgres password=qwerty dbname=cremiabuildings port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()
	if err := runMigrations(sqlDB); err != nil {
		log.Fatal("Migration failed: ", err)
	}
	ApartRep := repositories.NewApartmentsRepository(db)

	apartmentService := services.NewApartmentsService(ApartRep)
	apartmentControler := controlers.ApartmentController{
		ApartmentsServices: apartmentService,
	}

	buildRep := repositories.NewBuildingRepositories(db)
	buildingService := services.NewBuildingService(buildRep)
	buildingControler := controlers.BuildingController{
		BuildingService: buildingService,
	}

	app.Use(logger.New())
	app.Get("/apartments", apartmentControler.GetApartments)
	app.Get("/apartments/search", apartmentControler.GetApartmentsDetail)
	app.Get("/buildings", buildingControler.GetReq)
	app.Listen(":8000")

}
