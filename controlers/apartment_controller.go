package controlers

import (
	"test-fiber/auth"
	basefunc "test-fiber/base_func"
	"test-fiber/dto"
	"test-fiber/services"

	jwtware "github.com/gofiber/contrib/v3/jwt"
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

	user := jwtware.FromContext(ctx)

	var aparts []*dto.ApartmentDetailsResponse

	if user == nil {
		aparts, err = c.ApartmentsServices.GetApartmentsDetailToAnonym(*req)
	} else {
		claims := user.Claims.(*auth.AuthClaims)
		aparts, err = c.ApartmentsServices.GetApartmentsDetailToUser(claims.UserId, *req)
	}

	if err != nil {
		ctx.Status(fiber.ErrBadRequest.Code)
		return err
	}

	return ctx.JSON(aparts)
}

func (c *ApartmentController) PostFavorite(ctx fiber.Ctx) error {
	req := new(dto.FavoriteRequest)
	if err := ctx.Bind().Query(req); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	if req.ApartmentId == nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	user := jwtware.FromContext(ctx)
	claims := user.Claims.(*auth.AuthClaims)
	err := c.ApartmentsServices.AddFavoriteApartment(int64(*req.ApartmentId), claims.UserId)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}

func (c *ApartmentController) DeleteFavorites(ctx fiber.Ctx) error {
	req := new(dto.FavoritesRequest)
	if err := ctx.Bind().Query(req); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	if req.ApartmentIds == nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	user := jwtware.FromContext(ctx)
	claims := user.Claims.(*auth.AuthClaims)
	dto, err := c.ApartmentsServices.DeleteFavoriteApartments(*req.ApartmentIds, claims.UserId)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.JSON(dto)
}
