package models

type EcomctxKey string


type AuthData struct {
	Authorization string `header:"Authorization"`
	ISsandBox     bool   `header:"IsSandBox"`
}