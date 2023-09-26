package db

import (
	"context"
	"log"
	"short_url/config"
	"short_url/iternal/entity"
)

func SaveToUsers(user *entity.User) error {
	request := `
        INSERT INTO users (user_name, email)
        VALUES ($1, $2, $3)
    `
	_, err := config.DB.Exec(context.Background(), request, user.Name, user.Email)
	if err != nil {
		log.Print("Error with save to db: " + err.Error())
		return err
	}
	return nil
}
