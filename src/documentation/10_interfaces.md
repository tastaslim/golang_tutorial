# Interfaces in Go

## What is an Interface?

An interface is a type that defines a set of method signatures. It specifies what methods a type must implement, without defining how those methods work. Any type that implements all the methods of an interface automatically satisfies that interface.

## Basic Syntax

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

## How Interfaces Work in Go

### 1. **Implicit Implementation**
Types don't need to explicitly declare they implement an interface. If a type has all required methods, it automatically satisfies the interface.

```go
type Dog struct{}

func (d Dog) Speak() string {
    return "Woof!"
}

type Animal interface {
    Speak() string
}

var a Animal = Dog{} // Valid - Dog implements Animal
```

### 2. **Empty Interface**
An empty interface accepts any type:

```go
var x interface{} = "hello"
```

## Significance in Go

| Aspect | Benefit |
|--------|---------|
| **Polymorphism** | Write generic code that works with multiple types |
| **Loose Coupling** | Reduce dependencies between packages |
| **Composition** | Build flexible, reusable components |
| **Testing** | Easy to create mock implementations |
| **Abstraction** | Hide implementation details |

## Practical Example

```go
package main

type Writer interface {
    Write(data string) error
}

type FileWriter struct{}

func (fw FileWriter) Write(data string) error {
    // Write to file
    return nil
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data string) error {
    // Write to console
    return nil
}

func SaveData(w Writer, data string) {
    w.Write(data) // Works with any type implementing Writer
}
```
