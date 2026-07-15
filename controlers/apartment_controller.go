package controlers

import (
	basefunc "test-fiber/base_func"
	"test-fiber/dto"
	"test-fiber/services"

	"github.com/gofiber/fiber/v3"
)

type ApartmentController struct {
	ApartmentsServices *services.ApartmentServices
}

func (c *ApartmentController) GetApartments(ctx fiber.Ctx) error {

	user, err := c.ApartmentsServices.GetApartments()
	if err != nil {
		ctx.Status(fiber.ErrBadRequest.Code)
		return err
	}
	return ctx.JSON(user)
}

func (c *ApartmentController) GetApartmentsDetail(ctx fiber.Ctx) error {
	req := new(dto.ApartmentDetailsRequest)
	if err := ctx.Bind().Query(req); err != nil {
		return err
	}

	err := basefunc.CheckPaginationRequest(req.PaginatorRequest)
	if err != nil {
		ctx.Status(fiber.ErrBadRequest.Code)
		return err
	}

	aparts, err := c.ApartmentsServices.GetApartmentsDetail(*req)
	if err != nil {
		ctx.Status(fiber.ErrBadRequest.Code)
		return err
	}
	return ctx.JSON(aparts)
}
