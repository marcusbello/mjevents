package mongolayer

import (
	"context"
	"github.com/marcusbello/mjevents/lib/persistence"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	DB        = "mjevents"
	USERS     = "users"
	EVENTS    = "events"
	LOCATIONS = "locations"
)

type MongoDBLayer struct {
	session *mgo.Session
	client  *mongo.Client
}

func (mgoLayer *MongoDBLayer) AddUser(user persistence.User) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (mgoLayer *MongoDBLayer) AddBookingForUser(bytes []byte, booking persistence.Booking) error {
	//TODO implement me
	panic("implement me")
}

func (mgoLayer *MongoDBLayer) AddLocation(location persistence.Location) (persistence.Location, error) {
	//TODO implement me
	panic("implement me")
}

func (mgoLayer *MongoDBLayer) FindUser(s string, s2 string) (persistence.User, error) {
	//TODO implement me
	panic("implement me")
}

func (mgoLayer *MongoDBLayer) FindBookingsForUser(bytes []byte) ([]persistence.Booking, error) {
	//TODO implement me
	panic("implement me")
}

func (mgoLayer *MongoDBLayer) FindLocation(s string) (persistence.Location, error) {
	//TODO implement me
	panic("implement me")
}

func (mgoLayer *MongoDBLayer) FindAllLocations() ([]persistence.Location, error) {
	//TODO implement me
	panic("implement me")
}

func (mgoLayer *MongoDBLayer) AddEvent(e persistence.Event) ([]byte, error) {
	ctx := context.TODO()
	//client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1"))
	//if err != nil {
	//	return nil, err
	//}
	//defer client.Disconnect(ctx)

	c := mgoLayer.client.Database(DB)
	collection := c.Collection(EVENTS)
	if e.ID.IsZero() {
		e.ID = primitive.NewObjectID()
	}
	insertOne, err := collection.InsertOne(ctx, e)
	if err != nil {
		return nil, err
	}
	insertedID := insertOne.InsertedID

	//if !e.ID.Valid() {
	//	e.ID = bson.NewObjectId()
	//}
	//let's assume the method below checks if the ID is valid for the location object of the event
	if e.Location.ID.IsZero() {
		e.Location.ID = primitive.NewObjectID()
	}
	//if !e.Location.ID.Valid() {
	//	e.Location.ID = bson.NewObjectId()
	//}
	return []byte(insertedID.(primitive.ObjectID).String()), nil
}

func (mgoLayer *MongoDBLayer) FindEvent(id []byte) (persistence.Event, error) {
	s := mgoLayer.getFreshSession()
	defer s.Close()
	e := persistence.Event{}
	err := s.DB(DB).C(EVENTS).FindId(bson.ObjectId(id)).One(&e)
	return e, err
}

func (mgoLayer *MongoDBLayer) FindEventByName(name string) (persistence.Event,
	error) {
	s := mgoLayer.getFreshSession()
	defer s.Close()
	e := persistence.Event{}
	err := s.DB(DB).C(EVENTS).Find(bson.M{"name": name}).One(&e)
	return e, err
}

func (mgoLayer *MongoDBLayer) FindAllAvailableEvents() ([]persistence.Event,
	error) {
	s := mgoLayer.getFreshSession()
	defer s.Close()
	var events []persistence.Event
	err := s.DB(DB).C(EVENTS).Find(nil).All(&events)
	return events, err
}

func NewMongoDBLayer(connection string) (*MongoDBLayer, error) {
	s, err := mgo.Dial(connection)
	if err != nil {
		return nil, err
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connection))
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(ctx)

	return &MongoDBLayer{
		session: s,
		client:  client,
	}, err
}

func (mgoLayer *MongoDBLayer) getFreshSession() *mgo.Session {
	return mgoLayer.session.Copy()
}

func (mgoLayer *MongoDBLayer) getClientDatabase(database string) *mongo.Database {
	return mgoLayer.client.Database(database)
}
