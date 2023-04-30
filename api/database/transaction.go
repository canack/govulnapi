package database

import (
	"errors"
	"fmt"
	m "govulnapi/models"
	"time"
)

func (d *DB) AddTransaction(userId int, coinId string, address string, qty float64) error {
	user, err := d.GetUserById(userId)
	if err != nil {
		return err
	}

	var senderBalance m.CoinBalance
	for _, balance := range user.CoinBalances {
		if balance.CoinId == coinId {
			senderBalance = balance
			break
		}
	}

	if senderBalance.CoinId == "" {
		return errors.New("Coin with requested id doesn't exist!")
	}

	if qty <= 0 {
		return errors.New("Quantity needs to be > 0!")
	}

	if senderBalance.Qty < qty {
		return errors.New("Not enough coin!")
	}

	// CWE-89:  SQL Injection
	qBalanceReceiver := fmt.Sprintf(
		"UPDATE 'coin_balance' SET qty=qty+%v WHERE address='%s'",
		qty, address,
	)
	qBalanceSender := fmt.Sprintf(
		"UPDATE 'coin_balance' SET qty=qty-%v WHERE address='%s'",
		qty, senderBalance.Address,
	)
	qTransaction := fmt.Sprintf(
		"INSERT INTO 'transaction' (user_id,coin_id,address,qty,date) VALUES (%d,'%v','%s',%v,'%v')",
		user.Id, coinId, address, qty, time.Now(),
	)

	tx, err := d.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	r, _ := tx.Exec(qBalanceReceiver)
	rows, _ := r.RowsAffected()
	if rows == 0 {
		return errors.New("Receiver address doesn't exist!")
	}

	if _, err = tx.Exec(qBalanceSender); err != nil {
		return err
	}
	if _, err = tx.Exec(qTransaction); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
