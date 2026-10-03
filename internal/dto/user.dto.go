package dto

import "mime/multipart"

type UserProfile struct {
	Name          string  `json:"name"`
	Email 		  string  `json:"email"`
	Bio           *string `json:"bio"`
	User_location *string `json:"user_location"`
	Position      *string `json:"position"`
	Avatar_url    *string `json:"avatar_url"`
}

type SetUserProfile struct {
	Name          string                `form:"name"`
	Bio           *string               `form:"bio"`
	User_location *string               `form:"user_location"`
	Position      *string               `form:"position"`
	Avatar        *multipart.FileHeader `form:"avatar_url"`
}

type UserHeaderInfo struct {
	Name          string  `json:"name"`
	Email 		  string  `json:"email"`
	Avatar_url    *string `json:"avatar_url"`
}
