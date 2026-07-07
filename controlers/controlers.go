package controlers

import (
	"fmt"
	"test-fiber/services"

	"github.com/gofiber/fiber/v3"
)

type Controllers struct {
	ApartmentsServices *services.ApartmentsServices
}

func (c *Controllers) HelloHandler(ctx fiber.Ctx) {
	fmt.Println(string(ctx.Body()))
	ctx.JSON(ctx.Body())
}

func (c *Controllers) GetApartments(ctx fiber.Ctx) error {

	user, err := c.ApartmentsServices.GetApartments()
	if err != nil {
		ctx.Status(fiber.ErrBadRequest.Code)
		return err
	}
	return ctx.JSON(user)
}

func (c *Controllers) GetApartmentsExpirence(ctx fiber.Ctx) error {

	user, err := c.ApartmentsServices.GetApartmentsExpirence()
	if err != nil {
		ctx.Status(fiber.ErrBadRequest.Code)
		return err
	}
	return ctx.JSON(user)
}

func (c *Controllers) GetBedrooms(ctx fiber.Ctx) error {

	n_bedrooms, err := c.ApartmentsServices.GetBedrooms()
	if err != nil {
		ctx.Status(fiber.ErrBadRequest.Code)
		return err
	}
	return ctx.JSON(n_bedrooms)
}
