package analitic

import (
	"encoding/json"
	"fmt"
	"log"
	"short_url/iternal/db"
	"short_url/iternal/entity"
	"time"
)

func SaveLinksAnalysis(message []byte) {
	var link entity.Link

	err := json.Unmarshal(message, &link)
	if err != nil {
		log.Println("Error in mapping data", err)
	}

	userData := entity.AnalysisUserData{UserId: link.UserId,
		UserName: db.GetUserNameById(link.Id),
		ShortURL: link.ShortURL,
		LongURL:  link.LongURL,
		Time:     time.Now()}

	linkData := entity.AnalysisLinkData{
		ShortURL: link.ShortURL,
		LongURL:  link.LongURL,
		Time:     time.Now()}
	go saveUserToAnalytic(userData)
	go saveLinkToAnalytic(linkData)
}

// for using this api, you need to set up your account
// home page this api 'https://app.amplitude.com/analytics/testamplituda/home'
func saveUserToAnalytic(data entity.AnalysisUserData) {
	fmt.Println("Save user to amplitude")
	log.Println(data)
	//example code
	/*
		config.AnalyticClient.Track(amplitude.Event{
			UserID:    strconv.FormatUint(uint64(data.UserId), 10),
			EventType: "User Data",
			EventProperties: map[string]interface{}{
				"UserName": data.UserName,
				"Time":     data.Time,
			},
		})

		config.AnalyticClient.Flush()

	*/
}

func saveLinkToAnalytic(data entity.AnalysisLinkData) {
	fmt.Println("Save link to amplitude")
	log.Println(data)
	//example code
	/*
		config.AnalyticClient.Track(amplitude.Event{
			DeviceID:  strconv.FormatUint(uint64(data.LinkId), 10),
			EventType: "Link Data",
			EventProperties: map[string]interface{}{
				"ShortURL": data.ShortURL,
				"LongURL":  data.LongURL,
				"Time":     data.Time,
			},
		})

		config.AnalyticClient.Flush()

	*/
}
