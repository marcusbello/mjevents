package listener

import (
	"github.com/marcusbello/mjevents/contracts"
	"github.com/marcusbello/mjevents/lib/msgqueue"
	"github.com/marcusbello/mjevents/lib/persistence"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type EventProcessor struct {
	EventListener msgqueue.EventListener
	Database      persistence.DatabaseHandler
}

func (p *EventProcessor) ProcessEvents() error {
	log.Println("Listening to events...")
	received, errors, err := p.EventListener.Listen("event.created")
	if err != nil {
		return err
	}
	for {
		select {
		case evt := <-received:
			p.handleEvent(evt)
		case err = <-errors:
			log.Printf("received error while processing msg: %s", err)
		}
	}
}

func (p *EventProcessor) handleEvent(event msgqueue.Event) {
	switch e := event.(type) {
	case *contracts.EventCreatedEvent:
		log.Printf("event %s created: %s", e.ID, e)
		objectID, _ := primitive.ObjectIDFromHex(e.ID)
		p.Database.AddEvent(persistence.Event{ID: objectID})
	case *contracts.LocationCreatedEvent:
		log.Printf("location %s created: %v", e.ID, e)
		// TODO: No persistence for locations, yet
	default:
		log.Printf("unknown event: %t", e)
	}
}
