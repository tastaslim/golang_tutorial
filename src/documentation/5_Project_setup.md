# Go Project Setup, Modules & Running #

## Setting Up a Go Project ##

### Prerequisites ###
1. Install Go from [https://go.dev/dl/](https://go.dev/dl/)
2. Verify installation:
```bash
go version
```
3. Ensure `GOPATH` and `GOROOT` are set (Go sets sensible defaults automatically since Go 1.11+)

### Initializing a New Project ###

Every Go project starts with `go mod init`. This creates a `go.mod` file which is the identity card of your project.

```bash
mkdir my-service && cd my-service
go mod init github.com/your-org/my-service
```

The module path (`github.com/your-org/my-service`) is how other packages import your code. For enterprise projects, this is typically a private VCS path like:
```bash
go mod init git.company.com/team/service-name
```

This project, for example, uses:
```
module git.druva.org/cloudapps/shareddrive-node
```

---

## go.mod — The Module Definition File ##

`go.mod` is the **single source of truth** for your project's identity and dependencies.

### What It Contains ###
| Directive | Purpose |
|-----------|---------|
| `module` | The module path — uniquely identifies your project |
| `go` | The minimum Go version required |
| `require` | Direct and indirect dependencies with exact versions |
| `replace` | Override a dependency's source (useful for local dev or forks) |
| `exclude` | Prevent a specific version from being used |
| `retract` | Mark versions of your own module as broken (for library authors) |

### Minimal Example ###
```
module git.druva.org/cloudapps/shareddrive-node

go 1.25.5
```

### Real-World Example with Dependencies ###
```
module git.company.com/platform/auth-service

go 1.25.5

require (
    github.com/gin-gonic/gin v1.9.1
    github.com/lib/pq v1.10.9
    go.uber.org/zap v1.27.0
)

require (
    // indirect dependencies auto-managed by Go tooling
    golang.org/x/text v0.14.0 // indirect
)
```

### Key Behaviors ###
- Running `go get <package>` automatically adds the dependency to `go.mod`
- Running `go mod tidy` removes unused dependencies and adds missing ones
- **Never edit `go.mod` by hand unless you know what you're doing** — let the tooling manage it

---

## go.sum — The Integrity Lock File ##

`go.sum` is automatically generated and contains **cryptographic checksums** (SHA-256 hashes) of every dependency (and its `go.mod`).

### Why It Exists ###
1. **Tamper detection** — guarantees that the exact same bytes are downloaded every time, on every machine
2. **Reproducible builds** — CI, coworkers, and production all get identical dependencies
3. **Supply chain security** — detects if a dependency's content changes at the same version (a compromised or yanked module)

### What It Looks Like ###
```
github.com/gin-gonic/gin v1.9.1 h1:4+fr/el88TOO3e...=
github.com/gin-gonic/gin v1.9.1/go.mod h1:ReTOC3...=
```

Each line has: `module version hash-type:hash`

### Rules ###
- **Always commit `go.sum` to version control** — it is not a generated artifact you can skip
- Never edit it manually
- If checksums mismatch, `go` refuses to build (this is a safety feature)
- Run `go mod tidy` to clean up stale entries

---

## go.mod vs go.sum — Quick Comparison ##

| Aspect | go.mod | go.sum |
|--------|--------|--------|
| Purpose | Declares module identity and dependencies | Stores checksums of all dependencies |
| Editable by hand | Cautiously, yes | Never |
| Commit to VCS | Yes | Yes |
| Created by | `go mod init` | Automatically on first `go get` or `go mod tidy` |
| Updated by | `go get`, `go mod tidy` | Same, automatically |

---

## Enterprise Configuration ##

### 1. Private Module Access (GOPRIVATE) ###

By default, Go fetches modules via the public proxy (`proxy.golang.org`) and validates checksums against the public checksum database (`sum.golang.org`). For private/internal repos, you must tell Go to skip these:

```bash
# In your shell profile (~/.zshrc, ~/.bashrc) or CI config
export GOPRIVATE="git.company.com/*,github.com/your-org/*"
```

This single variable configures both `GONOSUMDB` and `GONOPROXY` for the matching patterns.

### 2. Private Git Authentication ###

Go uses `git` under the hood to fetch modules. Configure access for private repos:

```bash
# For HTTPS (most common in enterprise)
git config --global url."https://oauth2:${ACCESS_TOKEN}@git.company.com/".insteadOf "https://git.company.com/"

# For SSH
git config --global url."git@git.company.com:".insteadOf "https://git.company.com/"
```

### 3. Private Module Proxy (Optional but Recommended) ###

Large organizations run an internal Go module proxy (e.g., Athens, Artifactory, Nexus) for caching, auditing, and availability:

```bash
export GOPROXY="https://goproxy.company.com,https://proxy.golang.org,direct"
```

This tries the private proxy first, falls back to the public one, then fetches directly.

### 4. Vendoring (Offline / Air-Gapped Builds) ###

For environments without network access or for fully hermetic builds:

```bash
go mod vendor        # copies all dependencies into a ./vendor directory
go build -mod=vendor # builds using only vendored dependencies
```

Commit the `vendor/` directory to VCS when using this approach.

### 5. Dependency Governance ###

```bash
go mod tidy           # remove unused, add missing deps
go mod verify         # verify checksums of downloaded modules match go.sum
go list -m all        # list all dependencies (direct + transitive)
go list -m -u all     # check for available updates
```

### 6. Recommended .gitignore for Go Projects ###
```
# Binaries
*.exe
*.exe~
*.dll
*.so
*.dylib

# Build output
/bin/
/dist/

# Test binary
*.test

# Go coverage
*.out

# IDE
.idea/
.vscode/
*.swp

# OS
.DS_Store
Thumbs.db

# Do NOT ignore go.sum — it must be committed
```

---

## How to Build and Run ##

### Run Directly (Development) ###
```bash
go run src/tas.go
```
`go run` compiles and executes in one step. The binary is created in a temp directory and discarded after execution.

### Build a Binary ###
```bash
go build -o bin/app src/tas.go
./bin/app
```
`go build` produces a **statically linked binary** — no runtime dependencies needed on the target machine.

### Build with Version Info (Production) ###
```bash
go build -ldflags="-s -w -X main.version=1.2.3 -X main.buildTime=$(date -u +%Y-%m-%dT%H:%M:%SZ)" \
  -o bin/app src/tas.go
```

| Flag | Effect |
|------|--------|
| `-s` | Strip symbol table (smaller binary) |
| `-w` | Strip DWARF debug info (smaller binary) |
| `-X` | Inject values into variables at build time |

### Cross-Compilation ###

Go makes cross-compiling trivial — no toolchain setup required:

```bash
# Build for Linux (e.g., deploying to a container or server)
GOOS=linux GOARCH=amd64 go build -o bin/app-linux src/tas.go

# Build for Windows
GOOS=windows GOARCH=amd64 go build -o bin/app.exe src/tas.go

# Build for ARM (e.g., Raspberry Pi, AWS Graviton)
GOOS=linux GOARCH=arm64 go build -o bin/app-arm64 src/tas.go
```

### Run Tests ###
```bash
go test ./...                    # run all tests recursively
go test -v ./...                 # verbose output
go test -race ./...              # detect race conditions
go test -cover ./...             # show coverage percentage
go test -coverprofile=cover.out ./... && go tool cover -html=cover.out  # HTML coverage report
```

### Useful Tooling Commands ###
```bash
go fmt ./...          # format all code (enforced style)
go vet ./...          # static analysis for common mistakes
go mod tidy           # sync go.mod and go.sum with actual imports
go mod download       # pre-download all dependencies (useful in CI/Docker)
go clean -cache       # clear the build cache
```

---

## Recommended Enterprise Project Layout ##

```
my-service/
├── go.mod
├── go.sum
├── Makefile              # build, test, lint, run targets
├── Dockerfile
├── README.md
├── cmd/
│   └── server/
│       └── main.go       # application entrypoint
├── internal/             # private packages (cannot be imported by other modules)
│   ├── handler/
│   ├── service/
│   ├── repository/
│   └── config/
├── pkg/                  # public reusable packages (importable by other modules)
├── api/                  # protobuf/OpenAPI specs
├── scripts/              # build and CI helper scripts
└── test/                 # integration / e2e tests
```

The `internal/` directory is enforced by the Go compiler — packages under it cannot be imported by code outside the module. This is a powerful encapsulation tool for enterprise services.

---

## Quick Reference ##

| Task | Command |
|------|---------|
| Create a new project | `go mod init <module-path>` |
| Add a dependency | `go get <package>@<version>` |
| Remove unused deps | `go mod tidy` |
| Run the program | `go run <file.go>` |
| Build a binary | `go build -o <output> <file.go>` |
| Run all tests | `go test ./...` |
| Format code | `go fmt ./...` |
| Lint/vet code | `go vet ./...` |
| Vendor dependencies | `go mod vendor` |
| Verify dependency integrity | `go mod verify` |
