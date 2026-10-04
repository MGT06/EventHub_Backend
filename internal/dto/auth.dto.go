package dto

type Register struct {
	FullName string `json:"fullname" binding:"required,min=5"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	// Term     bool   `json:"term"`
}

type Login struct {
	Email    string `json:"email" binding:"required,email" example:"alfan@gmail.com"`
	Password string `json:"password" binding:"required" example:"alfan123"`
}

type ChangePassword struct {
	Email           string `json:"email" binding:"required,email"`
	CurrentPassword string `json:"currentPassword" binding:"required"`
	NewPassword     string `json:"newPassword" binding:"required,min=8"`
}
