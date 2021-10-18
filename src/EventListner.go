package puffgo

import (
	"time"
)

// Function represents any arbitary function that
// returns a boolean value after performing event checks.
type Function func(...interface{}) bool

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
	TriggerFunction Function

	// [TriggerFunctionArgs] Specifies the arguments to be
	// passed into [TriggerFunction]
	TriggerFunctionArgs []interface{}

	// [TerminationChannel] acts as stopping mechanism for the mainloop.
	// If a boolean true is passed into it, the mainloop is terminated.
	TerminationChannel chan bool
}

func NewListner(interval *time.Duration, tfunc Function, tfuncArgs interface{}) (el *EventListner) {
	el.TriggerChannel = make(chan bool)
	el.TriggerFunction = tfunc
	el.TriggerFunctionArgs = tfuncArgs
	el.TerminationChannel = make(chan bool)
	if interval == nil {
		el.Interval = time.Duration(500 * time.Millisecond)
	} else {
		el.Interval = interval
	}

	return
}

// Mainloop starts an infinite loop which checks for the event's
// occurence until terminated.
func (e *EventListner) Mainloop() {
	for {

		// Pass boolean value into the TriggerChannel
		e.TriggerChannel <- e.TriggerFunction(e.TriggerFunctionArgs...)

		// Sleep for Interval
		time.Sleep(e.Interval)

		// Check if the mainloop must be terminated
		if <-e.TerminationChannel {
			break
		}
	}
}

// Terminate terminates the mainloop checking for the event.
func (e *EventListner) Terminate() {
	e.TerminationChannel <- true
}
