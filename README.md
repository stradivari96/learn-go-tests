# learn-go-tests

https://quii.gitbook.io/learn-go-with-tests/

```
go mod init hello
go mod tidy
go test ./2.integers -v
go test ./3.iteration -bench="."
```

## Notes

- end tests with `_test.go`
- test function starts with `Test`: `func TestHello(t *testing.T)`
- Godoc examples, will run as tests https://go.dev/blog/examples
- `func BenchmarkRepeat(b *testing.B)`
- no polymorphism in Go
- Go implicitly deferences pointers (w.attribute is the same as (\*w).attribute)
