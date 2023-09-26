package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"github.com/spf13/viper"
	"log"
	"short_url/iternal/services/analitic"
	"time"
)

func CreateConsumerTopic() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{viper.GetString("kafka.brokerAddress")},
		Topic:    viper.GetString("kafka.topics.analysis_topic"),
		GroupID:  "my-group",
		MinBytes: 1,
		MaxBytes: 1e6,
		MaxWait:  1 * time.Second,
	})
	log.Print(viper.GetString("kafka.topics.analysis_topic") + " consumer run")
	runAbstractConsumer(reader, context.Background())
}

func runAbstractConsumer(reader *kafka.Reader, ctx context.Context) {
	for {
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			log.Println("receiver has trouble with data")
		} else {
			analitic.SaveLinksAnalysis(msg.Value)
		}
	}
}
