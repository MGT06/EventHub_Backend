package dto

type UserProfile struct {
	Name          string  `json:"name"`
	Bio           *string `json:"bio"`
	User_location *string `json:"user_location"`
	Position      *string `json:"position"`
	Avatar_url    *string `json:"avatar_url"`
}


