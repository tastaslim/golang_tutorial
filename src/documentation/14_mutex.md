
# Mutex in Go

## Overview
A **Mutex** (short for *mutual exclusion*) is a synchronization primitive from Go's `sync` package that protects shared data when multiple goroutines read or write it concurrently. Only one goroutine can hold the lock at a time — everyone else waits.

Without a mutex (or a channel), concurrent access to the same memory leads to **race conditions**: non-deterministic bugs that are notoriously hard to reproduce and debug.

---

## 1. The Problem: Race Conditions

Consider a simple counter incremented by 1000 goroutines:

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    counter := 0
    var wg sync.WaitGroup

    for i := 0; i < 1000; i++ {
        wg.Go(func(){
            counter++ // ❌ RACE
        })()
        
    }
    wg.Wait()
    fmt.Println(counter)
}
```

You **expect** `1000`, but you'll almost always get something less — `987`, `994`, `1000`, different every run.

### Why?
`counter++` is **not atomic**. It's actually three steps at the CPU level:
1. Read `counter` from memory into a register
2. Add 1 to the register
3. Write the register back to memory

If two goroutines read the value `42` at the same time, both increment to `43`, and both write back `43` — one increment is silently lost.

### Detecting the race
Go ships with a built-in race detector. Run your program with:
```bash
go run -race main.go
```
It will print a detailed report showing which goroutines raced on which memory address. **Always run your tests and long-running programs with `-race` during development.**

---

## 2. Fix with `sync.Mutex`

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    var (
        counter int
        mu      sync.Mutex
        wg      sync.WaitGroup
    )

    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            mu.Lock()   // only one goroutine past this line at a time
            counter++
            mu.Unlock()
        }()
    }
    wg.Wait()
    fmt.Println(counter) // always 1000
}
```

### Two methods you'll use 99% of the time
| Method | Purpose |
|---|---|
| `mu.Lock()` | Acquire the lock. Blocks until the lock is free. |
| `mu.Unlock()` | Release the lock. Must be called by whoever locked it. |

### The critical section
The code between `Lock()` and `Unlock()` is called the **critical section**. Keep it as small as possible — anything in it is effectively single-threaded.

### Always prefer `defer mu.Unlock()`
```go
mu.Lock()
defer mu.Unlock()
counter++
```
Why? If the function panics or returns early, `defer` still runs. Without it, a panic leaves the mutex locked forever → every other goroutine deadlocks.

---

## 3. Idiomatic Pattern: Embed the Mutex in a Struct

Don't pass a mutex around separately — bundle it with the data it protects.

```go
package main

import (
    "fmt"
    "sync"
)

type SafeCounter struct {
    mu    sync.Mutex
    count map[string]int
}

func NewSafeCounter() *SafeCounter {
    return &SafeCounter{count: make(map[string]int)}
}

func (c *SafeCounter) Inc(key string) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.count[key]++
}

func (c *SafeCounter) Value(key string) int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.count[key]
}

func main() {
    c := NewSafeCounter()
    var wg sync.WaitGroup

    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            c.Inc("hits")
        }()
    }
    wg.Wait()
    fmt.Println("hits =", c.Value("hits")) // 1000
}
```

### Rules of thumb
- **Always use a pointer receiver** (`*SafeCounter`) on methods that touch the mutex. A value receiver copies the mutex, which is a bug `go vet` will flag.
- **Never copy a `sync.Mutex` after first use** — the copied lock has its own state, so two goroutines can both "hold" what looks like the same lock.
- Place the mutex field **above** the fields it protects, and add a comment like `// protects count` if the protection relationship isn't obvious.

```go
type Server struct {
    mu       sync.Mutex       // protects the fields below
    requests int
    errors   int
    // ---- unprotected fields below ----
    name     string
}
```

---

## 4. `sync.RWMutex` — Read-Heavy Workloads

When reads vastly outnumber writes (caches, config, lookup tables), a regular mutex is overkill — reads don't actually conflict with each other.

`sync.RWMutex` lets **many readers** hold the lock simultaneously, but writers get **exclusive** access.

```go
package main

import (
    "fmt"
    "sync"
)

type Cache struct {
    mu   sync.RWMutex
    data map[string]string
}

func NewCache() *Cache {
    return &Cache{data: make(map[string]string)}
}

func (c *Cache) Get(key string) (string, bool) {
    c.mu.RLock()            // shared — many readers allowed
    defer c.mu.RUnlock()
    val, ok := c.data[key]
    return val, ok
}

func (c *Cache) Set(key, val string) {
    c.mu.Lock()             // exclusive — blocks ALL readers and writers
    defer c.mu.Unlock()
    c.data[key] = val
}

func main() {
    c := NewCache()
    c.Set("user:1", "Alice")
    if v, ok := c.Get("user:1"); ok {
        fmt.Println(v)
    }
}
```

### Methods
| Method | Purpose |
|---|---|
| `RLock()` / `RUnlock()` | Acquire/release a **read** lock (shared) |
| `Lock()` / `Unlock()` | Acquire/release a **write** lock (exclusive) |

### When to pick `RWMutex` over `Mutex`
- Reads are ≥ 10× more frequent than writes.
- The critical section for reads is non-trivial (expensive reads benefit more from parallelism).
- Benchmark it! `RWMutex` has higher overhead per call than `Mutex`. For short critical sections, a plain `Mutex` is often faster.

### Write-starvation warning
If reads are constant and overlapping, a pending writer may wait a long time. Go's `RWMutex` prevents starvation by blocking **new** readers once a writer is waiting, but this is still something to be aware of under heavy read pressure.

---

## 5. Common Pitfalls

