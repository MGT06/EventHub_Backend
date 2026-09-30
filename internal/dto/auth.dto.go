package dto

type Register struct {
	FullName string `json:"fullname"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password"`
	// Term     bool   `json:"term"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ChangePassword struct {
	NewPassword string `json:"newPassword"`
}
