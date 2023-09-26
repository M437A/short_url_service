package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"github.com/spf13/viper"
	"log"
	"time"
)

var myTopicChannel chan []byte

func CreateProducerTopic() {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:      []string{viper.GetString("kafka.brokerAddress")},
		Topic:        viper.GetString("kafka.topics.analysis_topic"),
		BatchSize:    10,
		BatchTimeout: 2 * time.Second,
	})
	log.Print(viper.GetString("kafka.topics.analysis_topic") + " producer run")
	myTopicChannel = make(chan []byte)

	runAbstractProducer(writer, context.Background())
}

func PushToMyTopic(message []byte) {
	myTopicChannel <- message
}

func runAbstractProducer(writer *kafka.Writer, ctx context.Context) {
	for {
		data := <-myTopicChannel
		err := writer.WriteMessages(ctx, kafka.Message{
			Value: data,
		})
		if err != nil {
			log.Println(writer.Topic + " producer data error")
		}
	}
}
