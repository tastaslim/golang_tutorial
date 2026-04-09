
# Goroutines in Go

## Overview
A goroutine is a lightweight thread managed by the Go runtime. Goroutines allow concurrent execution of functions with minimal overhead, making it easy to write scalable concurrent programs.

## Key Characteristics
- **Lightweight**: Thousands of goroutines can run simultaneously
- **Managed by runtime**: No manual thread management needed
- **Multiplexed**: Many goroutines run on fewer OS threads
- **Simple syntax**: Use the `go` keyword to launch

## So goroutines are threads?
To people new to Go, the word goroutine and thread get used a little interchangeably. This makes sense if you come from a language such as Java where you can quite literally make new OS threads. Go is different, and **a goroutine is not the same as a thread**. Threads are much more expensive to create, use more memory and switching between threads takes longer. Goroutines are an abstraction over threads and a single Operating System thread can run many goroutines.

## Basic Syntax
```go
go functionName()
go func() {
    // code here
}()
```

## Example 1: Simple Goroutine
```go
package main

import (
    "fmt"
)

func PrintNumbers(num int) {
    fmt.Println(num)
}

func main() {
    for i := 0; i <= 10; i++ {
        go PrintNumbers(i)
    }
}

// Above code won't print anything because of go routine scheduling this task to different threads and before they get completed the main function exists end up killing them.
```

# Go Concurrency: Why Goroutines Exit When `main()` Finishes

This is one of the **most important concepts in Go concurrency** — **Goroutines + Main Function Exit**.

---

## 1. Your Code

```go
func PrintNumbers(num int) {
    fmt.Println(num)
}

func main() {
    for i := 0; i <= 10; i++ {
        go PrintNumbers(i)
    }
}
```

You expect:
```
0 1 2 3 4 5 6 7 8 9 10
```

But instead nothing prints (or partial output).

---

## 2. Important Rule

**When main() exits → program exits → all goroutines are killed**
---

## 3. Are goroutines called in the order I declared them?
  
**No**

Most Operating Systems have something called a preemptive scheduler. This means that which thread is executed next is determined by the OS itself based on thread priority and other things like waiting to receive data over the network. Since goroutines are abstractions over threads, they all have the same priority and we therefore cannot control the order in which they run.
There has been discussions as far back as 2016 (you can read one such discussion [here](https://groups.google.com/g/golang-dev/c/HJcGESXfJfs)) about adding the ability to set priority on individual goroutines, but there is some pretty compelling points raised as to why its not a good idea.

---

## 4.How do I ensure my program is as performant as possible?

- There is an environment variable (**GOMAXPROCS**) that you can set which determines how many threads your go program will use simultaneously. You can use this great library from [Uber](https://github.com/uber-go/automaxprocs) to automatically set the GOMAXPROCS variable to match a Linux container CPU quota. If you are running Go workloads in Kubernetes, you should use this.
- If you set GOMAXPROCS to be 3, this means that the program will only execute code on 3 operating system threads at once, even if there are 1000s of goroutines.
- It begs the question though, does setting GOMAXPROCS to the biggest value possible mean your program will be faster?
- The answer is **no**, and it actually might make it slower. There are a few reasons for this, but the main reason is to do with context switching.
- Swapping between threads is a relatively slow operation, and can take up to 1000ns as oppose to switching between goroutines on the same thread which takes ~200ns. Therefore you may find that for your particular workload, your program is faster with a lower GOMAXPROCS value. Always profile and benchmark your programs and make sure the Go runtime configuration is only changed if absolutely required.
---

## 4. Internal Working
- The Main Goroutine is the Anchor. When you run a Go program, the Go runtime automatically creates a single, primary goroutine. This is known as the Main Goroutine. It is responsible for initializing the program and executing the main() function.
- Every other goroutine you create using the go keyword is spawned from this main "thread"(see, it's hard to avoid this word when discussing them!) of execution, but they are treated as independent, concurrent tasks by the Go scheduler.

### The Rule of Termination**
- The fundamental rule of Go concurrency is: The lifecycle of the entire Go program is tied strictly to the Main Goroutine.
- Unlike some other programming languages where the main thread will wait for all background threads to finish before shutting down, Go does not do this automatically. 
- The Go runtime does not keep track of whether spawned goroutines have finished their work.
- **main() finishes**: When the final line of code in main() is executed (or os.Exit() is called).
- **Runtime triggers shutdown**: The Go program immediately begins its teardown process.
- **Goroutines are killed**: Any other goroutines that are currently running, sleeping, or blocked are abruptly terminated. Their execution stops instantly, and any pending defer statements inside those goroutines are not executed.
```
Time
 |    [Main Goroutine]                 [Goroutine 1]            [Goroutine 2]
 |          |                                |                        |
 |        main() starts                      |                        |
 |          |                                |                        |
 |       go PrintNumbers(0) ---------------->+ (starts)               |
 |          |                                |                        |
 |       go PrintNumbers(1) ----------------------------------------->+ (starts)
 |          |                                |                        |
 |      main() does work                    ...                      ...
 |      (e.g., brief sleep)                 ...                      ...
 |          |                                |                        |
 |        main() EXITS                       |                        |
 V          |                                |                        |
===================================================================================
      PROGRAM TERMINATES                  (KILLED)                 (KILLED)
      Memory is freed.                Work left unfinished.    Work left unfinished.
===================================================================================
```
---
## How to Prevent This (Synchronization)
1. **Bad Fix (Sleep)**
Don't rely on sleep method to handle this because we don't know how much time the other go routines would take to get completed.
```go
time.Sleep(time.Second*10)
```
2. If you want the main program to wait for its spawned goroutines to finish their jobs before it exits, you must explicitly tell the Main Goroutine to wait.
The most common and idiomatic way to do this in Go is by using a sync.WaitGroup(We will learn about it in coming chapters):
- Add the number of goroutines you are spawning to the WaitGroup.
- Done is called by each goroutine when it finishes its work.
- Wait is called at the end of main(), which blocks the Main Goroutine from exiting until the WaitGroup counter reaches zero.
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func PrintNumbers(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutine completes
	
	for i := 1; i <= 3; i++ {
		fmt.Printf("Goroutine %d: %d\n", id, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2) // We are waiting for 2 goroutines

	go PrintNumbers(0, &wg)
	go PrintNumbers(1, &wg)

	wg.Wait() // main() pauses here until both goroutines call wg.Done()
	fmt.Println("All goroutines finished. Main exiting cleanly.")
}
```
---

## 6. Key Takeaways

- main is also a goroutine
- program ends when main ends
- goroutines are async & non-deterministic
- use WaitGroup / channels for sync

---

## Interview One-liner

Goroutines don't block main; when main exits, all goroutines terminate unless explicitly synchronized.

## Best Practices
- Always synchronize goroutines using channels or `sync.WaitGroup`
- Avoid leaving goroutines hanging without completion
- Handle panics in goroutines appropriately
- Be cautious with shared memory; prefer channels for communication
- We will read about channels and WaitGroup in coming lectures
- Refer https://www.youtube.com/watch?v=rAhmryWS3Ng&list=PLXQpH_kZIxTWUe-Ee-DZEX5gfeoo4tHV6&index=24 , 
