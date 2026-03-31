# Go Data Types #
- Data type is an important concept in programming. Data type specifies the size and type of variable values.
- Go is statically typed, meaning that once a variable type is defined, it can only store data of that type.
---
Go has three basic data types:

1. **bool**: represents a boolean value and is either true or false
2. **Numeric**: represents integer types, floating point values, and complex types
3. **string**: represents a string value

```go
package main
import ("fmt")

func main() {
  var a bool = true     // Boolean
  var b int = 5         // Integer
  var c float32 = 3.14  // Floating point number
  var d string = "Hi!"  // String

  fmt.Println("Boolean: ", a)
  fmt.Println("Integer: ", b)
  fmt.Println("Float:   ", c)
  fmt.Println("String:  ", d)
}
```

# Conditionals in Go #

## Go If else Statement ##

In Go, conditional statements (if, else if, else) are used to execute different blocks of code based on whether a condition evaluates to true or false. The if statement checks a condition, else if adds more conditions, and else runs when none are true. Parentheses are optional, but {} are required.

```go 
import "fmt"
func main() {
    temperature := 70

    if temperature < 60 {
        fmt.Println("Put on a jacket.")
    } else if temperature >= 60 && temperature < 75 {
        fmt.Println("Put on a light sweater.")
    } else {
        fmt.Println("Wear summer clothes.")
    }
}
```

## Go Switch Statement ##
The Go switch statement can be used as an alternative to a set of if followed by else if statements. The switch statement compares the expression inside a condition with a set of values encapsulated in cases. The code inside a matched case value is executed and the switch block terminates. A default case without a value can be appended to the last case and its code executed if no prior match is found.
```go
import "fmt"
day := "Tuesday"
switch day {
  case "Monday":
    fmt.Println("Monday is magnificent.")
  case "Tuesday":
    fmt.Println("Tuesday is terrific.")
  case "Wednesday":
    fmt.Println("Wednesday is wacky.")
  default:
    fmt.Println("We survived.")
}
```

