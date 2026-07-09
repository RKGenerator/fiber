package dto

type PaginatorRequest struct {
	Page     int64 `query:"page"`
	PageSize int64 `query:"page_size"`
}
