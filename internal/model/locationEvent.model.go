package model

import "time"

type locationEvent struct {
	LocationId int       `db:"id"`
	City       string    `db:"city"`
	Created_at time.Time `db:"created_at"`
}
