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
## 2. If, Else If, and Else

Use `else if` when you need to check multiple conditions.

Example: Checking exam marks.

```go
marks := 82

if marks >= 90 {
    fmt.Println("Excellent")
} else if marks >= 75 {
    fmt.Println("Good")
} else {
    fmt.Println("Keep practicing")
}
```

Output:

```text
Good
```

Key points:
- `if` checks the first condition.
- `else if` checks another condition if the previous one is false.
- `else` runs when none of the conditions are true.
- Go executes only the first matching block.

