package database

import (
	"errors"
	"fmt"
	m "govulnapi/models"
	"time"
)

func (d *DB) AddOrder(userId int, coinId string, price float64, isBuy bool, qty float64) error {
	user, err := d.GetUserById(userId)
	if err != nil {
		return err
	}

	var (
		orderValue         = qty * price
		newUsdBalance      float64
		currentCoinBalance m.CoinBalance
		newCoinBalance     float64
	)

	for _, c := range user.CoinBalances {
		if c.CoinId == coinId {
			currentCoinBalance = c
		}
	}

	// CWE-20: Improper Input Validation
	// An adversary can create infinite orders with 0 qty, which could lead
	// to constant disk writes and filling of storage
	// if qty <= 0 {
	// 	return errors.New("Quantity needs to be > 0!")
	// }

	if isBuy {
		if user.UsdBalance < orderValue {
			return errors.New("Not enough usd!")
		}
		newUsdBalance = user.UsdBalance - orderValue
		newCoinBalance = currentCoinBalance.Qty + qty
	} else {
		if currentCoinBalance.Qty < qty {
			return errors.New("Not enough coin!")
		}
		newUsdBalance = user.UsdBalance + orderValue
		newCoinBalance = currentCoinBalance.Qty - qty
	}

	// CWE-89:  SQL Injection
	qAddOrder := fmt.Sprintf(
		"INSERT INTO 'order' (user_id, coin_id, price, is_buy, qty, date) VALUES ('%v','%v','%v','%v','%v','%v')",
		user.Id, coinId, price, isBuy, qty, time.Now(),
	)
	qUpdateFiat := fmt.Sprintf(
		"UPDATE 'user' SET usd_balance = %v WHERE id = %d",
		newUsdBalance, user.Id,
	)
	qUpdateCoinBalance := fmt.Sprintf(
		"UPDATE 'coin_balance' SET qty = %v WHERE user_id = %d AND coin_id = '%s'",
		newCoinBalance, user.Id, coinId,
	)

	tx, err := d.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if _, err = tx.Exec(qAddOrder); err != nil {
		return err
	}
	if _, err = tx.Exec(qUpdateFiat); err != nil {
		return err
	}
	if _, err = tx.Exec(qUpdateCoinBalance); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
