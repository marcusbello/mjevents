package main

import (
	"flag"
	"fmt"

	"github.com/IBM/sarama"

	"github.com/marcusbello/mjevents/eventservice/rest"
	"github.com/marcusbello/mjevents/lib/configuration"
	"github.com/marcusbello/mjevents/lib/msgqueue"
	msgqueue_amqp "github.com/marcusbello/mjevents/lib/msgqueue/amqp"
	"github.com/marcusbello/mjevents/lib/msgqueue/kafka"
	"github.com/marcusbello/mjevents/lib/persistence/dblayer"
	"github.com/streadway/amqp"
)

func main() {
	var eventEmitter msgqueue.EventEmitter

	confPath := flag.String("conf", `../lib/configuration/config.json`, "flag to set the path to the configuration json file")
	flag.Parse()
	//extract configuration
	config, _ := configuration.ExtractConfiguration(*confPath)

	switch config.MessageBrokerType {
	case "amqp":
		conn, err := amqp.Dial(config.AMQPMessageBroker)
		if err != nil {
			panic(err)
		}

		eventEmitter, err = msgqueue_amqp.NewAMQPEventEmitter(conn, "events")
		if err != nil {
			panic(err)
		}
	case "kafka":
		conf := sarama.NewConfig()
		conf.Producer.Return.Successes = true
		conn, err := sarama.NewClient(config.KafkaMessageBrokers, conf)
		if err != nil {
			panic(err)
		}

		eventEmitter, err = kafka.NewKafkaEventEmitter(conn)
		if err != nil {
			panic(err)
		}
	default:
		panic("Bad message broker type: " + config.MessageBrokerType)
	}

	fmt.Println("Connecting to database")
	dbHandler, _ := dblayer.NewPersistenceLayer(config.Databasetype, config.DBConnection)

	fmt.Println("Serving API")
	//RESTful API start
	err := rest.ServeAPI(config.RestfulEndpoint, dbHandler, eventEmitter)
	if err != nil {
		panic(err)
	}
}
