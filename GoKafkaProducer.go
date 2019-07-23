package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"os"
	"time"
)

const (
	kafkaConn = "localhost:9092"
	topic     = "users"
)

func main() {

	// create producer

	producer, err := initProducer()

	if err != nil {
		fmt.Println("Error producer: ", err.Error())
		os.Exit(1)
	}

	for i := 0; i < 1000; i++ {
		timeString := time.Now().String()
		msg := ("myLifeMytest -" + timeString)
		fmt.Println("Sent to Kafka %s ", msg)

		//Some go Routine just to see a small paralelism
		go publish(msg, producer)
	}

}

func initProducer() (sarama.SyncProducer, error) {
	// setup sarama log to stdout
	sarama.Logger = log.New(os.Stdout, "", log.Ltime)

	// producer config
	config := sarama.NewConfig()
	config.Producer.Retry.Max = 5
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	// async producer
	//prd, err := sarama.NewAsyncProducer([]string{kafkaConn}, config)

	// sync producer
	prd, err := sarama.NewSyncProducer([]string{kafkaConn}, config)

	return prd, err
}

func publish(message string, producer sarama.SyncProducer) {

	// publish sync
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}
	p, o, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Println("Error publish: ", err.Error())
	}

	// publish async
	//producer.Input() <- &sarama.ProducerMessage{

	fmt.Println("Partition: ", p)
	fmt.Println("Offset: ", o)

}
