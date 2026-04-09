# Generics in Go

## Overview
Generics allow you to write reusable code that works with multiple types while maintaining type safety. Go 1.18 introduced support for type parameters.

## Basic Syntax

### Generic Functions
```go
func Print[T any](value T) {
    fmt.Println(value)
}
```

### Generic Types
```go
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(value T) {
    s.items = append(s.items, value)
}

func (s *Stack[T]) Pop() T {
    value := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return value
}
```

## Constraints

### Using Built-in Constraints
```go
import "golang.org/x/exp/constraints"

func Sum[T constraints.Integer](nums []T) T {
    var sum T
    for _, num := range nums {
        sum += num
    }
    return sum
}
```

### Custom Constraints
```go
type Comparable interface {
    Compare(other Comparable) int
}

func Max[T Comparable](a, b T) T {
    if a.Compare(b) > 0 {
        return a
    }
    return b
}

//Only allow int and string types
func PrintNumbers[T int | string](items []T) []T {
	fmt.Println(items)
	return items
}
```

## Key Points
- Use square brackets `[T any]` for type parameters
- Type parameters appear before regular parameters
- Constraints limit what types can be used
- `any` is equivalent to `interface{}`