package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/corybuecker/go-kafka-producer/command_line"
	"github.com/corybuecker/go-kafka-producer/kafka"
	"github.com/corybuecker/go-kafka-producer/messages"
	"github.com/urfave/cli"
)

func main() {
	rand.Seed(time.Now().Unix())

	arguments := &command_line.Arguments{}

	arguments.Action = func(c *cli.Context) error {
		messages := messages.GenerateRawMessages(arguments.NumberOfMessages)
		kafka.SendMessages(messages, arguments.KafkaHost, arguments.Concurrency)
		return nil
	}

	arguments.GetApp().Run(os.Args)
}
