package services

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/spf13/viper"
	"log"
	"math/rand"
	"short_url/config"
	"short_url/iternal/db"
	"short_url/iternal/entity"
	mykafka "short_url/iternal/kafka"
	"time"
)

var queueOpenUrl *entity.Queue

func GetNewShortUrl(link *entity.Link) error {
	link.ShortURL = queueOpenUrl.Get()

	errorSave := db.SaveToLinks(link)
	if errorSave != nil {
		return errorSave
	}

	go func() {
		linkJSON, err := json.Marshal(link)
		if err != nil {
			log.Printf("Error marshaling link to JSON: %v", err)
			return
		}

		err = config.Redis.Set(config.Ctx, link.ShortURL, linkJSON, 24*time.Hour).Err()
		if err != nil {
			log.Printf("Redis error: %v", err)
		}
	}()

	return nil
}

func GetLongUrl(link entity.Link) (entity.Link, error) {
	link, err := gettingLongUrl(link)
	if err != nil {
		return entity.Link{}, err
	}

	jsonData, err := json.Marshal(link)
	if err != nil {
		log.Println("Can't map data for analysis in 'GetLongUrl' method in 'urlService'")
	}
	mykafka.PushToMyTopic(jsonData)

	return link, nil
}

func gettingLongUrl(link entity.Link) (entity.Link, error) {
	linkJSON, err := config.Redis.Get(config.Ctx, link.ShortURL).Result()
	if err != nil {
		log.Printf("Error getting link JSON from Redis: %v", err)
	}

	var retrievedLink entity.Link
	err = json.Unmarshal([]byte(linkJSON), &retrievedLink)
	if err != nil {
		log.Printf("Error unmarshaling link JSON: %v", err)
	}

	if retrievedLink.ShortURL != "" {
		return retrievedLink, nil
	}

	return getLongUrlInDb(link)
}

func getLongUrlInDb(link entity.Link) (entity.Link, error) {
	link.LongURL = db.GetLongUrlByShortUrl(link.ShortURL)
	if link.LongURL == "" {
		return entity.Link{}, errors.New("link doesn't exist")
	}
	return link, nil
}

func CheckQueueOpenUrl() {
	queueMaxLen := viper.GetInt("url.queue")
	queueOpenUrl = entity.NewQueue()
	for {
		if queueOpenUrl.Len() < queueMaxLen {

			for {
				url, err := generateUniqueURL()
				if err != nil {
					log.Println(err)
				}
				if !db.ExistShortUrl(url) && url != "" {
					queueOpenUrl.Add(url)
					break
				}
			}

		} else {
			time.Sleep(5 * time.Second)
		}
	}
}

func generateUniqueURL() (string, error) {
	randomBytes := make([]byte, viper.GetInt("url.len_new_url"))
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	uniqueURL := base64.URLEncoding.EncodeToString(randomBytes)
	return "/" + uniqueURL, nil
}
