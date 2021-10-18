package puffgo

import (
	"os"
	"time"
)

// EventListener ...
// TODO: Write some docs for this type
type EventListener struct {

	// TriggerChannel is to detect if the event has occured.
	// When it does, a boolean true is passed into it.
	TriggerChannel chan bool

	// [Interval] specifies the delay between each check of for the event.
	Interval *time.Duration

	// [TriggerFunction] specifies the function which will be
	// used to check when the event has occured.
	//
	// The function will return a boolean, and will be executed
	// on intervals of [Interval] (type: *time.Duration).
	// If Interval is nil, it performs checks every 500ms.
	//
	// When the event occurs, true is passed into [TriggerChannel],
	// which can be used as a detection mechanism.
	TriggerFunction func() bool

	// [TerminationChannel] acts as stopping mechanism for the mainloop.
	// If a boolean true is passed into it, the mainloop is terminated.
	TerminationChannel chan bool

	// PID specifies the process-id of the persistent mainloop.
	// It holds nil until the mainloop of the event-listner is started.
	PID *int
}

func NewListener(interval *time.Duration, tfunc func() bool) *EventListener {
	var el EventListener
	pid := os.Getpid()

	el.PID = &pid
	el.TriggerChannel = make(chan bool)
	el.TriggerFunction = tfunc
	el.TerminationChannel = make(chan bool)
	if interval == nil {
		td := time.Duration(500 * time.Millisecond)
		el.Interval = &td
	} else {
		el.Interval = interval
	}

	return &el
}

// Mainloop starts an infinite loop which checks for the event's
// occurence until terminated.
func (e *EventListener) Mainloop() {
	*(e.PID) = os.Getpid()

	for {

		// Pass boolean value into the TriggerChannel
		e.TriggerChannel <- e.TriggerFunction()

		// Sleep for Interval
		time.Sleep(*(e.Interval))

		// Check if the mainloop must be terminated
		if <-e.TerminationChannel {
			break
		}
	}
}

// Terminate terminates the mainloop checking for the event.
func (e *EventListener) Terminate() {
	e.TerminationChannel <- true
}
