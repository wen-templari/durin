package model

type Account struct {
	Id       string `json:"id" redis:"id"`
	Name     string `json:"name" redis:"name"`
	Password string `json:"password" redis:"password"`
	Avatar   string `json:"avatar" redis:"avatar"`
}

type AccountDTO struct {
	Id             string    `json:"id" redis:"id"`
	Name           string    `json:"name" redis:"name"`
	Status         int       `json:"status"` //1 for online and 0 for offline
	Avatar         string    `json:"avatar" redis:"avatar"`
	Token          string    `json:"token"`
	OfflineMessage []Message `json:"offlineMessage"`
}
