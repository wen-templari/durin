package model

type Account struct {
	Id       string `json:"id" redis:"id"`
	Name     string `json:"name" redis:"name"`
	Password string `json:"password" redis:"password"`
}

type AccountDTO struct {
	Id       string `json:"id" redis:"id"`
	Name     string `json:"name" redis:"name"`
	Password string `json:"password" redis:"password"`
}
