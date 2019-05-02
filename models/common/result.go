package common

type Result struct {
	Status string `json:"status"`
	Data   interface{} `json:"info"`
	ErrMsg string `json:"err_msg"`
}
