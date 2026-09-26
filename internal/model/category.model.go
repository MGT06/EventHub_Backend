package model

import "time"

type category struct {
	Id           int    `db:"id"`
	Category_name string `db:"category_name"`
	Created_at   time.Time `db:"created_at"`
}
