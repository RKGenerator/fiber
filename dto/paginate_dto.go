package dto

type PaginatorRequest struct {
	Page     int     `query:"page"`
	PageSize int     `query:"page_size"`
	Order    *string `query:"order"`
}
