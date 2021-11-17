package handler

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"mesage"`
	Data    interface{} `json:"data"`
}
