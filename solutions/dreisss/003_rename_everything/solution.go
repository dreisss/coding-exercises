package main

import (
	"fmt"
	"time"
)

type L string

type P struct {
	n string
	g string
	l []L

	// optional
	b time.Time
}

func newP(n string, g string, l []L) P {
	return P{
		n: n,
		g: g,
		l: l,
	}
}

func (p P) withB(b time.Time) P {
	p.b = b
	return p
}

func main() {
	ps := []P{
		newP("Davi Picanço dos Reis", "dreisss", []L{"Rust", "Go"}).
			withB(time.Date(2003, time.May, 26, 0, 0, 0, 0, time.Local)),
		newP("Fábio Akita", "akitaonrails", []L{"Ruby", "Crystal"}),
		newP("Lucas Montano", "lucasmontano", []L{"Kotlin", "Java"}),
	}

	for i := range ps {
		person := ps[i]
		fmt.Printf("%s has username '%s' on github and likes this programming languages: %s.\n", person.n, person.g, person.l)
	}
}
