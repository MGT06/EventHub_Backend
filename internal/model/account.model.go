package model

import "time"

type Account struct {
	Id         int        `db:"id"`
	Name       string     `db:"name"`
	Email      string     `db:"email"`
	Bio        *string    `db:"bio"`
	User_location   *string    `db:"user_location"`
	Position   *string    `db:"position"`
	Password   string     `db:"password"`
	Role       string     `db:"role"`
	Avatar_url *string    `db:"avatar_url"`
	Status     string     `db:"status"`
	Terms      bool       `db:"terms"`
	Created_at time.Time  `db:"created_at"`
	Updated_at *time.Time `db:"updated_at"`
}
