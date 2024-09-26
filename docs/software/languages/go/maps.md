# Maps

Collection of Key, value pairs.

All keys and values must be of the same type

## Definition

```go
// Way 1) all keys and values are of type string
colors := map[string]string{
  "red":   "#ff0000",
  "green": "#4bf745",
  "white": "#ffffff",
}

fmt.Println(colors)
// map[green:#4bf745 red:#ff0000]

// Way 2) 
var colors map[string]string
// Way 3)
colors := make(map[string]string)
```

## Adding Properties

```go
colors["white"] = "#ffffff"
```

## Delete Properties

```go
delete(colors, "white")
```

## Iteration

```go
for color, hex := range colors {
  println(color, hex)
}

/*
red #ff0000
green #4bf745
white #ffffff\n*/
```
