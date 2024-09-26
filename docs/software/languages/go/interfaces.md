# Interfaces

`englishBot` and `spanishBot` both implement the `getGreeting` method so according to `bot interface`, they are both of considered `interface`s.

```go
package main
import "fmt"

// Declare interface 
type bot interface {
  getGreeting() string
}

type spanishBot struct{}
type englishBot struct{}

func main() {
  eb := englishBot{}
  sb := spanishBot{}
  
  printGreeting(eb)
  printGreeting(sb)
}


// Declare that function uses interface
func printGreeting(b bot) {
  fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
  return "Hi there!"
}

func (spanishBot) getGreeting() string {
  return "Hola!"
}

```

- Interfaces are not generic types.
- Interfaces are implicit (no need to `implements` a certain type to the interface type)
- Are a contract to help us manage types.

### `Reader` Interface and `io.Copy`

Very useful interface is the `Reader` which takes in any source of information (i.e. HTTP response body, text/image file on disk, user input) and outputs a byteslice (`[]byte`).

To read data from an HTTP source:

```go
resp, err := http.Get("https://httpbin.org/anything")
if err != nil {
  print(err)
}

// resp.Body implements the `Reader` interface
// os.Stdout implments the `Wrtier` interface
io.Copy(os.Stdout, resp.Body)
```
