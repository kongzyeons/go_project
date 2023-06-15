package repository

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Create(user Users) error
	GetByID(getuser Users) (user Users, err error)
	GetAll() (users []Users, err error)
}

type userRepository struct {
	db *mongo.Collection
}

func NewUserReoository(db *mongo.Collection) UserRepository {
	return userRepository{db}
}

func (obj userRepository) Create(user Users) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := obj.db.InsertOne(ctx, user)
	return err

}
func (obj userRepository) GetByID(getuser Users) (user Users, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	quary := handlerUser(getuser)
	err = obj.db.FindOne(ctx, quary).Decode(&user)
	return user, err

}
func (obj userRepository) GetAll() (users []Users, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := obj.db.Find(ctx, bson.M{})
	if err != nil {
		err := fmt.Errorf("empty database user")
		return users, err
	}
	if err = cursor.All(context.TODO(), &users); err != nil {
		err := fmt.Errorf("empty database user")
		return users, err
	}
	return users, err
}
