package main

import (
	"fmt"

	puffgo "github.com/NovusEdge/puffgo/src"
)

func main() {
	tfunc := func() bool { return true }
	execFunc := func() { fmt.Println("BOOM!!!") }

	el := puffgo.NewListener(nil, tfunc)
	b := puffgo.NewBomb(*el, execFunc)

	b.Arm()
}
