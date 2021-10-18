package puffgo

import (
	"crypto/rand"
	"encoding/hex"
)

//LogicBomb: type implementing a logic-bomb for UNIX/Linux based systems.
type LogicBomb struct {

	// Listner represents an event-listner defined in this library.
	Listner *EventListner

	// ExecutionFunction specifies the function
	// which will be executed when the bomb is triggered and it goes off
	ExecutionFunction func(...interface{})

	// ExecutionFunctionArgs specifies the arguments to be passed into ExecutionFunction
	ExecutionFunctionArgs interface{}

	// BombId represents a random hex string used to identify the bomb.
	BombID string
}

// NewBomb returns an instance of LogicBomb, which can be
// implanted by Implant(), to be triggered when conditions
// for it to do so are met.
func NewBomb(listner Listner, execFunc func(...interface{}), execFuncArgs interface{}) (lb *LogicBomb) {
	lb.Listner = listner
	lb.BombID, _ = randomHex(10)
	lb.ExecutionFunction = execFunc
	lb.ExecutionFunctionArgs = execFuncArgs

	return
}

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
