# Run Tests in `go`

1) Create file ending with `_test.go`.
2) Run:

```bash
go test
```

Test functions need to have the following signature:

```go
func TestNewDeck(t *testing.T) {
  // testing logic...
}
```
