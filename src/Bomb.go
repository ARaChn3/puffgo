package puffgo

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"os/exec"
	"time"
)

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
	Deadline *time.Time
	Exec     string
	ID       string
}

func NewBomb(deadline *time.Time, e string) Bomb {
	id, _ := randomHex(10)

	return Bomb{
		ID:       id,
		Deadline: deadline,
		Exec:     e,
	}
}

//Implant: Implants a logic bomb.
func (b *Bomb) Implant() {
	var stdout, stderr bytes.Buffer
	f := fmt.Sprintf("usr/sbin/LB%s_%s.service", b.ID, b.Deadline)
	cmd := exec.Command("touch", []string{f}...)

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil || stderr.String() != "" {
		log.Fatal(fmt.Sprintf("%s[!] E: Error while implanting bomb.%s", ColorRed, ColorReset))
		log.Printf("%s[*] Log: %s\nError: %s\n\nStderr: %s\n", ColorYellow, ColorReset, err, stderr.String())
	}

	log.Printf("%sStdout: %s%s\n", ColorCyan, stdout.String(), ColorReset)
}

//Remove: removes previously implanted bomb
func (b *Bomb) Remove() {
	f := fmt.Sprintf("usr/sbin/LB%s_%s.service", b.ID, b.Deadline)

	err := exec.Command("rm", []string{f}...).Run()
	if err != nil {
		log.Fatal(fmt.Sprintf("%s[!]E: %s%s\n", ColorRed, err, ColorReset))
	}
}

//randomHex generates a randomized hex string of length "n"
func randomHex(n int) (string, error) {
	bytes := make([]byte, n)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}
