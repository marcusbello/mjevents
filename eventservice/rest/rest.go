package rest

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/marcusbello/mjevents/lib/msgqueue"
	"github.com/marcusbello/mjevents/lib/persistence"
	"net/http"
)

func ServeAPI(endpoint, tlsendpoint string, databasehandler persistence.DatabaseHandler, eventEmitter msgqueue.EventEmitter) (chan error, chan error) {
	handler := NewEventHandler(databasehandler, eventEmitter)
	r := mux.NewRouter()
	eventsrouter := r.PathPrefix("/events").Subrouter()
	eventsrouter.Methods("GET").Path("/{SearchCriteria}/{search}").HandlerFunc(handler.FindEventHandler)
	eventsrouter.Methods("GET").Path("").HandlerFunc(handler.AllEventHandler)
	eventsrouter.Methods("POST").Path("").HandlerFunc(handler.NewEventHandler)
	httpErrChan := make(chan error)
	httptlsErrChan := make(chan error)

	server := handlers.CORS()(r)
	go func() {
		httptlsErrChan <- http.ListenAndServeTLS(tlsendpoint, "cert.pem", "key.pem", server)
	}()
	go func() {
		httpErrChan <- http.ListenAndServe(endpoint, server)
	}()

	return httpErrChan, httptlsErrChan
}
