Arrays are used to store multiple values of the same type in a single variable, instead of declaring separate variables for each value.

In Go, there are two ways to declare an array:

1. With the var keyword:
```go
var arrayName = [length]datatype{values}
var arrayName = [...]datatype{values} // here length is inferred
```
2. With the := sign:
```go
arrayName := [length]datatype{values} // here length is defined
arrayName := [...]datatype{values} // here length is inferred
```
---

```go
package main
import ("fmt")

func main() {
  var arr1 = [3]int{1,2,3}
  arr2 := [5]int{4,5,6,7,8}

  fmt.Println(arr1)
  fmt.Println(arr2)
}
```
# Slices #

If you’re new to Go, you’ve likely encountered slices, which are a powerful and flexible data structure. Slices make it easy to manage collections of elements like numbers, strings, or custom objects. However, one of the most fascinating parts of slices is how they grow dynamically. 

Let’s break this down in simple terms.

# What Is a Slice in Go? #
Think of a slice as a resizable box that holds a list of items. You can start with just a few items, but if you need more room, the box can automatically grow bigger to fit more stuff.

**A slice has three main properties:**

1. Length: How many items are currently in the slice?
2. Capacity: How many items can the slice hold before needing to grow?
3. Pointer: A reference to the underlying “storage area” (called an array) where the items are kept.

# How Do Slices Grow? #
When you create a slice, you can either specify its size upfront or leave it to grow dynamically.

```go
nums := []int{1, 2, 3} // Slice with 3 items
```

Here, the slice has:

- Length: 3 (because it holds 3 items)
- Capacity: 3 (because the initial storage space matches the length)

If you try to add more items using append:

```go
nums = append(nums, 4) // Add another item
```
Go checks if the slice has enough capacity. If not:

- It creates a bigger storage area (usually double the size).
- It copies the existing items into the new storage.
- It adds the new item.

# Growth Pattern: From Small to Big #
Slices grow in predictable steps, but the way they grow depends on their size.

## Small Slices (Rapid Growth) ##
For smaller slices, the capacity typically doubles when you exceed it. This ensures that you don’t keep reallocating for every small addition.
**Capacity: 2 → 4 → 8 → 16 → ...**
This rapid growth is efficient for small data because:

- It reduces the number of times Go needs to allocate new storage.
- It avoids frequent copying of data.

## Large Slices (Slower Growth) ##
As slices get larger, the growth becomes more conservative to save memory. Instead of doubling, Go increases capacity by about 25%. For instance:

**Capacity: 1024 → 1280 → 1600 → ...**
Why slow down? Because doubling a large slice wastes a lot of memory and takes longer to copy all the elements to the new storage.

```go
package main
import "fmt"
func main() {
    nums := make([]int, 0, 2) // Start with capacity 2
    for i := 0; i < 10; i++ {
        nums = append(nums, i)
        fmt.Printf("After appending %d: len=%d, cap=%d\n", i, len(nums), cap(nums))
    }
}
/*
Output:

After appending 0: len=1, cap=2
After appending 1: len=2, cap=2
After appending 2: len=3, cap=4
After appending 3: len=4, cap=4
After appending 4: len=5, cap=8
After appending 5: len=6, cap=8
After appending 6: len=7, cap=8
After appending 7: len=8, cap=8
After appending 8: len=9, cap=16
After appending 9: len=10, cap=16
*/
```