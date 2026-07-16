package controlers

import (
	basefunc "test-fiber/base_func"
	"test-fiber/dto"
	"test-fiber/services"

	"github.com/gofiber/fiber/v3"
)

type BuildingController struct {
	BuildingService *services.BuildingService
}

func (c *BuildingController) GetBuildingDetail(ctx fiber.Ctx) error {
	req := new(dto.BuildingRequest)
	err := ctx.Bind().Query(req)
	if err != nil {
		return err
	}

	err = basefunc.CheckPaginationRequest(req.PaginatorRequest)
	if err != nil {
		return fiber.ErrBadRequest
	}

	dto, err := c.BuildingService.GetBuildingsDetail()
	if err != nil {
		return err
	}
	return ctx.JSON(dto)
}
