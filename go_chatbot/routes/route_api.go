package routes

import (
	"go_chatbot/controller"
	"go_chatbot/repository"
	"go_chatbot/services"
	"go_chatbot/utils"
	"go_chatbot/ws"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func Routers(app *fiber.App, db *mongo.Collection) {
	userRpository := repository.NewUserReoository(db)
	usrService := services.NewUserService(userRpository)
	userController := controller.NewUserController(usrService)

	app.Get("/websoc", websocket.New(ws.GetWs))

	app.Post("/user/create", userController.CreateUser_api)
	app.Post("/user/Login", userController.LoginUser_api)
	app.Post("/websoc", userController.Post_ws)

	auth := app.Group("", utils.AuthorizationMiddleware)
	auth.Get("/user/get", userController.GetUser_api)

}
