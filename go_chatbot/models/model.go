package models

type UserCreate struct {
	Username string `json:"username" bson:"username" validate:"required"`
	Password string `json:"password" bson:"password" validate:"required"`
}

type UserLogin struct {
	Username string `json:"username" bson:"username" validate:"required"`
	Password string `json:"password" bson:"password" validate:"required"`
}

type UserGet struct {
	UserID string `json:"user_id" bson:"user_id" validate:"required"`
}
