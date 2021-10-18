package puffgo

import (
	"crypto/rand"
	"encoding/hex"
	"sync"
)

//LogicBomb: type implementing a logic-bomb for UNIX/Linux based systems.
type LogicBomb struct {

	// Listener represents an event-listner defined in this library.
	Listener *EventListener

	// ExecutionFunction specifies the function
	// which will be executed when the bomb is triggered and it goes off
	ExecutionFunction func()

	// BombId specifies a random hex string used to identify the bomb.
	BombID string

	// PID specifies the process-id of the persistent logic-bomb program running.
	// It holds nil until the mainloop of the event-listner is started.
	PID *int
}

// NewBomb returns an instance of LogicBomb, which can be
// implanted by Implant(), to be triggered when conditions
// for it to do so are met.
func NewBomb(listener EventListener, execFunc func()) *LogicBomb {
	var lb LogicBomb

	lb.Listener = &listener
	lb.BombID, _ = randomHex(10)
	lb.ExecutionFunction = execFunc
	lb.PID = lb.Listener.PID

	return &lb
}

func (lb *LogicBomb) Arm() {
	var wg sync.WaitGroup

	// Run listner's mainloop
	go func() {
		defer wg.Done()
		lb.Listener.Mainloop()
	}()

	// Check for trigger...
	go func() {
		defer wg.Done()
		for {
			if isTriggered := <-lb.Listener.TriggerChannel; isTriggered {
				lb.ExecutionFunction()
				lb.Listener.Terminate()
				break
			}
		}
	}()

	wg.Add(2)
	wg.Wait()
}

func (lb *LogicBomb) Disarm() {
	lb.Listener.Terminate()
}

// Used for generating the BombID on creation of LogicBombs
func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
