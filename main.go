package main

import (
	"database/sql"
	"embed"
	"fmt"
	"log"
	"os"
	"test-fiber/auth"
	"test-fiber/controlers"
	"test-fiber/model"
	"test-fiber/repositories"
	"test-fiber/services"

	jwtware "github.com/gofiber/contrib/v3/jwt"

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
		return fmt.Errorf("could not run migrate: %w", err)
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

	apartRep := repositories.NewApartmentsRepository(db)
	apartmentService := services.NewApartmentsService(apartRep)
	apartmentControler := controlers.ApartmentController{
		ApartmentsServices: apartmentService,
	}

	buildRep := repositories.NewBuildingRepositories(db)
	buildingService := services.NewBuildingService(buildRep)
	buildingControler := controlers.BuildingController{
		BuildingService: buildingService,
	}

	userRep := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRep)
	userController := controlers.UserController{
		UserService: userService,
	}

	app.Use(logger.New())

	app.Post("/login", userController.UserLogin)

	app.Get("/apartments", apartmentControler.GetApartments)
	app.Get("/apartments/search", apartmentControler.GetApartmentsDetail)
	app.Get("/buildings", buildingControler.GetBuildingDetail)

	app.Get("/logint", func(c fiber.Ctx) error {
		tokenString, err := auth.GenerateToken(model.User{})
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Could not generate token",
			})
		}

		return c.JSON(fiber.Map{"token": tokenString})
	})

	test := app.Group("/test")
	test.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
		Claims:     &auth.AuthClaims{},
	}))

	test.Get("/i", func(c fiber.Ctx) error {
		user := jwtware.FromContext(c)
		claims := user.Claims.(*auth.AuthClaims)
		username := claims.UserId

		return c.JSON(fiber.Map{
			"message": "Welcome to your profile!",
			"user":    username,
		})
	})

	api := app.Group("/api")
	api.Use(auth.OptionalJWTMiddleware())
	api.Post("/login", userController.UserLogin)
	api.Get("/i", func(c fiber.Ctx) error {
		user := jwtware.FromContext(c)
		claims := user.Claims.(*auth.AuthClaims)
		username := claims.UserId

		return c.JSON(fiber.Map{
			"message": "Welcome to your profile!",
			"user":    username,
		})
	})

	app.Listen(":8000")
}
