package models

type SignUp struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Login struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}
