package models

type User struct {
	Id                 int     `db:"id"`
	Email              string  `db:"email"`
	Password           string  `db:"password"`
	UsdBalance         float64 `db:"usd_balance"`
	UsdStartingBalance float64 `db:"usd_starting_balance"`
	CoinBalances       []CoinBalance
	Transactions       []Transaction
	Orders             []Order
}

type CoinBalance struct {
	CoinId  string  `db:"coin_id"`
	Address string  `db:"address"`
	Qty     float64 `db:"qty"`
}

type Transaction struct {
	Id         int     `db:"id" swaggerignore:"true"`
	SenderId   int     `db:"sender_id" swaggerignore:"true"`
	ReceiverId int     `db:"receiver_id" swaggerignore:"true"`
	CoinId     string  `db:"coin_id" example:"bitcoin"`
	Address    string  `db:"address" example:""`
	Qty        float64 `db:"qty" example:"1"`
	Date       string  `db:"date" swaggerignore:"true"`
	Note       string  `db:"note"`
}

type Order struct {
	UserId int     `db:"user_id" swaggerignore:"true"`
	CoinId string  `db:"coin_id" example:"bitcoin"`
	Price  float64 `db:"price" swaggerignore:"true"`
	IsBuy  bool    `db:"is_buy"`
	Qty    float64 `db:"qty" example:"1"`
	Date   string  `db:"date" swaggerignore:"true"`
}
