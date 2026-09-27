package model

import "time"

type Notification struct {
	NotifId    int        `db:"id"`
	Account_id int        `db:"account_id"`
	Title      string     `db:"title"`
	Message    string     `db:"message"`
	Type       string     `db:"type"`
	Read_at    *time.Time `db:"read_at"`
	Created_at time.Time  `db:"created_at"`
	Updated_at time.Time  `db:"updated_at"`
}

type NotificationDetail struct {
	Account
	Notification
}
