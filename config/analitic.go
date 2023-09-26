package config

import (
	"github.com/amplitude/analytics-go/amplitude"
	"github.com/spf13/viper"
)

var AnalyticClient amplitude.Client

func RunAnalytic() amplitude.Client {
	analytics := amplitude.NewClient(
		amplitude.NewConfig(viper.GetString("analytic.api")),
	)

	AnalyticClient = analytics

	return analytics
}
