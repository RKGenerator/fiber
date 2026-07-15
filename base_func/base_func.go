package basefunc

import (
	"test-fiber/dto"
	apperrors "test-fiber/errors"
)

func CheckPaginationRequest(req dto.PaginatorRequest) error {
	if req.Page < 1 || req.PageSize < 1 {
		return apperrors.ErrBadRequest
	}
	return nil
}
