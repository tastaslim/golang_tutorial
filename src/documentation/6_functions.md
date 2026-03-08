
# Go Functions

## Function Basics
Functions are reusable blocks of code. Basic syntax:
```go
func functionName(param1 type1, param2 type2) returnType {
    // function body
    return value
}
```

## Anonymous Functions
Functions without names, often used as closures:
```go
func() {
    fmt.Println("Anonymous function")
}()

// As a variable
greet := func(name string) {
    fmt.Println("Hello, " + name)
}
greet("Alice")
```

## Multiple Return Values
Go functions can return multiple values:
```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

result, err := divide(10, 2)
```

## Named Return Values
Returns can be named in the function signature:
```go
func rectangle(length, width float64) (area, perimeter float64) {
    area = length * width
    perimeter = 2 * (length + width)
    return
}
```

## Variadic Functions
Functions accepting variable number of arguments:
```go
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

sum(1, 2, 3, 4, 5)
```

## Call by Value
Go passes arguments by value (copies are made):
```go
func modify(x int) {
    x = 100  // only modifies local copy
}

num := 5
modify(num)
fmt.Println(num)  // prints 5
```
