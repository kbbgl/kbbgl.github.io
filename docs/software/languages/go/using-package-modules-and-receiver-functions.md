# Package Modules and Receiver Functions

If we have 2 files:

```bash
> ls .
main.go
deck.go
```

We can use the logic in `deck.go` in `main.go` as so:

```go
// deck.go
package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// function using a receiver 'deck'
func (d deck) print() {
  for i, card := range d {
    fmt.Println(i, card)
  }
}

```

```go
// main.go

package main

import "fmt"

func main() {
  cards := deck{newCard(), newCard(), "Ace of Diamonds"}
  
  // using print method from deck module
  cards.print()
}

func newCard() string {
  return "Five of Diamonds"
}
```
