# Enums in Go

## Overview
Go doesn't have a built-in `enum` type. Instead, enums are implemented using **constants** with the `iota` identifier.

## How They Work

`iota` is a special constant that automatically increments, starting from 0 in each `const` block.

## Basic Example

```go
package main

import "fmt"

const (
    Sunday = iota    // 0
    Monday           // 1
    Tuesday          // 2
    Wednesday        // 3
    Thursday         // 4
    Friday           // 5
    Saturday         // 6
)

func main() {
    fmt.Println(Monday) // Output: 1
}
```

## String Enums

```go
const (
    Red = iota
    Green
    Blue
)

func colorName(c int) string {
    colors := []string{"Red", "Green", "Blue"}
    return colors[c]
}
```

## Typed Enums

```go
type Vehicle string

const (
    CAR Vehicle = "Car"
    TRAIN Vehicle = "Train"
    Ship Vehicle = "Ship"
    AIRPLANE Vehicle = "AirPlane"
)

func main() {
    fmt.Println(CAR) // Output: Car
}
```

## Using `String()` Method

```go
type Day int

const (
    Mon Day = iota
    Tue
    Wed
)

func (d Day) String() string {
    return []string{"Monday", "Tuesday", "Wednesday"}[d]
}

func main() {
    fmt.Println(Mon) // Output: Monday
}
```