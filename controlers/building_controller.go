package controlers

import (
	"test-fiber/dto"
	"test-fiber/services"

	"github.com/gofiber/fiber/v3"
)

type BuildingController struct {
	BuildingService *services.BuildingService
}

func (c *BuildingController) GetReq(ctx fiber.Ctx) error {
	req := new(dto.BuildingRequest)
	err := ctx.Bind().Query(req)
	if err != nil {
		return err
	}
	dto, err := c.BuildingService.GetBuildings()
	if err != nil {
		return err
	}
	return ctx.JSON(dto)
}
