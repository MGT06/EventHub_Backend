package model

import "time"

type Testimony struct {
	Id         int       `db:"id"`
	Account_id string    `db:"account_id"`
	Message    string    `db:"message"`
	Created_at time.Time `db:"created_at"`
}

type TestimonyDetail struct {
	Account
	Testimony
}