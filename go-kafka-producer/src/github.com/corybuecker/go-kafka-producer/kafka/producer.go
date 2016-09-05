package kafka

import (
	"log"
	"sync"
	"time"

	"github.com/Shopify/sarama"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s: %s", name, elapsed)
}

func SendMessages(messages [][]byte, kafkaHost string, concurrency int) {
	var wg sync.WaitGroup

	log.Println(kafkaHost)

	defer timeTrack(time.Now(), "loading and confirming load into Kafka")

	wg.Add(concurrency)

	batchSize := len(messages)/concurrency + 1

	for i := 0; i < len(messages); i += batchSize {
		if i+batchSize > len(messages) {
			go sendBatch(messages[i:], kafkaHost, &wg)
		} else {
			go sendBatch(messages[i:i+batchSize], kafkaHost, &wg)
		}
	}

	wg.Wait()
}

func sendBatch(messages [][]byte, kafkaHost string, wg *sync.WaitGroup) {
	producer, _ := sarama.NewSyncProducer([]string{kafkaHost}, sarama.NewConfig())

	for _, message := range messages {
		rawMessage := sarama.ByteEncoder(message)
		producer.SendMessage(&sarama.ProducerMessage{
			Topic: "kafka.test",
			Value: rawMessage,
		})
	}

	wg.Done()
}
