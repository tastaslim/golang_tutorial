# Introduction #
1. Go is a cross-platform, open source programming language
2. Go can be used to create high-performance applications
3. Go is a fast, statically typed, compiled language known for its simplicity and efficiency
4. Go was developed at Google by Robert Griesemer, Rob Pike, and Ken Thompson in 2007
5. Go's syntax is similar to C++
---

# What is Go Used For? #
1. Web development (server-side)
2. Developing network-based programs
3. Developing cross-platform enterprise applications
4. Cloud-native development

---

# Why Use Go? #
1. Go is fun and easy to learn
2. Go has fast run time and compilation time
3. Go supports concurrency
4. Go has memory management
5. Go works on different platforms (Windows, Mac, Linux, Raspberry Pi, etc.)

---

# Go Compared to Python and C++ #

| Go | Python | C++ |
|----|--------|-----|
| Statically typed | Dynamically typed | Statically typed |
| Fast run time | Slow run time | Fast run time |
| Compiled | Interpreted | Compiled |
| Fast compile time | Interpreted | Slow compile time |
| Supports concurrency through goroutines and channels | No built-in concurrency mechanism | Supports concurrency through threads |
| Has automatic garbage collection | Has automatic garbage collection | Does not have automatic garbage collection |
| Does not support classes and objects | Has classes and objects | Has classes and objects |
| Does not support inheritance | Supports inheritance | Supports inheritance |


*_Compilation time refers to translating the code into an executable program*
*_Concurrency is performing multiple things out-of-order, or at the same time, without affecting the final outcome*
*_Statically typed means that the variable types are known at compile time*

# How to Build and Run #
We need to run 
```go 
go build filePath
```
 to build go package and then ./executableFilePath

We can also do it in one command using 
```go
go run filePath
```

# Install Packages in go #

```go
go get <PACKAGE_NAME>
```
---

Go Follows same pattern for code comment and document string like C++.
// ==> TO SINGLE LINE COMMENT
/* */ ==> TO MULTI LINE COMMENT


---

A Go file consists of the following parts:

Package declaration
Import packages
Functions
Statements and expressions

```go
package main
import ("fmt")

func main() {
  fmt.Println("Hello World!")
}
```

1. In Go, every program is part of a package. We define this using the package keyword. In this example, the program belongs to the main package.
2. import ("fmt") lets us import files included in the fmt package.
3. A blank line. Go ignores white space. Having white spaces in code makes it more readable.
4. func main() {} is a function. Any code inside its curly brackets {} will be executed.
5. fmt.Println() is a function made available from the fmt package. It is used to output/print text. In our example it will output "Hello World!".

*_Note: In Go, any executable code belongs to the main package.*

--- 

# Naming Convention #
1. We should give the package name for go files present in a directory as package name.
```golang
/*
src -> conditionals -> 1.go, 2.go

1.go

package conditionals
...

2.go

package conditionals
...

*/
```

2. The methods which will be imported and used in different files should start with capital letter and then follows camel casing convention(Basically came casing convention with first letter as capital). Rest methods which are private can start with small letter exactly like camel casing. Same goes for any variable, struct member or anything which is present in different package and needs to be imported and used to be in different package must start with Capital letter.
```golang
//conditionals/1.go
package conditionals
import (
  "fmt"
)

func addition(a int, b int) int {
  return a+b
}

func ArithmeticAddition(a int, b int) int {
  return addition(a,b)
}

// main.go

package main
import (
  "fmt"
  "git.druva.org/cloudapps/shareddrive-node/src/code/conditionals"
)

func main(){
  var summation int = conditionals.ArithmeticAddition(1,3)
  fmt.Println(summation)
}
```

---

# Running non Main Go File #
1. Go can not run any other package than main. If you try to run it, you will get below error
```text
package command-line-arguments is not a main package
```