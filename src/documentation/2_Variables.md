# Declaring Variables #
In Go, there are two ways to declare a variable:
1. With the var keyword:
Use the var keyword, followed by variable name and type. In this case you always have to specify either type or value (or both).

```go
var variablename type = value
```
Note: 

2. With the **:=** sign:
Use the := sign, followed by the variable value. In this case, the type of the variable is inferred from the value (means that the compiler decides the type of the variable, based on the value). It is not possible to declare a variable using :=, without assigning a value to it.

```go
variablename := value
```

---
# Default Arguments #

Go does not support default parameters. We usually use wrapper functions or config structs.
1. 
```go
func helper(name string, age int, address string, salary float32, isResident bool) {

    println(name, age, address, salary, isResident)
}

func helperDefault(name string, age int, address string, salary float32) {

    helper(name, age, address, salary, true)
}
```

2. The enterprise grade practice
```go
type Employee struct {
    Name       string
    Age        int
    Address    string
    Salary     float32
    IsResident bool
}

func helper(emp Employee) {

    if !emp.IsResident {
        emp.IsResident = true
    }

    println(emp.Name, emp.Age, emp.Address, emp.Salary, emp.IsResident)
}```
---

# 🖨️ Print vs Println vs Printf in Go #

All three functions belong to the fmt package and are used to print output to the console.

🔹 1️⃣ fmt.Print()

Prints output without adding a new line

Does not automatically add spaces between arguments (unless needed for formatting)
```go
fmt.Print("Hello")
fmt.Print("World")
```

```shell
Output:
HelloWorld
```

🔹 2️⃣ fmt.Println()

Prints output with a new line at the end

Automatically adds a space between arguments
```go
fmt.Println("Hello", "World")
fmt.Println("Go")
```

```shell
Output:
Hello World
Go
```

🔹 3️⃣ fmt.Printf()

Used for formatted output. 
Does NOT add a new line automatically
Uses format specifiers
```go
name := "Taslim"
age := 25
fmt.Printf("Name: %s, Age: %d", name, age)
```

```shell
Output:
Name: Taslim, Age: 25
```


If you want a new line:

```go
fmt.Printf("Name: %s\n", name)
```

---

# Go: `var` vs `:=` #

## 1. Using `var`

```go
var a1, b1 int = 1, 2

var x int      // default value = 0
var y = 10     // type inferred

```
- Explicit variable declaration
- Type can be specified or inferred
- Can be used inside and outside functions
- Allows declaration without initialization (zero values)

## Using 2. := (Short Declaration) ##
```go
a1, b1 := 1, 2
```

- Type automatically inferred
- Can only be used inside functions
- Must declare at least one new variable

```go
a := 10
a, b := 20, 30  // valid because b is new
```

## Key Differences #

| Feature | `var` | `:=` |
|----------|--------|------|
| Explicit keyword | Yes | No |
| Type required | Optional | No (inferred) |
| Works outside functions | Yes | No |
| Preferred inside functions | No | Yes |

---

# Best Practice #

- Use `var` at package/global level.
- Use `:=` inside functions for cleaner, idiomatic Go.


---

1. **Go** Follows CamelCasing covention for methods and variables and Pascal casing for Classes and same rules as C++ for variable naming.

- A variable name must start with a letter or an underscore character (_)
- A variable name cannot start with a digit
- A variable name can only contain alpha-numeric characters and underscores (a-z, A-Z, 0-9, and _ )
- Variable names are case-sensitive (age, Age and AGE are three different variables)
- There is no limit on the length of the variable name
- A variable name cannot contain spaces
- The variable name cannot be any Go keywords


# Go Constants #

If a varibles in go has fixed value and is not going to be changed in future(Immutable variable), we can declare it using 
```go
var const tas = "Tas"
```

---
# Basic Format Specifiers #
# Go Format Specifiers (%s, %d, %v and others)

This guide covers `%s`, `%d`, `%v`, and other `%` format specifiers in Go, used with:

- `fmt.Printf()`
- `fmt.Sprintf()`
- `fmt.Fprintf()`

---

## Basic Format Specifiers

| Specifier | Meaning | Example |
|----------|---------|---------|
| `%s` | String | `"John"` |
| `%d` | Integer (decimal) | `10` |
| `%f` | Float | `3.14` |
| `%t` | Boolean | `true` |
| `%v` | Default format (any type) | struct, int, string |
| `%T` | Type of variable | `int`, `string` |
| `%c` | Character | `'A'` |
| `%p` | Pointer address | `0xc00001a0` |

---

## Example

```go
package main

import "fmt"

func main() {

	name := "John"
	age := 25
	price := 12.5
	active := true

	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Price: %f\n", price)
	fmt.Printf("Active: %t\n", active)
}
```
**Most Important One → %v**
%v works for almost everything.

```go
fmt.Printf("%v\n", 10)
fmt.Printf("%v\n", "hello")
fmt.Printf("%v\n", []int{1,2,3})
```
---

| Specifier | Meaning                  |
| --------- | ------------------------ |
| `%v`      | Default value            |
| `%+v`     | Struct with field names  |
| `%#v`     | Go syntax representation |

```go
type User struct {
	Name string
	Age  int
}

u := User{"John", 25}

fmt.Printf("%v\n", u)
fmt.Printf("%+v\n", u)
fmt.Printf("%#v\n", u)

/*
{John 25}
{Name:John Age:25}
main.User{Name:"John", Age:25}
*/
```