# Go Sample Project — Book Server API

A small but realistic REST API written in Go that manages a collection of
books. The project demonstrates the **standard Go project layout** with
multiple packages, external dependencies, tests, and a Makefile.

---

## Project Structure

```
go-sample-project/
├── cmd/
│   └── bookserver/
│       └── main.go            # Application entry point (wiring & server start)
├── internal/
│   ├── handler/
│   │   └── book.go            # HTTP handlers (request parsing, response writing)
│   ├── model/
│   │   └── book.go            # Core data types / DTOs
│   └── service/
│       ├── book.go            # Business logic (in-memory CRUD)
│       └── book_test.go       # Unit tests for the service layer
├── pkg/
│   └── response/
│       └── response.go        # Reusable JSON response helpers
├── go.mod                     # Module definition & dependency list
├── go.sum                     # Dependency checksums (auto-generated)
├── Makefile                   # Common build/test/run targets
├── .gitignore                 # Files excluded from version control
├── .vscode/
│   ├── settings.json          # VS Code workspace settings (Go, formatting)
│   ├── launch.json            # Debug/run configurations
│   └── extensions.json        # Recommended extensions
├── .zed/
│   ├── settings.json          # Zed workspace settings (gopls, formatting)
│   └── tasks.json             # Zed task runner definitions
└── README.md                  # You are here
```

### Key directories

| Directory    | Purpose |
|--------------|---------|
| `cmd/`       | Each subdirectory is a separate executable. Here we have one: `bookserver`. |
| `internal/`  | Packages that are **private** to this module — the Go compiler enforces this. |
| `pkg/`       | Packages that are safe to be imported by **other** projects. |

---

## Prerequisites

| Tool | Minimum version | Check with |
|------|-----------------|------------|
| **Go** | 1.22+ | `go version` |
| **Make** *(optional)* | any | `make --version` |
| **curl** *(optional, for testing)* | any | `curl --version` |

---

## Getting Started

### 1. Clone the repository

```bash
git clone <repo-url>
cd go-playground/src/go-sample-project
```

### 2. Download dependencies

```bash
go mod tidy
```

This reads `go.mod`, downloads missing packages, and updates `go.sum`.

**Dependencies used:**

| Package | Why |
|---------|-----|
| [`github.com/go-chi/chi/v5`](https://github.com/go-chi/chi) | Lightweight, idiomatic HTTP router |
| [`github.com/google/uuid`](https://github.com/google/uuid) | UUID generation for book IDs |
| [`github.com/rs/zerolog`](https://github.com/rs/zerolog) | Fast, structured JSON/console logging |

### 3. Run the tests

```bash
go test -v -race ./...
```

The `-race` flag enables the Go race detector — always use it during
development to catch concurrency bugs.

### 4. Build the binary

```bash
go build -o bin/bookserver ./cmd/bookserver
```

Or, if you have Make installed:

```bash
make build
```

### 5. Start the server

```bash
# Option A — run the compiled binary
./bin/bookserver

# Option B — build and run in one step
go run ./cmd/bookserver

# Option C — via Make
make run
```

By default the server listens on port **8080**. Override with:

```bash
PORT=3000 go run ./cmd/bookserver
```

---

## API Endpoints

| Method   | Path                   | Description          |
|----------|------------------------|----------------------|
| `GET`    | `/health`              | Health check         |
| `GET`    | `/api/v1/books`        | List all books       |
| `POST`   | `/api/v1/books`        | Create a new book    |
| `GET`    | `/api/v1/books/{id}`   | Get book by ID       |
| `PUT`    | `/api/v1/books/{id}`   | Update book by ID    |
| `DELETE` | `/api/v1/books/{id}`   | Delete book by ID    |

### Example requests (curl)

```bash
# Health check
curl http://localhost:8080/health

# Create a book
curl -X POST http://localhost:8080/api/v1/books \
  -H "Content-Type: application/json" \
  -d '{"title":"The Go Programming Language","author":"Donovan & Kernighan","isbn":"978-0134190440","pages":380}'

# List all books
curl http://localhost:8080/api/v1/books

# Get a specific book (replace <id> with a real UUID)
curl http://localhost:8080/api/v1/books/<id>

# Update a book
curl -X PUT http://localhost:8080/api/v1/books/<id> \
  -H "Content-Type: application/json" \
  -d '{"title":"Updated Title","author":"Updated Author","isbn":"000-000","pages":999}'

# Delete a book
curl -X DELETE http://localhost:8080/api/v1/books/<id>
```

---

## Makefile Targets

```
make build   # Compile binary → bin/bookserver
make run     # Build + start server
make test    # Run all tests with -race
make lint    # Run go vet
make clean   # Remove bin/ directory
make help    # Show this list
```

---

## How This Project Was Created (Step by Step)

Below is the exact sequence of commands that produced this project from
scratch, for anyone who wants to reproduce it or understand the process.

```bash
# 1. Create the project directory and initialise the Go module.
mkdir -p go-sample-project && cd go-sample-project
go mod init github.com/example/go-sample-project

# 2. Create the standard directory layout.
mkdir -p cmd/bookserver internal/{handler,model,service} pkg/response

# 3. Write the source files (main.go, model, service, handler, response).
#    (See the files listed in the "Project Structure" section above.)

# 4. Add external dependencies.
go get github.com/go-chi/chi/v5@latest
go get github.com/google/uuid@latest
go get github.com/rs/zerolog@latest

# 5. Tidy the module — removes unused deps, adds missing ones, updates go.sum.
go mod tidy

# 6. Verify the code compiles cleanly.
go build ./...

# 7. Run the linter.
go vet ./...

# 8. Run all tests with the race detector.
go test -v -race ./...

# 9. Build the final binary.
go build -o bin/bookserver ./cmd/bookserver

# 10. Start the server!
./bin/bookserver
```

---

## Editor Setup

### VS Code

1. Install the [Go extension](https://marketplace.visualstudio.com/items?itemName=golang.Go) (recommended automatically via `.vscode/extensions.json`).
2. Open this folder as the workspace root — settings and debug configs are picked up automatically.
3. **Debug**: Press **F5** → choose *Launch Book Server* or *Test Current Package*.

Included configs:
| File | Purpose |
|------|---------|
| `.vscode/settings.json` | Format-on-save, Go test flags (`-race -v`), lint tool, search exclusions |
| `.vscode/launch.json` | Debug launch configs (server & tests) |
| `.vscode/extensions.json` | Recommends the official Go extension |

### Zed

1. Install `gopls` (`go install golang.org/x/tools/gopls@latest`) — Zed uses it as the language server.
2. Open this folder as the project root; Zed reads `.zed/settings.json` automatically.
3. **Tasks**: Open the command palette → *task: spawn* → pick a task (Build, Run, Test, Vet).

Included configs:
| File | Purpose |
|------|---------|
| `.zed/settings.json` | Format-on-save, gopls hints & analyses, staticcheck, hard tabs |
| `.zed/tasks.json` | One-click build, run, test, and vet tasks |

---

## License

This sample project is provided for educational purposes. Use it however you
like.
