
# Goroutines in Go

## Overview
Goroutines allow concurrent execution of functions with minimal overhead, making it easy to write scalable concurrent programs.

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

----
# Go WaitGroup
- Go makes it very easy to write concurrent code. Start a few goroutines, wait for all of them to finish, and done. The go keyword makes this a breeze, and with the 
**sync.WaitGroup** primitive that allows us to wait for a group of goroutines to finish, synchronization can be achieved with no hassle. So it so hassle free? ==> **NO**
- If a WaitGroup is explicitly passed into functions, it should be done by **pointer**.
- A goroutine is done when the function it invokes returns/completes the execution.
  
## How It Works ?
WaitGroup exports 3 methods.
1. **Add(int)**	It increases WaitGroup counter by given integer value.
2. **Done()**	It decreases WaitGroup counter by 1, we will use it to indicate termination of a goroutine.
3. **Wait()**	It Blocks the execution until it's internal counter becomes 0.
Note: WaitGroup is concurrency safe, so its safe to pass pointer to it as argument for Goroutines.
```go
package main

import (
    "fmt"
	"sync"
)

// func PrintNumbers(num int) {
//     fmt.Println(num)
// }

// func main() {
//     for i := 0; i <= 10; i++ {
//         go PrintNumbers(i) // This won't print unless we use waitGroups
//     }
// }

func PrintNumbers(num int, wg *sync.WaitGroup) {
	defer wg.Done()
    fmt.Println(num)
}

func main(){
	var waitGroup sync.WaitGroup
	for i:= 0; i<=10; i++{
		waitGroup.Add(1)
		go PrintNumbers(i, &waitGroup)
	}
	waitGroup.Wait()
}
```
## Depth About WaitGroup

The WaitGroup implementation, although just 129 lines of code, is actually quite interesting to look at. We can learn a lot about writing concurrent code in Go and about the runtime and the Go scheduler. Let's take a look at the WaitGroup struct:

