# What is a pointer? #
A pointer is a variable that stores the memory address of another variable.

# Pointers in Go #
1. *T is the type of the pointer variable which points to a value of type T. It is quite similar to pointers in C++.

2. Let’s write a program that declares a pointer.

```go
package main

import (
	"fmt"
)

func main() {
	b := 255
	var a *int = &b
    fmt.Printf("Type of a is %T\n", a)
	fmt.Println("address of b is", a)
}
```

3. The & operator is used to get the address of a variable. In above example, a is said to point to b. When we print the value in a, the address of b will be printed.
4. **The zero value of a pointer is nil. This means if a variabe is not assigned any value, it's value would be nil**
```go
package variableTask
import (
    "fmt"
)
func main(){
    var tas *string
    if tas == nil{
        fmt.Println("YES")
    }
}

// Output ==> YES
```

# Creating pointers using the new function #

Go also provides a handy function new to create pointers. The new function takes a type as an argument and returns a pointer to a newly allocated zero value of the type passed as argument.The following example will make things more clear.
```go
package main

import (
	"fmt"
)

func main() {
	size := new(int)
	fmt.Printf("Size value is %d, type is %T, address is %v\n", *size, size, size)
	*size = 85
	fmt.Println("New size value is", *size)
}
```

# Pass By Reference #
- In Golang, **slices** and **maps** are passed by reference by default. This means modifications inside the function will affect the original data structure.
- To pass variable as passby reference, we use & and * keywords

```go
package main
import (
	"fmt"
)

func PassByReference(pbr * int){
	*pbr = 10
}

func main(){
	pbr := 5
	PassByReference(&pbr)
	fmt.Println(pbr) // 10
}
```