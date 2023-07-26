package couchbaselayer

import "github.com/marcusbello/mjevents/lib/persistence"

type CouchBaseLayer struct {
}

func (c *CouchBaseLayer) AddEvent(event persistence.Event) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CouchBaseLayer) FindEvent(bytes []byte) (persistence.Event, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CouchBaseLayer) FindEventByName(s string) (persistence.Event, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CouchBaseLayer) FindAllAvailableEvents() ([]persistence.Event, error) {
	//TODO implement me
	panic("implement me")
}

func NewCouchbaseLayer(connection string) (*CouchBaseLayer, error) {

	return &CouchBaseLayer{}, nil
}
