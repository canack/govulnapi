package database

import (
	"errors"
	m "govulnapi/models"
)

func (d *DB) GetCoins() ([]m.Coin, error) {
	var (
		coins []m.Coin
		query = "SELECT id FROM 'coin'"
	)

	if err := d.db.Select(&coins, query); err != nil {
		return nil, errors.New("Unable to load coins from the database!")
	}

	return coins, nil
}
