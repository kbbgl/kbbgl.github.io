# Functions

To return multiple values from a function

```go
// Returns two decks

func deal(cards deck, handSize int) (deck, deck) {
  return d[:handSize], d[handSize:]
}
```

To capture both function outputs:

```go
deck1, deck2 := deal(cards, 5)
```

A lot of functions return an `error` type:

```go
byteslice, err := ioutil.ReadFile(filename)
if err != nil
```

## Function Literal (unnamed, anonymous, lambda)

```go
func() {
  // some code...
}()
