package dto

type Register struct {
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	// Term     bool   `json:"term"`
}
