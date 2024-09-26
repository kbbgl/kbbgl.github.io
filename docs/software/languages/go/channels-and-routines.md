# Channels (`chan`) and `goroutines` (Concurrency)

Provides concurrency features.

When we execute a go program, a `goroutine` executes it line by line.

A `goroutine` is a thread managed by the Go runtime which has its own scheduler to manage the routines. The scheduler distributes routines across CPU cores.

A channel provides communication between the main routine and the child routines.

## Sending/Receiving Data to/from Channel

```go
// send the value '5' into this channel
channel <- 5
// wait for a value to be sent into the channel
// when we get one, assign back to myNumber
// this is a blocking call
myNumber <- channel

// wait for a value to be sent into the channel
// when we get one, log it out
// this is a blocking call
fmt.Println(<- channel)
```

A channel may cause the main routine to continue execution if the channel awaits data to be sent through it.

## Usage

To make a function run in concurrently:

```go
func main() {
  links := []string{
    "https://google.com",
    "https://amazon.com",
    "https://facebook.com",
    "https://stackoverflow.com",
    "https://golang.org",
  }
  
  channel := make(chan string)
  
  for _, link := range links {
    go checkLink(link, channel)
  }
  // for loop used with channels
  // will start loop when channel receives a value
  // We need to add the link as an argument to the function literal to ensure that the routine is pointing to right data
  for link := range channel {
    go func(link string) {
      time.Sleep(5 * time.Second)
      checkLink(link, channel)
    }(link)
  }
}

func checkLink(link string, channel chan string) {
  _, err := http.Get(link) // blocking call
  if err != nil {
    fmt.Println(link, "might be down")
    channel <- link
    return
  }
  
  fmt.Println(link, "is up")
  
  channel <- link
}
```
