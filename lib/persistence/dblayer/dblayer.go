package dblayer

import (
	"github.com/marcusbello/mjevents/lib/persistence"
	"github.com/marcusbello/mjevents/lib/persistence/couchbaselayer"
	"github.com/marcusbello/mjevents/lib/persistence/mongolayer"
)

type DBTYPE string

const (
	MONGODB    DBTYPE = "mongodb"
	DOCUMENTDB DBTYPE = "documentdb"
	DYNAMODB   DBTYPE = "dynamodb"
	COUCHBASE  DBTYPE = "couchbaselayer"
)

func NewPersistenceLayer(options DBTYPE, connection string) (persistence.DatabaseHandler, error) {

	switch options {
	case MONGODB:
		return mongolayer.NewMongoDBLayer(connection)
	case COUCHBASE:
		return couchbaselayer.NewCouchbaseLayer(connection)
	}

	return nil, nil
}
