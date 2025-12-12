# Terminal-based Todo App (Go)

A minimal day-0 Go CLI project that lets you manage todos from the terminal using a CSV file for storage.

Built to understand:
- How Go handles CLI arguments
- File I/O
- CSV reading/writing
- Building & installing a global binary

**No frameworks. No abstractions. Just Plain Go.**

---

## Features

- Add todos from the terminal
- List all todos
- Mark a todo as done (currently removes it)
- Stores data in a simple `todo.csv` file
- Works as a global CLI command

---

## Commands

```bash
# Add a new todo
todo add buy milk from store

# List all todos
todo list

# Mark todo #1 as done (removes it)
todo done 1
```

---

## CSV Storage Format

Currently the CSV file looks like this:

```csv
buy milk from store
learn golang
```

Each row represents one todo.

---

## How It Works (High Level)

### CLI Input
Go reads terminal arguments via `os.Args`:
- `os.Args[1]` → command (`add`, `list`, `done`)
- `os.Args[2:]` → task text or ID

### Persistence
- Todos are stored in a CSV file (`todo.csv`)
- Uses Go's `encoding/csv` package
- For delete/done:
  1. Read entire CSV
  2. Modify data in memory
  3. Rewrite the file

### Global Command
- Compiled into a binary
- Installed into `/usr/local/bin`
- Works like any real CLI tool

---

## Installation

### Build the binary
```bash
go build -o todo
```

### Move it into PATH
```bash
sudo mv todo /usr/local/bin/
```

### Verify
```bash
todo list
```

---

## My Key Takeaways

### Go CLI Arguments
- `os.Args[0]` → program name
- `os.Args[1]` → command
- `os.Args[2:]` → arguments
- `go run` vs global binary changes argument indexing

### CSV in Go
- CSV rows are always `[]string`
- Entire file = `[][]string`
- You cannot delete a row in place
- To delete/update:
  1. Read all rows
  2. Modify in memory
  3. Rewrite the file

### File Handling
- `os.OpenFile` flags matter (`APPEND`, `CREATE`, `WRONLY`)
- Files must be closed (`defer file.Close()`)
- CSV writers buffer data → `writer.Flush()` is mandatory

### defer
- Runs when the function exits
- Ensures cleanup even on errors
- Used for:
  - `file.Close()`
  - `writer.Flush()`

---

## Possible Improvements

- [ ] Add `done` status instead of deleting rows
  - Store CSV as: `task,done`
- [ ] Add `todo delete` command
- [ ] Add `[ ]` / `[x]` indicators in list output
- [ ] Add `--help` flag
- [ ] Replace CSV with SQLite later

---
