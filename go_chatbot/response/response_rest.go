package response

import (
	"github.com/gofiber/fiber/v2"
)

type messageResponse struct {
	Status  bool       `json:"status"`
	Message string     `json:"message"`
	Data    *fiber.Map `json:"data"`
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func Response(result fiber.Map, massage string) messageResponse {
	return messageResponse{
		Status:  true,
		Message: massage,
		Data:    &result,
	}
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func Err_response(err error) messageResponse {
	data := &fiber.Map{"data": err.Error()}
	return messageResponse{
		Status:  false,
		Message: "error",
		Data:    data}
}