### 5.1 Forgetting to `Unlock`
```go
func (s *Store) Save(item Item) error {
    s.mu.Lock()
    if err := validate(item); err != nil {
        return err   // ❌ lock is never released → deadlock
    }
    s.items = append(s.items, item)
    s.mu.Unlock()
    return nil
}
```
Fix with `defer`:
```go
s.mu.Lock()
defer s.mu.Unlock()
```

### 5.2 Copying a mutex
```go
type Counter struct {
    mu sync.Mutex
    n  int
}

func (c Counter) Inc() { // ❌ VALUE receiver copies the mutex
    c.mu.Lock()
    defer c.mu.Unlock()
    c.n++
}
```
`go vet` catches this: *"Inc passes lock by value: Counter contains sync.Mutex"*. Always use **pointer receivers**.

### 5.3 Double-locking (mutexes are NOT reentrant)
```go
func (s *Store) Outer() {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.Inner() // ❌ Inner() also calls Lock() → deadlock
}

func (s *Store) Inner() {
    s.mu.Lock()
    defer s.mu.Unlock()
    // ...
}
```
Unlike Java's `synchronized` or C#'s `lock`, Go mutexes **cannot be re-acquired** by the same goroutine. Refactor so only one entry point locks:
```go
func (s *Store) Outer() {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.innerLocked() // unexported helper that assumes lock is held
}

func (s *Store) innerLocked() {
    // caller must hold s.mu
}
```

### 5.4 Locking too much (over-serialization)
```go
func (s *Server) Handle(req Request) Response {
    s.mu.Lock()
    defer s.mu.Unlock()
    data := s.fetchFromDB(req.ID)        // ❌ slow I/O inside the lock
    return process(data)                  // ❌ CPU work inside the lock
}
```
Every request now runs single-threaded. Lock only what touches shared state:
```go
func (s *Server) Handle(req Request) Response {
    s.mu.Lock()
    cached, ok := s.cache[req.ID]
    s.mu.Unlock()
    if ok {
        return process(cached)
    }
    data := s.fetchFromDB(req.ID) // slow, but outside the lock
    s.mu.Lock()
    s.cache[req.ID] = data
    s.mu.Unlock()
    return process(data)
}
```

### 5.5 Lock ordering and deadlocks
If two goroutines each hold one lock and wait for the other, neither can proceed.
```go
// Goroutine A: locks mu1 → tries to lock mu2
// Goroutine B: locks mu2 → tries to lock mu1
// → deadlock
```
**Rule:** if your code ever holds two mutexes at once, always acquire them in the **same global order** (e.g., alphabetical by field name, or by ID). Or better — redesign so you never hold two at the same time.

### 5.6 Unlocking a lock you don't hold
Calling `Unlock()` without a matching `Lock()` panics with `sync: unlock of unlocked mutex`. Pair them religiously.

---

## 6. Mutex vs Channel — When to Use Which

Go's proverb:
> **"Don't communicate by sharing memory; share memory by communicating."**

But mutexes are still the right tool for many problems. Here's the split:

| Use a **Mutex** when | Use a **Channel** when |
|---|---|
| Protecting access to a shared variable, map, or struct field | Passing ownership of a value between goroutines |
| Reference counting, caches, in-memory stats | Producer/consumer pipelines |
| The critical section is small and state-focused | Coordinating the lifecycle of multiple goroutines |
| You want the simplest correct solution | You want cancellation, fan-out/fan-in, or back-pressure |

In practice, production Go codebases use **both** — mutexes for per-object state, channels for coordination between components.

---

## 7. Related Primitives (Quick Reference)

- **`sync.Once`** — run an initialization function exactly once, safely, across goroutines.
  ```go
  var once sync.Once
  once.Do(initConfig) // runs initConfig only on the first call, even if called concurrently
  ```
- **`sync.Map`** — a concurrent map optimized for two specific patterns: write-once-read-many, and disjoint key sets per goroutine. For general use, a `sync.Mutex` + regular `map` is usually simpler and faster.
- **`sync/atomic`** — lock-free atomic operations on integers and pointers (`atomic.Int64`, `atomic.Pointer[T]`). Use only for single variables where a mutex would be overkill. Simple counters are the classic case:
  ```go
  var hits atomic.Int64
  hits.Add(1)
  fmt.Println(hits.Load())
  ```

---

## 8. Key Takeaways

- `sync.Mutex` serializes access to shared state — only one goroutine in the critical section at a time.
- Always pair `Lock()` with `defer Unlock()`.
- **Never copy** a mutex (or a struct containing one) — use pointer receivers.
- Go mutexes are **not reentrant** — don't lock the same mutex twice in the same goroutine.
- Use `sync.RWMutex` when reads massively outnumber writes, and benchmark.
- Run your code with `go test -race` and `go run -race` during development.
- Keep critical sections **small**. No I/O, no long computation, no calls to user code inside a lock.
- Prefer channels for coordinating goroutines; prefer mutexes for guarding state.

---

## Interview One-liner

*A `sync.Mutex` enforces mutual exclusion so only one goroutine at a time can execute a critical section; use `defer Unlock()` to avoid leaked locks, never copy it, and remember it is not reentrant.*

---

## Best Practices
- Embed the mutex in the struct it protects; document which fields it guards.
- Expose **thread-safe methods**, not the mutex itself — callers shouldn't need to know locking exists.
- Lock at the **narrowest scope** that still protects the invariant.
- If you're holding a lock across a function call, ask: *could that function ever try to acquire this lock again?*
- Prefer `atomic` for single-variable counters; prefer `Mutex` for anything that requires multiple fields to stay consistent together.
- When in doubt, write a test and run it with `-race`.

# The One Rule to Remember

**Every access to shared state — read or write — must be under the lock.**