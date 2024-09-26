# Type Conversion

## To byte slice

```go
greeting := "Hi there!"

fmt.Println([]byte(greeting))
[72 105 32 116 104 101 114 101 33]
```

## Joining Slice to string

```go
import strings
strings.Join(someSlice, ",")
```

## Byte slice to String

```go
string(byteslice)
```
