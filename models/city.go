package models

import (
	"context"
	"database/sql"
	"fmt"
	"unsia/pb/cities"
)

type City struct {
	Pb cities.City
}

func (u *City) Get(ctx context.Context, db *sql.DB, in *cities.Id) error {
	query := `SELECT id, name FROM cities WHERE id = $1;`
	err := db.QueryRowContext(ctx, query, in.Id).Scan(&u.Pb.Id, &u.Pb.Name)
	if err != nil {
		return err
	}
	return nil
}

func (u *City) Create(ctx context.Context, db *sql.DB, in *cities.CityInput) error {
	query := `INSERT INTO cities (name) VALUES ($1) RETURNING id;`
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	err = stmt.QueryRowContext(ctx, in.Name).Scan(&u.Pb.Id)
	if err != nil {
		return err
	}

	u.Pb.Name = in.Name
	
	return nil

func (u *City) Delete(ctx context.Context, db *sql.DB, in *cities.Id) error {
	query := `DELETE FROM cities WHERE id = $1;`
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	rs, err:= stmt.ExecContext(ctx, in.Id)
	if err != nil {
		return err
	}

	affected, err := rs.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
	return fmt.Errorf("DATA NOT FOUND")
	}

	return nil
}
