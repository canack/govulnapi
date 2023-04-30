package models

type Coin struct {
	Id    string `db:"id"`
	Price float64
}
