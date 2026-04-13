
# Go Channels

## What are Channels?

Channels are a core feature of Go that enable safe communication between goroutines. They allow you to send and receive values across goroutines in a synchronized manner, preventing race conditions.

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
