package main

import (
	"flag"
	"fmt"
	msgqueue_amqp "github.com/marcusbello/mjevents/lib/msgqueue/amqp"
	"github.com/streadway/amqp"
	"log"

	"github.com/marcusbello/mjevents/eventservice/rest"
	"github.com/marcusbello/mjevents/lib/configuration"
	"github.com/marcusbello/mjevents/lib/persistence/dblayer"
)

func main() {
	confPath := flag.String("conf", `../lib/configuration/config.json`, "flag to set the path to the configuration json file")
	flag.Parse()

	//extract configuration
	config, _ := configuration.ExtractConfiguration(*confPath)
	// AMQP
	conn, err := amqp.Dial(config.AMQPMessageBroker)
	if err != nil {
		panic(err)
	}
	emitter, err := msgqueue_amqp.NewAMQPEventEmitter(conn)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connecting to database")
	dbhandler, _ := dblayer.NewPersistenceLayer(config.Databasetype, config.DBConnection)

	httpErrChan, httptlsErrChan := rest.ServeAPI(config.RestfulEndpoint, config.RestfulTLSEndPint, dbhandler, emitter)
	select {
	case err := <-httpErrChan:
		log.Fatal("HTTP Error: ", err)
	case err := <-httptlsErrChan:
		log.Fatal("HTTPS Error: ", err)
	}
}
