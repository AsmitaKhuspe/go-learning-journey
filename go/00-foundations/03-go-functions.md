# Functions in Go

## What is a Function?

A function is a reusable block of code that performs a specific task.

Instead of writing the same code again and again, we can put it inside a function and call that function whenever we need it.

### Basic Structure

```go
func functionName() {
    // code
}
## Why Do We Use Functions?

Functions help us:

1. **Reuse code** — write code once and use it many times.
2. **Organize code** — divide a large program into smaller tasks.
3. **Make code easier to read** — each function can have one clear responsibility.
4. **Make code easier to test** — individual functions can be tested separately.

### Real-World Example

Think of a restaurant.

Instead of the chef doing everything in one huge process, different tasks can be separated:

- `takeOrder()`
- `prepareFood()`
- `serveFood()`
- `takePayment()`

Each function performs one specific task.

Go programs can be organized in the same way.