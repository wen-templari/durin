package model

type Account struct {
	Id       string `json:"id" redis:"id"`
	Name     string `json:"name" redis:"name"`
	Password string `json:"password" redis:"password"`
}

type AccountDTO struct {
	Id             string    `json:"id" redis:"id"`
	Name           string    `json:"name" redis:"name"`
	Status         int       `json:"status"` //1 for online and 0 for offline
	Token          string    `json:"token"`
	OfflineMessage []Message `json:"offlineMessage"`
}
