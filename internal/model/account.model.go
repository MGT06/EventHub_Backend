package model

import "time"

type Account struct {
	Id         int        `db:"id"`
	Name       string     `db:"name"`
	Email      string     `db:"email"`
	Bio        *string    `db:"bio"`
	Location   *string    `db:"location"`
	Position   *string    `db:"position"`
	Password   string     `db:"password"`
	Role       string     `db:"role"`
	Image_url  *string    `db:"image_url"`
	Created_at time.Time  `db:"created_at"`
	Updated_at *time.Time `db:"updated_at"`
}
