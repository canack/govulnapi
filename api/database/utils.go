package database

import (
	"crypto/md5"
	"errors"
	"fmt"
	"net/mail"
)

func md5sum(password string) string {
	// CWE-327: Use of a Broken or Risky Cryptographic Algorithm
	// CWE-328: Use of Weak Hash
	// CWE-759: Use of a One-Way Hash without a Salt
	// CWE-916: Use of Password Hash With Insufficient Computational Effort
	password = fmt.Sprintf("%x", md5.Sum([]byte(password)))
	return password
}

func validateEmail(email string) error {
	if _, err := mail.ParseAddress(email); err != nil {
		return errors.New("Email invalid!")
	}

	return nil
}
