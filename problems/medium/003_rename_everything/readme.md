# Rename Everything

Rename variables, functions and objects so that the code makes sense:

```go
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

	// optional bithday
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
```

<details>
	<summary>Tips?</summary>

1. Understand what the code does. What are each object, function and variable? What can i get reading strings and comments?
2. Based in this context, what names are coherent? Make a list.
3. Make thing intuitive. As a programmer, make things clear, think about you in the future or other programmer in your team. We should have no doubt about the code execution.

</details>

<details>
	<summary>Solution?</summary>

None here, but you can check other submissions.

If you want a feedback, submit it!

</details>
