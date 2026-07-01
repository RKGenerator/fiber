package repositories

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

type ApartmentsRepository struct {
	conn *pgx.Conn
}

func NewApartmentsRepository(conn *pgx.Conn) *ApartmentsRepository {
	return &ApartmentsRepository{
		conn: conn,
	}
}

func (r *ApartmentsRepository) GetApartments() (*int, error) {
	var id int
	rows, err := r.conn.Query(context.Background(), "SELECT id FROM parser_apartment where id = 3")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
	}
	return &id, nil
}

func (r *ApartmentsRepository) GetNBedrooms() (int, error) {
	var n_bedrooms int
	row := r.conn.QueryRow(context.Background(), "SELECT id FROM parser_apartment WHERE n_bedrooms = $1", 2)

	err := row.Scan(&n_bedrooms)
	if err != nil {
		log.Fatal(err)
	}

	return n_bedrooms, nil
}
