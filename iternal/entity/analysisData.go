package entity

import "time"

type AnalysisUserData struct {
	UserId   uint
	UserName string
	ShortURL string
	LongURL  string
	Time     time.Time
}

type AnalysisLinkData struct {
	LinkId   uint
	ShortURL string
	LongURL  string
	Time     time.Time
}
