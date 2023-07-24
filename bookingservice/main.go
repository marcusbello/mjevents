package main

import (
	"flag"
	"github.com/marcusbello/mjevents/bookingservice/listener"
	"github.com/marcusbello/mjevents/lib/configuration"
	msgqueue_amqp "github.com/marcusbello/mjevents/lib/msgqueue/amqp"
	"github.com/marcusbello/mjevents/lib/persistence/dblayer"
	"github.com/streadway/amqp"
)

func main() {
	confPath := flag.String("config", "./lib/configuration/config.json", "path to config file")
	flag.Parse()
	config, err := configuration.ExtractConfiguration(*confPath)
	dblayer, err := dblayer.NewPersistenceLayer(config.Databasetype, config.DBConnection)
	if err != nil {
		panic(err)
	}
	conn, err := amqp.Dial(config.AMQPMessageBroker)
	if err != nil {
		panic(err)
	}
	eventListener, err := msgqueue_amqp.NewAMQPEventListener(conn, "events")
	if err != nil {
		panic(err)
	}
	processor := &listener.EventProcessor{eventListener, dblayer}
	go processor.ProcessEvents()
}
