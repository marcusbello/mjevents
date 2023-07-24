package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/marcusbello/mjevents/src/eventsservice/rest"
	"github.com/marcusbello/mjevents/src/lib/configuration"
	"github.com/marcusbello/mjevents/src/lib/persistence/dblayer"
)

func main() {
	confPath := flag.String("conf", `../lib/configuration/config.json`, "flag to set the path to the configuration json file")
	flag.Parse()

	//extract configuration
	config, _ := configuration.ExtractConfiguration(*confPath)
	fmt.Println("Connecting to database")
	dbhandler, _ := dblayer.NewPersistenceLayer(config.Databasetype,
		config.DBConnection)

	httpErrChan, httptlsErrChan := rest.ServeAPI(config.RestfulEndpoint, config.RestfulTLSEndPint, dbhandler)
	select {
	case err := <-httpErrChan:
		log.Fatal("HTTP Error: ", err)
	case err := <-httptlsErrChan:
		log.Fatal("HTTPS Error: ", err)
	}
}
