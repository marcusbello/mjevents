package configuration

import (
	"encoding/json"
	"fmt"
	"github.com/marcusbello/mjevents/src/lib/persistence/dblayer"
	"os"
)

var (
	DBTypeDefault       = dblayer.DBTYPE("mongodb")
	DBConnectionDefault = "mongodb://127.0.0.1"
	RestfulEPDefault    = "localhost:8181"
	RestfulTLSEPDefault = "localhost:9191"
)

type ServiceConfig struct {
	Databasetype      dblayer.DBTYPE `json:"databasetype"`
	DBConnection      string         `json:"dbconnection"`
	RestfulEndpoint   string         `json:"restfulapi_endpoint"`
	RestfulTLSEndPint string         `json:"restfulapi-tlsendpoint"`
}

func ExtractConfiguration(filename string) (ServiceConfig, error) {
	conf := ServiceConfig{
		DBTypeDefault,
		DBConnectionDefault,
		RestfulEPDefault,
		RestfulTLSEPDefault,
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Configuration file not found. Continuing with default values.")
		return conf, err
	}
	err = json.NewDecoder(file).Decode(&conf)
	return conf, err
}
