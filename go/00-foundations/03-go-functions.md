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
## Multiple Functions in Go

A Go program can have multiple functions. Each function performs a specific task.

We define functions separately and call them whenever we need them.

### Example

```go
func greet() {
    fmt.Println("Hello!")
}

func sayBye() {
    fmt.Println("Goodbye!")
}
```

Here:
- `greet()` prints `Hello!`.
- `sayBye()` prints `Goodbye!`.

Defining these functions does not execute them. We must call them from `main()` or another function to execute them.

### Important

- `func` is used to define a function.
- A function name identifies the function.
- `()` is used when calling a function.
- `main()` is the starting point of an executable Go program.
## Functions with Parameters

### 1. What is a Parameter?

A parameter lets a function receive information when we call it.

For example, instead of creating separate functions to greet different people, we can create one function that accepts a person's name.

### 2. Example

```go
func greet(name string) {
    fmt.Println("Hello", name)
}
```

Here:
- `greet` is the function name.
- `name` is the parameter — it receives a value.
- `string` tells Go that the value must be text.

When we call:

```go
greet("Asmita")
```

Go passes `"Asmita"` into the `name` parameter. The function prints `Hello Asmita`.

### 3. Parameter vs Argument

- Parameter: `name` — the variable written in the function definition.
- Argument: `"Asmita"` — the value supplied when calling the function.

### 4. Why Use Parameters?

Parameters let us reuse one function with different inputs instead of writing a separate function for every person.

### Remember

Define the function once. Pass different values whenever you call it.