```go
// A WaitGroup waits for a collection of goroutines to finish.
// The main goroutine calls Add to set the number of goroutines to wait for.
// Then each of the goroutines runs and calls Done when finished. At the same
// time, Wait can be used to block until all goroutines have finished.
//
// A WaitGroup must not be copied after first use.
type WaitGroup struct {
	noCopy noCopy

	// Bits (high to low):
	//   bits[0:32]  counter
	//   bits[32]    flag: synctest bubble membership
	//   bits[33:64] wait count
	state atomic.Uint64
	sema  uint32
}

// waitGroupBubbleFlag indicates that a WaitGroup is associated with a synctest bubble.
const waitGroupBubbleFlag = 0x8000_0000

// Add adds delta, which may be negative, to the [WaitGroup] task counter.
// If the counter becomes zero, all goroutines blocked on [WaitGroup.Wait] are released.
// If the counter goes negative, Add panics.
//
// Callers should prefer [WaitGroup.Go].
//
// Note that calls with a positive delta that occur when the counter is zero
// must happen before a Wait. Calls with a negative delta, or calls with a
// positive delta that start when the counter is greater than zero, may happen
// at any time.
// Typically this means the calls to Add should execute before the statement
// creating the goroutine or other event to be waited for.
// If a WaitGroup is reused to wait for several independent sets of events,
// new Add calls must happen after all previous Wait calls have returned.
// See the WaitGroup example.
func (wg *WaitGroup) Add(delta int) {
	if race.Enabled {
		if delta < 0 {
			// Synchronize decrements with Wait.
			race.ReleaseMerge(unsafe.Pointer(wg))
		}
		race.Disable()
		defer race.Enable()
	}
	bubbled := false
	if synctest.IsInBubble() {
		// If Add is called from within a bubble, then all Add calls must be made
		// from the same bubble.
		switch synctest.Associate(wg) {
		case synctest.Unbubbled:
		case synctest.OtherBubble:
			// wg is already associated with a different bubble.
			fatal("sync: WaitGroup.Add called from multiple synctest bubbles")
		case synctest.CurrentBubble:
			bubbled = true
			state := wg.state.Or(waitGroupBubbleFlag)
			if state != 0 && state&waitGroupBubbleFlag == 0 {
				// Add has been called from outside this bubble.
				fatal("sync: WaitGroup.Add called from inside and outside synctest bubble")
			}
		}
	}
	state := wg.state.Add(uint64(delta) << 32)
	if state&waitGroupBubbleFlag != 0 && !bubbled {
		// Add has been called from within a synctest bubble (and we aren't in one).
		fatal("sync: WaitGroup.Add called from inside and outside synctest bubble")
	}
	v := int32(state >> 32)
	w := uint32(state & 0x7fffffff)
	if race.Enabled && delta > 0 && v == int32(delta) {
		// The first increment must be synchronized with Wait.
		// Need to model this as a read, because there can be
		// several concurrent wg.counter transitions from 0.
		race.Read(unsafe.Pointer(&wg.sema))
	}
	if v < 0 {
		panic("sync: negative WaitGroup counter")
	}
	if w != 0 && delta > 0 && v == int32(delta) {
		panic("sync: WaitGroup misuse: Add called concurrently with Wait")
	}
	if v > 0 || w == 0 {
		return
	}
	// This goroutine has set counter to 0 when waiters > 0.
	// Now there can't be concurrent mutations of state:
	// - Adds must not happen concurrently with Wait,
	// - Wait does not increment waiters if it sees counter == 0.
	// Still do a cheap sanity check to detect WaitGroup misuse.
	if wg.state.Load() != state {
		panic("sync: WaitGroup misuse: Add called concurrently with Wait")
	}
	// Reset waiters count to 0.
	wg.state.Store(0)
	if bubbled {
		// Adds must not happen concurrently with wait when counter is 0,
		// so we can safely disassociate wg from its current bubble.
		synctest.Disassociate(wg)
	}
	for ; w != 0; w-- {
		runtime_Semrelease(&wg.sema, false, 0)
	}
}

// Done decrements the [WaitGroup] task counter by one.
// It is equivalent to Add(-1).
//
// Callers should prefer [WaitGroup.Go].
//
// In the terminology of [the Go memory model], a call to Done
// "synchronizes before" the return of any Wait call that it unblocks.
//
// [the Go memory model]: https://go.dev/ref/mem
func (wg *WaitGroup) Done() {
	wg.Add(-1)
}

// Wait blocks until the [WaitGroup] task counter is zero.
func (wg *WaitGroup) Wait() {
	if race.Enabled {
		race.Disable()
	}
	for {
		state := wg.state.Load()
		v := int32(state >> 32)
		w := uint32(state & 0x7fffffff)
		if v == 0 {
			// Counter is 0, no need to wait.
			if race.Enabled {
				race.Enable()
				race.Acquire(unsafe.Pointer(wg))
			}
			if w == 0 && state&waitGroupBubbleFlag != 0 && synctest.IsAssociated(wg) {
				// Adds must not happen concurrently with wait when counter is 0,
				// so we can disassociate wg from its current bubble.
				if wg.state.CompareAndSwap(state, 0) {
					synctest.Disassociate(wg)
				}
			}
			return
		}
		// Increment waiters count.
		if wg.state.CompareAndSwap(state, state+1) {
			if race.Enabled && w == 0 {
				// Wait must be synchronized with the first Add.
				// Need to model this is as a write to race with the read in Add.
				// As a consequence, can do the write only for the first waiter,
				// otherwise concurrent Waits will race with each other.
				race.Write(unsafe.Pointer(&wg.sema))
			}
			synctestDurable := false
			if state&waitGroupBubbleFlag != 0 && synctest.IsInBubble() {
				if race.Enabled {
					race.Enable()
				}
				if synctest.IsAssociated(wg) {
					// Add was called within the current bubble,
					// so this Wait is durably blocking.
					synctestDurable = true
				}
				if race.Enabled {
					race.Disable()
				}
			}
			runtime_SemacquireWaitGroup(&wg.sema, synctestDurable)
			isReset := wg.state.Load() != 0
			if race.Enabled {
				race.Enable()
				race.Acquire(unsafe.Pointer(wg))
			}
			if isReset {
				panic("sync: WaitGroup is reused before previous Wait has returned")
			}
			return
		}
	}
}
```
- The first thing that stands out is the **noCopy** field. This isn't data; it's a clever trick. If you try to copy a WaitGroup after its first use, the Go vet tool will yell at you. Why? Because copying it would mean the counter and waiters wouldn't be shared correctly, leading to chaos. Think of it like trying to photocopy a shared to-do list – everyone ends up with different versions!
```go
wg := sync.WaitGroup{}

wg.Add(1)

wgCopy := wg   // ❌ Copy happens here

go func() {
    wgCopy.Done() // modifies copied version
}()

wg.Wait() // original still waiting forever
```

- The second thing is the state field, an atomic.Uint64. This is where the magic happens. Instead of using separate variables (and a mutex!) for the goroutine counter (how many Done() calls are still needed) and the waiter counter (how many goroutines are blocked on Wait()), it packs them both into one 64-bit integer. The high 32 bits track the main counter, and the low 32 bits track the waiters. This atomic variable allows multiple goroutines to update and read the state safely and efficiently without needing locks, in most cases. Pretty neat, huh?

- Finally, the sema field is a semaphore used internally by the Go runtime. When a goroutine calls Wait() and the counter isn't zero, it essentially tells the runtime, "Okay, put me to sleep on this semaphore (runtime_SemacquireWaitGroup)." When the counter does hit zero (because the last Done() was called), the runtime is signaled to wake up all the goroutines sleeping on that semaphore (runtime_Semrelease). This sema field is key to understanding potential issues, as we'll see later.

In a nutshell, the WaitGroup lifecycle is:

- Call Add(n) before starting your goroutines to tell the WaitGroup how many Done() calls to expect. A common pattern is wg.Add(1) right before each go statement.
- Inside each goroutine, call defer wg.Done() immediately to ensure the counter is decremented when the goroutine finishes, no matter what.
- Call Wait() where you need to block until all n goroutines have called Done().
Now let's talk about where things can go wrong.






