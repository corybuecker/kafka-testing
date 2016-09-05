package command_line

import "github.com/urfave/cli"

type Arguments struct {
	Concurrency      int
	KafkaHost        string
	NumberOfMessages int
	Action           func(*cli.Context) error
}

func (arguments *Arguments) GetApp() *cli.App {
	app := cli.NewApp()
	app.Name = "Kafka Producer"
	app.Usage = "write messages to Kafka broker"
	app.Action = arguments.Action
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "kafka_host",
			Value:       "localhost:9092",
			Usage:       "Kafka broker to write messages into",
			Destination: &arguments.KafkaHost,
		},
		cli.IntFlag{
			Name:        "concurrency",
			Value:       2,
			Usage:       "Number of threads to use when writing to Kafka",
			Destination: &arguments.Concurrency,
		},
		cli.IntFlag{
			Name:        "number_of_messages",
			Value:       1000,
			Usage:       "Number of messages to write to Kafka",
			Destination: &arguments.NumberOfMessages,
		},
	}

	return app
}
