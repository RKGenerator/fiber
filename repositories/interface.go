package repositories

type ApartmentsInterface interface {
	GetApartments() (*int, error)
	GetNBedrooms() (int, error)
}
