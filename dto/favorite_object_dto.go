package dto

type FavoriteRequest struct {
	ApartmentId *int `query:"id"`
}

type FavoritesRequest struct {
	ApartmentIds *[]int64 `query:"id"`
}

type FavoritesResponse struct {
	DeletedCount int64 `json:"deleted_count"`
}
