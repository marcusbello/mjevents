package contracts

// EventBookedEvent is emitted whenever an event is booked
type EventBookedEvent struct {
	EventID string `json:"eventId"`
	UserID  string `json:"userId"`
}

func (c *EventBookedEvent) PartitionKey() string {
	//TODO implement me
	panic("implement me")
}

// EventName returns the event's name
func (c *EventBookedEvent) EventName() string {
	return "eventBooked"
}
