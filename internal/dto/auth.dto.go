package dto

type Register struct {
	FullName string `json:"fullname"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password"`
	// Term     bool   `json:"term"`
}

type Login struct {
	Email    string `json:"email" example:"alfan@gmail.com"`
	Password string `json:"password" example:"alfan123"`
}

type ChangePassword struct {
	NewPassword string `json:"newPassword"`
}
