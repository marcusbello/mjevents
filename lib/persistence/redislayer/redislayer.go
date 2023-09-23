package redislayer

import (
	"github.com/marcusbello/mjevents/lib/persistence"
	"github.com/redis/go-redis/v9"
)

type RedisLayer struct {
	client *redis.Client
}

func (r *RedisLayer) AddUser(user persistence.User) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (r *RedisLayer) AddEvent(event persistence.Event) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (r *RedisLayer) AddBookingForUser(bytes []byte, booking persistence.Booking) error {
	//TODO implement me
	panic("implement me")
}

func (r *RedisLayer) AddLocation(location persistence.Location) (persistence.Location, error) {
	//TODO implement me
	panic("implement me")
}

func (r *RedisLayer) FindUser(s string, s2 string) (persistence.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *RedisLayer) FindBookingsForUser(bytes []byte) ([]persistence.Booking, error) {
	//TODO implement me
	panic("implement me")
}

func (r *RedisLayer) FindEvent(bytes []byte) (persistence.Event, error) {
	//TODO implement me
	panic("implement me")
}

func (r *RedisLayer) FindEventByName(s string) (persistence.Event, error) {
	//TODO implement me
	panic("implement me")
}

func (r *RedisLayer) FindAllAvailableEvents() ([]persistence.Event, error) {
	//TODO implement me
	panic("implement me")
}

func (r *RedisLayer) FindLocation(s string) (persistence.Location, error) {
	//TODO implement me
	panic("implement me")
}

func (r *RedisLayer) FindAllLocations() ([]persistence.Location, error) {
	//TODO implement me
	panic("implement me")
}

func NewRedisLayer(connection string) (*RedisLayer, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     connection,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &RedisLayer{client: rdb}, nil
}
