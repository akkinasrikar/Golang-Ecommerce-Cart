package models

type SignUp struct {
	Name     string `json:"user_name" binding:"required"`
	Email    string `json:"user_email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Login struct {
	Name     string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}
