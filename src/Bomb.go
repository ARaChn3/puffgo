package puffgo

import "time"

/*
Bomb: ...
Fields:
	ID [string]
		A random hex string used to identify the bomb

	Exec [string]
		A string of commands to execute when the bomb goes off

	RunOnBoot [bool]
		-_-

	Timer [*time.Timer]
		Timer for the bomb


*/
type Bomb struct {
	ID        string
	Exec      string
	RunOnBoot bool
	Timer     *time.Timer
}

func NewBomb(deadline *time.Time, e string) Bomb {
	return Bomb{}
}

func (b *Bomb) Implant() {}
