
# Go Channels

## What are Channels?

- Channels are a core feature of Go that enable safe communication between goroutines. They allow you to send and receive values across goroutines in a synchronized manner, preventing race conditions. Think of a channel as a pipe — one goroutine puts data in, another takes it out.
- A channel is a typed, goroutine-safe queue with a lock, a buffer, and two wait queues (senders and receivers) — all managed by the Go runtime.

## Basic Syntax

```go
// Create a channel
ch := make(chan int)

// Send value
ch <- 42

// Receive value
value := <-ch
```

## Example: Simple Channel Communication

```go
package main
import "fmt"
func main() {
    ch := make(chan string)
    go func() {
        ch <- "Hello from goroutine"
    }()
    message := <-ch
    fmt.Println(message)
}
```

```go
make(chan int, 3)

┌─────────────────────────────────────────────┐
│              hchan (runtime struct)         │
│                                             │
│  buf      → [ 1 | 2 | 3 ]  (circular queue) │
│  sendx    → write index                     │
│  recvx    → read index                      │
│  qcount   → elements currently in buf       │
│  dataqsiz → buffer capacity                 │
│  sendq    → waiting senders (goroutine list)│
│  recvq    → waiting receivers               │
│  lock     → mutex (internal)                │
│  closed   → 0 or 1                          │
└─────────────────────────────────────────────┘
```
## Channel States & Behavior Matrix
- Channels should be closed by the sender when no more values will be sent.
```go
close(ch)
```
- **Receiving from a closed channel will not cause a deadlock; it will either receive the remaining values or zero value of the channel's type.**
```go
ch := make(chan int, 3)
ch <- 10
ch <- 20
ch <- 30
close(ch)

fmt.Println(<-ch)   // 10  ✅ real value
fmt.Println(<-ch)   // 20  ✅ real value
fmt.Println(<-ch)   // 30  ✅ real value
fmt.Println(<-ch)   // 0   ✅ zero value — no panic, no deadlock
fmt.Println(<-ch)   // 0   ✅ zero value — forever
```

## Types of Channels and Buffered Channels
Channels by default are **unbuffered**, meaning they will block on send unless there's a corresponding receive. **Buffered** channels allow you to send multiple values before blocking.
A buffered channel can hold a specified number of values without blocking. When you attempt to send to a full buffered channel or receive from an empty buffered channel, it will block until space becomes available or a value becomes available respectively.

**Full/Empty Buffered Channel === UnBuffered Channel**

Buffered channels are declared/defined with make keyword with at least of size 1.

```go
// ch := make(chan int)      // unbuffered int channel
// ch := make(chan string)   // unbuffered string channel
// ch := make(chan bool)     // unbuffered bool channel
// ch := make(chan int)     // unbuffered 
// ch := make(chan int, 0)  // unbuffered(explicit, same thing)
// ch := make(chan int, 1)  // buffered(capacity of 1)

package main

import (
	"fmt"
	"time"
	"sync"
)

func worker(id int, ch chan int) {
	// Pretend we're doing some work
	fmt.Printf("Worker %d started\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d finished\n", id)
	<-ch // Signal we're done
}

func main() {
    waitGroup := &sync.WaitGroup{}
	// Create a buffered channel
	bufferedChannel := make(chan int, 3)

	// Start 5 workers
	for i := 1; i <= 5; i++ {
        waitGroup.Go(func(){
            worker(i, bufferedChannel)
            bufferedChannel <- i // Fill the channel slot to represent a worker in-progress
	    })
    }
    waitGroup.Wait()
	// Wait for all workers to finish
	// This is a simple way to wait; in real-world scenarios, you might use sync.WaitGroup or similar
	
}
```

---
## Directional Channels (best practice)
You can restrict a channel to send-only or receive-only in function signatures:
```go
func producer(ch chan<- int) {   // can only send
    ch <- 10
}

func consumer(ch <-chan int) {   // can only receive
    fmt.Println(<-ch)
}
```
This is idiomatic Go — it makes intent clear and is caught at compile time.

## Closing a Channel + range
```go
func main() {
    ch := make(chan int, 5)

    go func() {
        for i := 0; i < 5; i++ {
            ch <- i
        }
        close(ch)   // signal: no more values
    }()

    for val := range ch {   // reads until channel is closed
        fmt.Println(val)
    }
}
```
- Rules about closing:

1. Only the sender should close
2. Sending on a closed channel panics
3. Receiving from a closed channel returns the zero value immediately

## Analogies

### Node.js
Channels are similar to **event emitters** or **Promise-based messaging**:
- `ch <- value` = `emitter.emit('event', value)` or `promise.resolve(value)`
- `<-ch` = `emitter.on('event', callback)` or `await promise`

### Python
Channels resemble **queues** from the `queue` module or **asyncio channels**:
- `ch <- value` = `queue.put(value)`
- `<-ch` = `queue.get()` or `await asyncio.gather()`

Both create a communication bridge between concurrent tasks, but Go's channels are built into the language for safer concurrency.
