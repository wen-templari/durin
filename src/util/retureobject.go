package util

type ReturnObject struct {
	ResCode int         `json:"resCode"`
	ResMsg  string      `json:"resMsg"`
	Data    interface{} `json:"data,omitempty"`
}

func NewReturnObject(resCode int, resMsg string, data interface{}) *ReturnObject {
	return &ReturnObject{
		ResCode: resCode,
		ResMsg:  resMsg,
		Data:    data,
	}
}
