package kafka

func KafkaRun() {
	go CreateProducerTopic()
	go CreateConsumerTopic()
}
