package database

import (
	"encoding/base64"
	"errors"
	"fmt"
	"log"

	m "govulnapi/models"
)

func (d *DB) getUser(queryUser string) (m.User, error) {
	// This function is inefficient as it fetches all user data
	// (even when not called for), but made this way for simplicity

	// Get user
	var user m.User
	if err := d.db.Get(&user, queryUser); err != nil {
		return m.User{}, err
	}

	// CWE-89:  SQL Injection
	qBalances := fmt.Sprintf("SELECT coin_id, address, qty FROM 'coin_balance' WHERE user_id = %d", user.Id)
	qOrders := fmt.Sprintf("SELECT coin_id, price, is_buy, qty, date FROM 'order' WHERE user_id = %d", user.Id)
	qTransactions := fmt.Sprintf("SELECT * FROM 'transaction' WHERE sender_id = %d OR receiver_id = %d", user.Id, user.Id)

	d.db.Select(&user.CoinBalances, qBalances)     // Get user balances
	d.db.Select(&user.Orders, qOrders)             // Get user orders
	d.db.Select(&user.Transactions, qTransactions) // Get user transactions

	return user, nil
}

func (d *DB) GetUserByCredentials(email string, password string) (m.User, error) {
	password = md5sum(password)

	// CWE-89:  SQL Injection
	query := fmt.Sprintf("SELECT * FROM 'user' WHERE user.email = '%s' and user.password = '%s'", email, password)

	user, err := d.getUser(query)
	if err != nil {
		return m.User{}, errors.New("No user with matching credentials found!")
	}

	return user, nil
}

func (d *DB) GetUserByEmail(email string) (m.User, error) {
	// CWE-89:  SQL Injection
	query := fmt.Sprintf("SELECT * FROM 'user' WHERE user.email = '%s'", email)

	user, err := d.getUser(query)
	if err != nil {
		return m.User{}, errors.New("No user with matching email found!")
	}

	return user, nil
}

func (d *DB) GetUserById(userId int) (m.User, error) {
	// CWE-89:  SQL Injection
	query := fmt.Sprintf("SELECT * FROM 'user' WHERE user.id = %d", userId)

	user, err := d.getUser(query)
	if err != nil {
		return m.User{}, errors.New("No user with matching id found!")
	}

	return user, nil
}

func (d *DB) AddUser(email string, password string) error {
	if err := validateEmail(email); err != nil {
		return err
	}

	if _, err := d.GetUserByEmail(email); err == nil {
		return errors.New("Email already registered!")
	}

	// CWE-521: Weak Password Requirements
	if len(password) < 6 {
		return errors.New("Password needs to be at least 6 characters long!")
	}

	hashedPassword := md5sum(password)

	// CWE-89:  SQL Injection
	query := fmt.Sprintf("INSERT INTO 'user' (email, password) VALUES ('%s', '%s')", email, hashedPassword)
	r, err := d.db.Exec(query)
	if err != nil {
		return err
	}
	user_id, _ := r.LastInsertId()

	coins, err := d.GetCoins()
	if err != nil {
		return err
	}

	// Initialize empty balances for every coin
	for _, coin := range coins {
		addressData := fmt.Sprintf("%v-%v-%v", coin.Id, email, user_id)
		address := base64.StdEncoding.EncodeToString([]byte(addressData))

		// CWE-89:  SQL Injection
		query = fmt.Sprintf(
			"INSERT INTO 'coin_balance' (user_id, coin_id, address, qty) VALUES (%d,'%v','%v',%v)",
			user_id, coin.Id, address, 0.0,
		)
		d.db.Exec(query)
	}

	// CWE-532: Insertion of Sensitive Information into Log File
	log.Printf("Registered user: email: '%s', password: '%s'\n", email, password)

	return nil
}

func (d *DB) UpdateEmail(userId int, newEmail string) error {
	// if err := validateEmail(newEmail); err != nil {
	// 	return err
	// }

	// CWE-89:  SQL Injection
	query := fmt.Sprintf("UPDATE 'user' SET email=%s WHERE id=%d", newEmail, userId)

	_, err := d.db.Exec(query)
	if err != nil {
		return err
	}

	// CWE-223: Omission of Security-relevant Information
	log.Println("Updated email for user")
	return nil
}

func (d *DB) UpdatePassword(userId int, newPassword string) error {
	newPassword = md5sum(newPassword)

	// CWE-89:  SQL Injection
	query := fmt.Sprintf("UPDATE 'user' SET password='%s' WHERE id=%d", newPassword, userId)

	_, err := d.db.Exec(query)
	if err != nil {
		return err
	}

	// CWE-778: Insufficient Logging
	// log.Printf("Updated password for user %d\n", userId)
	return nil
}
