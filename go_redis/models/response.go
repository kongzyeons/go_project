package models

type messageResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Count   int         `json:"count"`
	Data    interface{} `json:"data"`
}

func Response(result interface{}, count int, massage string) messageResponse {
	return messageResponse{
		Status:  true,
		Message: massage,
		Count:   count,
		Data:    result,
	}
}

func Err_response(err error) messageResponse {
	var data interface{}
	return messageResponse{
		Status:  false,
		Message: err.Error(),
		Data:    data}
}
