package controlers

import (
	"test-fiber/dto"
	"test-fiber/services"

	"github.com/gofiber/fiber/v3"
)

type Controllers struct {
	ApartmentsServices *services.ApartmentsServices
}

func (c *Controllers) GetApartments(ctx fiber.Ctx) error {

	user, err := c.ApartmentsServices.GetApartments()
	if err != nil {
		ctx.Status(fiber.ErrBadRequest.Code)
		return err
	}
	return ctx.JSON(user)
}

func (c *Controllers) GetReq(ctx fiber.Ctx) error {
	req := new(dto.ApartmentDetailsRequest)
	if err := ctx.Bind().Query(req); err != nil {
		return err
	}

	aparts, err := c.ApartmentsServices.GetReq(*req)
	if err != nil {
		ctx.Status(fiber.ErrBadRequest.Code)
		return err
	}
	return ctx.JSON(aparts)
}
