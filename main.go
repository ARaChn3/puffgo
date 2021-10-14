package main

import (
	"flag"
	pfgo "puffgo/src"
	"time"
)

func main() {
	var t, e string

	flag.StringVar(&e, "e", "echo \"BOOM!!!\"", "Specify a command to be run when the bomb goes off (Required)\n{command|script} [string]")
	flag.StringVar(&t, "t", "", "Specify in what time the bomb goes off. (Required)\n{MONTH DD, YYYY at TIME (TZ) | YYYY-MONTH-DD}")

	flag.Parse()

	T := time.Parse(t, "Feb 3, 2013 at 7:54pm (PST)")

	bomb := pfgo.NewBomb(t, e)
	bomb.Implant()

}
