package myopenid

import (
	"errors"
	"fmt"
	"log"
	"time"
)

type MysqlNonceStore struct{}

func (d MysqlNonceStore) Accept(endpoint, nonce string) error {
	maxNonceAge := 60 * time.Second
	// Value: A string 255 characters or less in length, that MUST be
	// unique to this particular successful authentication response.
	if len(nonce) < 20 || len(nonce) > 256 {
		return errors.New("Invalid nonce")
	}

	// The nonce MUST start with the current time on the server, and MAY
	// contain additional ASCII characters in the range 33-126 inclusive
	// (printable non-whitespace characters), as necessary to make each
	// response unique. The date and time MUST be formatted as specified in
	// section 5.6 of [RFC3339], with the following restrictions:

	// All times must be in the UTC timezone, indicated with a "Z".  No
	// fractional seconds are allowed For example:
	// 2005-05-15T17:11:51ZUNIQUE
	ts, err := time.Parse(time.RFC3339, nonce[0:20])
	if err != nil {
		return err
	}
	now := time.Now()
	diff := now.Sub(ts)
	if diff > maxNonceAge {
		return fmt.Errorf("Nonce too old: %ds", diff.Seconds())
	}

	s := nonce[20:]

	rows, err := db.Query("SELECT nonce, endpoint, time FROM noncestore WHERE endpoint = ?", endpoint)
	_, err = db.Exec("DELETE FROM noncestore")
	defer rows.Close()
	if err != nil {
		log.Printf("\nError 1: %s", err.Error())
	}

	noRows := 0
	for rows.Next() {
		noRows++
		var storedNonce, storedEndpoint string
		var nonceTime time.Time
		err := rows.Scan(&storedNonce, &storedEndpoint, &nonceTime)
		if err != nil {
			log.Printf("\nError 2: %s", err.Error())
		}

		if ts == nonceTime && storedNonce == s {
			return errors.New("Nonce already used")
		}
		if now.Sub(nonceTime) < maxNonceAge {
			_, err := db.Exec("INSERT INTO noncestore SET endpoint=?, time=?, nonce=?", endpoint, ts, s)
			if err != nil {
				log.Printf("\nError 3: %s", err.Error())

			}
		}
	}
	if noRows == 0 {
		_, err := db.Exec("INSERT INTO noncestore SET endpoint=?, time=?, nonce=?", endpoint, ts, s)
		if err != nil {
			log.Printf("\nError 4: %s", err.Error())
		}
	}

	return nil
}
