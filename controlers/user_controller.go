package controlers

import (
	"test-fiber/dto"
	"test-fiber/services"

	"github.com/gofiber/fiber/v3"
)

type UserController struct {
	UserService *services.UserService
}

func (c *UserController) UserLogin(ctx fiber.Ctx) error {
	req := new(dto.AuthRequest)
	err := ctx.Bind().Query(req)
	if err != nil {
		return err
	}

	if req.Email == "" || req.Password == "" {
		return fiber.ErrBadRequest
	}

	dto, err := c.UserService.UserLogin(*req)
	if err != nil {
		return err
	}

	return ctx.JSON(dto)

}
