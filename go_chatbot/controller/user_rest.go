package controller

import (
	"fmt"
	"go_chatbot/models"
	"go_chatbot/response"
	"go_chatbot/services"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gorilla/websocket"
)

type UserController interface {
	CreateUser_api(ctx *fiber.Ctx) error
	LoginUser_api(ctx *fiber.Ctx) error
	GetUser_api(ctx *fiber.Ctx) error
	Post_ws(ctx *fiber.Ctx) error
}

type userController struct {
	userSrv services.UserService
}

func NewUserController(userSrv services.UserService) UserController {
	return userController{userSrv}
}

func (obj userController) Post_ws(ctx *fiber.Ctx) error {
	dial, _, err := websocket.DefaultDialer.Dial("ws://localhost:8000/websoc", nil)

	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusInternalServerError).JSON(response.Err_response(err))
	}
	defer dial.Close()

	err = dial.WriteMessage(websocket.TextMessage, []byte("hello_5555555"))
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusInternalServerError).JSON(response.Err_response(err))
	}
	return ctx.Status(http.StatusOK).JSON(response.Response(fiber.Map{}, "Post success"))
}

func (obj userController) CreateUser_api(ctx *fiber.Ctx) error {
	user := models.UserCreate{}
	if err := ctx.BodyParser(&user); err != nil {
		log.Println("error", err)
		return ctx.Status(http.StatusBadRequest).JSON(response.Err_response(err))
	}
	validate := validator.New()
	if err := validate.Struct(&user); err != nil {
		log.Println("error", err)
		return ctx.Status(http.StatusBadRequest).JSON(response.Err_response(err))
	}
	err := obj.userSrv.CreateUser(user)
	if err != nil {
		log.Println("error", err)
		return ctx.Status(http.StatusInternalServerError).JSON(response.Err_response(err))
	}
	log.Println("success", ctx.Status(http.StatusCreated))
	return ctx.Status(http.StatusCreated).JSON(response.Response(fiber.Map{}, "Create success"))

}
func (obj userController) LoginUser_api(ctx *fiber.Ctx) error {
	user := models.UserLogin{}
	if err := ctx.BodyParser(&user); err != nil {
		log.Println("error", err)
		return ctx.Status(http.StatusBadRequest).JSON(response.Err_response(err))
	}
	validate := validator.New()
	if err := validate.Struct(&user); err != nil {
		log.Println("error", err)
		return ctx.Status(http.StatusBadRequest).JSON(response.Err_response(err))
	}
	token, err := obj.userSrv.LoginUser(user)
	if err != nil {
		log.Println("error", err)
		return ctx.Status(http.StatusInternalServerError).JSON(response.Err_response(err))
	}
	log.Println("success", ctx.Status(http.StatusOK))
	return ctx.Status(http.StatusOK).JSON(response.Response(fiber.Map{"data": token}, "Login success"))
}

func (obj userController) GetUser_api(ctx *fiber.Ctx) error {

	user := models.UserGet{}
	if err := ctx.BodyParser(&user); err != nil {
		log.Println("error", err)
		return ctx.Status(http.StatusBadRequest).JSON(response.Err_response(err))
	}
	validate := validator.New()
	if err := validate.Struct(&user); err != nil {
		log.Println("error", err)
		return ctx.Status(http.StatusBadRequest).JSON(response.Err_response(err))
	}

	id_token := ctx.Locals("user_id")
	if user.UserID != id_token {
		err := fmt.Errorf("token or user_id invalid")
		return ctx.Status(http.StatusBadRequest).JSON(response.Err_response(err))
	}

	getUser, err := obj.userSrv.GetUser(user.UserID)
	if err != nil {
		log.Println("error", err)
		return ctx.Status(http.StatusInternalServerError).JSON(response.Err_response(err))
	}
	log.Println("success", ctx.Status(http.StatusOK))
	return ctx.Status(http.StatusOK).JSON(response.Response(fiber.Map{"data": getUser}, "Get success"))
}
