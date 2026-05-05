
# Go Files

## What are Files in Go?

- In Go, file handling is done through the standard library, primarily the `os`, `io`, `bufio`, and `io/ioutil` (now `os` + `io` in modern Go) packages. A file is represented by the `*os.File` type, which wraps a file descriptor provided by the operating system.
- Go treats files as streams of bytes. Reading and writing are done through interfaces (`io.Reader`, `io.Writer`), which makes file operations composable with network connections, buffers, and other streams.

## Key Packages for File Operations

| Package | Purpose |
|---------|---------|
| `os` | Open, create, remove, stat files; `*os.File` type |
| `io` | Core interfaces: `Reader`, `Writer`, `Closer`, `Copy` |
| `bufio` | Buffered I/O for efficient line/chunk reading |
| `io/ioutil` | Deprecated helpers (use `os` + `io` instead) |
| `path/filepath` | Cross-platform path manipulation |

## Opening and Closing Files

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("data.txt")   // read-only
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()                  // always close with defer

    fmt.Println("File opened:", file.Name())
}
```

- `os.Open` opens a file in **read-only** mode.
- `os.Create` creates (or truncates) a file for **writing**.
- `os.OpenFile` gives full control over flags and permissions.
- **Always `defer file.Close()`** immediately after a successful open to prevent leaks.

## File Open Flags

```go
file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
```

| Flag | Meaning |
|------|---------|
| `os.O_RDONLY` | Read-only |
| `os.O_WRONLY` | Write-only |
| `os.O_RDWR`   | Read and write |
| `os.O_APPEND` | Append to end |
| `os.O_CREATE` | Create if missing |
| `os.O_TRUNC`  | Truncate to zero length |
| `os.O_EXCL`   | Fail if file already exists (with `O_CREATE`) |

The third argument (`0644`) is the Unix permission mode: owner read/write, group read, others read.

## Reading Files

### Read Entire File into Memory (small files)

```go
data, err := os.ReadFile("data.txt")
if err != nil {
    panic(err)
}
fmt.Println(string(data))
```

### Read Line-by-Line with bufio.Scanner (large files)

```go
file, _ := os.Open("data.txt")
defer file.Close()

scanner := bufio.NewScanner(file)
for scanner.Scan() {
    fmt.Println(scanner.Text())   // one line at a time
}
if err := scanner.Err(); err != nil {
    fmt.Println("scan error:", err)
}
```

### Read Fixed-Size Chunks

```go
file, _ := os.Open("data.txt")
defer file.Close()

buf := make([]byte, 1024)
for {
    n, err := file.Read(buf)
    if n > 0 {
        fmt.Print(string(buf[:n]))
    }
    if err == io.EOF {
        break
    }
    if err != nil {
        panic(err)
    }
}
```

## Writing Files

### Write Entire Content at Once

```go
data := []byte("Hello, Go files!\n")
err := os.WriteFile("output.txt", data, 0644)
if err != nil {
    panic(err)
}
```

### Write with bufio.Writer (buffered, efficient)

```go
file, _ := os.Create("output.txt")
defer file.Close()

writer := bufio.NewWriter(file)
writer.WriteString("Line 1\n")
writer.WriteString("Line 2\n")
writer.Flush()   // MUST flush or data stays in buffer
```

### Append to an Existing File

```go
file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
if err != nil {
    panic(err)
}
defer file.Close()

file.WriteString("new log entry\n")
```

## File Metadata (Stat)

```go
info, err := os.Stat("data.txt")
if err != nil {
    panic(err)
}

fmt.Println("Name:   ", info.Name())
fmt.Println("Size:   ", info.Size(), "bytes")
fmt.Println("Mode:   ", info.Mode())
fmt.Println("ModTime:", info.ModTime())
fmt.Println("IsDir:  ", info.IsDir())
```

Use `os.IsNotExist(err)` to check if a file does not exist:

```go
if _, err := os.Stat("missing.txt"); os.IsNotExist(err) {
    fmt.Println("file does not exist")
}
```

## Copying Files

`io.Copy` streams data from any `Reader` to any `Writer` — works for files, HTTP bodies, sockets, etc.

```go
src, _ := os.Open("source.txt")
defer src.Close()

dst, _ := os.Create("destination.txt")
defer dst.Close()

n, err := io.Copy(dst, src)
if err != nil {
    panic(err)
}
fmt.Printf("Copied %d bytes\n", n)
```

## Deleting and Renaming

```go
os.Remove("old.txt")              // delete a single file
os.RemoveAll("some_dir")          // delete directory recursively
os.Rename("old.txt", "new.txt")   // rename/move
```

## Working with Directories

```go
os.Mkdir("newdir", 0755)          // create one directory
os.MkdirAll("a/b/c", 0755)        // create nested path

entries, _ := os.ReadDir(".")
for _, entry := range entries {
    fmt.Println(entry.Name(), entry.IsDir())
}
```

## Cross-Platform Paths

Never hard-code `/` or `\`. Use `filepath`:

```go
import "path/filepath"

path := filepath.Join("data", "logs", "app.log")
// → "data/logs/app.log" on Unix, "data\logs\app.log" on Windows

ext  := filepath.Ext("report.pdf")    // ".pdf"
base := filepath.Base(path)           // "app.log"
dir  := filepath.Dir(path)            // "data/logs"
```

## The io.Reader / io.Writer Abstraction

`*os.File` satisfies both `io.Reader` and `io.Writer`. This is why the same functions work across files, network connections, and in-memory buffers:

```go
func CountBytes(r io.Reader) (int64, error) {
    return io.Copy(io.Discard, r)
}

// Works for all of these:
CountBytes(file)              // *os.File
CountBytes(httpResponse.Body) // network stream
CountBytes(strings.NewReader("hello")) // in-memory
```

This uniform interface is one of Go's most powerful design choices.

## Common Pitfalls

1. **Forgetting `defer file.Close()`** — leaks file descriptors; on long-running servers this eventually hits the OS limit.
2. **Forgetting `writer.Flush()`** — `bufio.Writer` holds data in memory until flushed; if the program exits first, data is lost.
3. **Loading huge files with `os.ReadFile`** — it reads the entire file into RAM. Use `bufio.Scanner` or chunked reads for large files.
4. **Ignoring the `n` from `Read`** — `Read` may return `n > 0` **and** `err == io.EOF` on the same call. Process the bytes before checking the error.
5. **Race conditions on shared files** — concurrent writers need external synchronization (mutex, channel, or file lock).

## Analogies

### Node.js
- `os.Open` / `os.Create` ≈ `fs.open` / `fs.createWriteStream`
- `bufio.Scanner` ≈ `readline.createInterface`
- `io.Copy` ≈ `readable.pipe(writable)`
- `defer file.Close()` ≈ `try/finally` around `file.close()`

### Python
- `os.Open` ≈ `open(path, 'r')`
- `defer file.Close()` ≈ `with open(...) as f:` context manager
- `bufio.Scanner` ≈ iterating directly over the file object (`for line in f`)
- `io.Copy` ≈ `shutil.copyfileobj(src, dst)`

Go's file API is lower-level and more explicit than Python's, but the interface-based design (`io.Reader`/`io.Writer`) makes it extraordinarily composable.
