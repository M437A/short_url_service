package db

import (
	"context"
	"log"
	"short_url/config"
	"short_url/iternal/entity"
)

func GetLinkByLongUrl(longUrl string) string {
	var link entity.Link
	request := "SELECT * FROM links WHERE long_url=$1"
	err := config.DB.QueryRow(context.Background(), request, longUrl).Scan(&link.Id, &link.UserId, &link.ShortURL, &link.LongURL)
	if err != nil {
		return ""
	}
	return link.ShortURL
}

func GetLongUrlByShortUrl(shortUrl string) string {
	var longUrl string
	request := "SELECT long_url FROM links WHERE short_url=$1"
	err := config.DB.QueryRow(context.Background(), request, shortUrl).Scan(&longUrl) // Обратите внимание на &longUrl
	if err != nil {
		return ""
	}
	return longUrl
}

func SaveToLinks(link *entity.Link) error {
	request := `
        INSERT INTO Links (user_id, short_url, long_url)
        VALUES ($1, $2, $3)
    `
	_, err := config.DB.Exec(context.Background(), request, link.UserId, link.ShortURL, link.LongURL)
	if err != nil {
		log.Print("Error with save to db: " + err.Error())
		return err
	}
	return nil
}

func ExistShortUrl(shortUrl string) bool {
	request := "SELECT EXISTS (SELECT 1 FROM links WHERE short_url = $1)"
	var exists bool
	err := config.DB.QueryRow(context.Background(), request, shortUrl).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}

func GetUserNameById(id uint) string {
	var name string
	request := "SELECT user_name FROM links WHERE id=$1"
	err := config.DB.QueryRow(context.Background(), request, id).Scan(name)
	if err != nil {
		return ""
	}
	return name
}
