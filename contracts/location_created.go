package contracts

import "github.com/marcusbello/mjevents/lib/persistence"

// LocationCreatedEvent is emitted whenever a location is created
type LocationCreatedEvent struct {
	ID      string             `json:"id"`
	Name    string             `json:"name"`
	Address string             `json:"address"`
	Country string             `json:"country"`
	Halls   []persistence.Hall `json:"halls"`
}

// EventName returns the event's name
func (e *LocationCreatedEvent) EventName() string {
	return "locationCreated"
}

func (e *LocationCreatedEvent) PartitionKey() string {
	return e.ID
}
