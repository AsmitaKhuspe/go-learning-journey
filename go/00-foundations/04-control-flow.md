# Control Flow in Go

## 1. If and Else

Control flow decides which instructions a program executes.

Analogy: If it is raining, take an umbrella. Otherwise, leave it at home.

### `if`

Runs code when a condition is true.

### `else`

Runs code when the condition is false.

### Example

```go
package main

import "fmt"

func main() {
    age := 20

    if age >= 18 {
        fmt.Println("Adult")
    } else {
        fmt.Println("Minor")
    }
}
```

Output:

```text
Adult
```

Remember:
- `if` checks a condition.
- `else` handles the alternative.
- Conditions evaluate to `true` or `false`.
- Go uses `{}` to enclose the code blocks.
