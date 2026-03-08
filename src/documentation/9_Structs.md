# Structs in Go

## Definition

A struct is a composite data type that groups together variables (fields) of different types under a single name. It allows you to create complex data structures by combining simpler types.

```go
type Person struct {
    Name string
    Age  int
    City string
}
```

## Usage

### Creating a Struct Instance

```go
/* Person represents a human being with basic attributes. This struct acts like a class in object-oriented programming,encapsulating data and methods related to a person.
*/

p := Person{
    Name: "Alice",
    Age:  30,
    City: "New York",
}

// Using positional values
p := Person{"Bob", 25, "Boston"}

// Partial initialization
p := Person{Name: "Charlie"}
```

### Accessing Fields

```go
fmt.Println(p.Name)  // Output: Alice
p.Age = 31           // Modify field
```

### Struct with Methods
In Go, methods are functions associated with a specific type (struct). The syntax `(p Person)` is the receiver, which binds the method to the `Person` struct. Unlike traditional OOP languages (Java, C++, Python), Go doesn't use classes—instead, it uses structs with associated methods.

**Key differences:**
- **Go**: Methods are defined separately from the struct definition
- **OOP languages**: Methods are defined inside the class

**Comparison:**
```go
// Go - Method on struct
func (p Person) Greet() string {
    return fmt.Sprintf("Hello, I'm %s", p.Name)
}
```

```java
// Java - Method inside class
class Person {
    String greet() {
        return "Hello, I'm " + this.name;
    }
}
```

Go receivers can be **value receivers** (copy of struct) or **pointer receivers** (`*Person`) for modifying fields.
```go
func (p Person) Greet() string {
    return fmt.Sprintf("Hello, I'm %s", p.Name)
}

fmt.Println(p.Greet())
```

### Embedded Structs

```go
type Address struct {
    Street string
    City   string
}

type Employee struct {
    Name    string
    Address  // Embedded struct
}

e := Employee{Name: "John", Address: Address{Street: "Main St", City: "NYC"}}
fmt.Println(e.City)  // Accessing embedded field
```
