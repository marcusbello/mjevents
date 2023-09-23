package dblayer

import (
	"github.com/marcusbello/mjevents/lib/persistence"
	"github.com/marcusbello/mjevents/lib/persistence/mongolayer"
	"github.com/marcusbello/mjevents/lib/persistence/redislayer"
)

type DBTYPE string

const (
	MONGODB    DBTYPE = "mongodb"
	DOCUMENTDB DBTYPE = "documentdb"
	DYNAMODB   DBTYPE = "dynamodb"
	REDIS      DBTYPE = "redis"
)

func NewPersistenceLayer(options DBTYPE, connection string) (persistence.DatabaseHandler, error) {

	switch options {
	case MONGODB:
		return mongolayer.NewMongoDBLayer(connection)
	case REDIS:
		return redislayer.NewRedisLayer(connection)
	}

	return nil, nil
}
