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

type RequestForgotPassword struct {
	Email string `json:"email" binding:"required,email"`
}

type ResponseForgotPassword struct {
	UserId int `json:"id"`
}

type ResetPassword struct {
	UserId int `json:"id" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required,min=8"`
}