# `parenlint`

Checks that arguments to function calls are all on the same line, or on multiple lines.

Good:

```go
fmt.Println("single line")

fmt.Println(
  "multiple",
  "lines",
)

t.Run("test", func(t *testing.T) {
  t.Log("hooray!")
})
```


Bad:

```go
fmt.Println("not",
  "happy")

fmt.Println(
  "also", "frowned upon"
)
```
