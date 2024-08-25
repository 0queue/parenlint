# `parenlint`

Checks that arguments in call expressions are all on separate lines, and also not on the same lines as the parentheses

Good:

```go
fmt.Println("single line")

fmt.Println(
  "multiple",
  "lines",
)
```


Bad:

```go
fmt.Println("not",
  "happy")
```
