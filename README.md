# learn-go-tests

https://quii.gitbook.io/learn-go-with-tests/

```
go mod init hello
go mod tidy
go test ./integers -v
```

## Tests

- end with `_test.go`
- function name starts with `Test`
- function example: `func TestHello(t *testing.T)`
- Godoc examples, will run as tests https://go.dev/blog/examples
