package main

import (
	"fmt"
	"go_chatbot/database"
	"go_chatbot/routes"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/app/config")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func main() {
	dsn := fmt.Sprintf("%v", viper.GetString("db.host"))
	db := database.NewMongoDB(dsn)
	app := fiber.New()
	routes.Routers(app, db)
	go app.Listen("localhost:8000")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
		// w,r IS WRITE AND DELETE YOU INDEX.HTML

	})
	println("You server run 8080")
	http.ListenAndServe(":8081", nil)

}
