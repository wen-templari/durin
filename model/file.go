package model

type FileDTO struct {
	Path string `json:"path" redis:"path"`
}
