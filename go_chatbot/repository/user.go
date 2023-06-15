package repository

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
)

type Users struct {
	UserID   string `json:"user_id" bson:"user_id"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

func handlerUser(user Users) bson.M {
	data := map[string]string{}
	data_json, _ := json.Marshal(user)
	json.Unmarshal(data_json, &data)
	results := bson.M{}
	for k, v := range data {
		if v != "" {
			results[k] = v
		}
	}
	return results
}
