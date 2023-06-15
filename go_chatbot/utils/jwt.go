package utils

import (
	"fmt"
	"go_chatbot/repository"
	"go_chatbot/response"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

const (
	SecretKey = "secret"
)

type myJWTClaims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func JWT_NewWithClaims(user repository.Users) *jwt.Token {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, myJWTClaims{
		UserID:   user.UserID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    user.UserID,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		},
	})
}

func AuthorizationMiddleware(ctx *fiber.Ctx) error {
	s := ctx.Get("Authorization")
	tokenString := strings.TrimPrefix(s, "Bearer ")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SecretKey), nil
	})
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(response.Err_response(err))
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		ctx.Locals("user_id", claims["user_id"])
		ctx.Locals("username", claims["username"])
	}

	return ctx.Next()
}
