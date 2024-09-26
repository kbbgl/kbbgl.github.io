# Variables and Types

Go is a statically defined language (like Java, C++)

```go
// 'string' here tells the go compiler that this variable will only be of string type 
var cards string = "Ace of Spades"
```

Other types:

```go
bool //(true|false)
string
int
float64
```

A more efficient way to declare a variable and let go figure out the type of data:

```go
card := "Ace of Spades"
```

Can also declare a variable and then define it

```go
var deckSize int
deckSize = 52
```

To define a method return of type string:

```go
func newCard() string {
  return "Five of Diamonds"
}

```

### Adding Variable Values to Formatted String (`%d`, `%v`)

```go
t.Errorf("Expected deck length is 52, but got %d", len(d))
```
