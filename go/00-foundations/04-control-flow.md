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

## 3. Switch Statement

A `switch` checks one value against multiple possible cases.

Example: Checking a day number.

```go
day := 2

switch day {
case 1:
    fmt.Println("Monday")
case 2:
    fmt.Println("Tuesday")
case 3:
    fmt.Println("Wednesday")
default:
    fmt.Println("Invalid day")
}
```

Output:

```text
Tuesday
```

Key points:
- `switch` checks the value.
- `case` defines a possible match.
- `default` runs if no case matches.
- Go automatically stops after the matching case, so `break` is usually unnecessary.
