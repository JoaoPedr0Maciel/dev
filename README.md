<div align="center">

# dev

**A minimal task runner with a TUI for task discovery.**

[![Go](https://img.shields.io/badge/Go-1.24+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

[🇧🇷 Português](README.pt-br.md)

</div>

---

`dev` is a simple alternative to Makefiles and Taskfiles. Define tasks in a `dev.yaml` file and run them through an interactive TUI or directly from the command line.

```
┌────────────────────────────────────────┐
│ Dev                                    │
├────────────────────────────────────────┤
│ ▶ build                                │
│   lint                                 │
│   run                                  │
│   test                                 │
├────────────────────────────────────────┤
│ Description:                           │
│ Build application                      │
└────────────────────────────────────────┘

  ↑/↓ navigate  •  enter run  •  q quit
```

---

## Table of Contents

- [Installation](#installation)
  - [Linux](#linux)
  - [macOS](#macos)
  - [Windows](#windows)
  - [Build from source](#build-from-source)
- [Configuration](#configuration)
- [Usage](#usage)
- [Keybindings](#keybindings)
- [Examples](#examples)

---

## Installation

### Requirements

- [Go 1.24+](https://go.dev/dl/)

### Linux

**Using `go install` (recommended):**

```bash
go install github.com/JoaoPedr0Maciel/dev/cmd/dev@latest
```

This places the binary in `$GOPATH/bin` (usually `~/go/bin`). Make sure it's in your `PATH`:

```bash
# Add to ~/.bashrc or ~/.zshrc
export PATH="$PATH:$HOME/go/bin"
```

Then reload your shell:

```bash
source ~/.bashrc   # or source ~/.zshrc
```

**Verify the installation:**

```bash
dev --help   # or just: dev
```

---

### macOS

**Using `go install` (recommended):**

```bash
go install github.com/JoaoPedr0Maciel/dev/cmd/dev@latest
```

Make sure `~/go/bin` is in your `PATH`. Add to `~/.zshrc` (default shell on macOS):

```bash
export PATH="$PATH:$HOME/go/bin"
```

Reload:

```bash
source ~/.zshrc
```

**Verify the installation:**

```bash
dev
```

---

### Windows

**Using `go install`:**

Open **PowerShell** or **Command Prompt** and run:

```powershell
go install github.com/JoaoPedr0Maciel/dev/cmd/dev@latest
```

The binary is placed in `%USERPROFILE%\go\bin`. Add it to your `PATH`:

1. Open **Start** → search for **"Environment Variables"**
2. Under **User variables**, select `Path` → **Edit**
3. Click **New** and add: `%USERPROFILE%\go\bin`
4. Click **OK** and restart your terminal

**Verify in a new terminal window:**

```powershell
dev
```

---

### Build from source

Works on Linux, macOS, and Windows:

```bash
git clone https://github.com/JoaoPedr0Maciel/dev.git
cd dev
go build -o dev ./cmd/dev/

# Move to a directory in your PATH (Linux/macOS)
sudo mv dev /usr/local/bin/

# Or on Windows, move dev.exe to a folder already in PATH
```

---

## Configuration

Create a `dev.yaml` file in the **root of your project**:

```yaml
tasks:
  build:
    description: Build the application
    cmd: go build .

  test:
    description: Run all tests
    cmd: go test ./...

  run:
    description: Start the development server
    cmd: go run .

  lint:
    description: Run the linter
    cmd: go vet ./...
```

### Fields

| Field | Required | Description |
|---|---|---|
| `description` | No | Short text shown in the TUI description panel |
| `cmd` | Yes | Shell command to execute (`sh -c` on Linux/macOS, `sh -c` via Git Bash on Windows) |

> **Note:** `dev` looks for `dev.yaml` in the **current working directory**. Make sure you run `dev` from your project root.

---

## Usage

### TUI mode

Run `dev` with no arguments to open the interactive interface:

```bash
dev
```

Use the keyboard to navigate the task list, read the description, and run tasks. Results are shown inline with exit status, duration, and output.

### Direct mode

Run a task by name without opening the TUI:

```bash
dev <task-name>
```

**Example:**

```
$ dev test

Running test...

✓ Success

Task:     test
Duration: 421ms
```

**On failure:**

```
$ dev build

Running build...

✗ Failed

./main.go:12:2: undefined: someFunc
```

---

## Keybindings

| Key | Action |
|---|---|
| `↑` or `k` | Move selection up |
| `↓` or `j` | Move selection down |
| `Enter` | Run selected task |
| `q` or `Ctrl+C` | Quit |

---

## Examples

### Typical Go project

```yaml
tasks:
  build:
    description: Compile the binary
    cmd: go build -o bin/app ./cmd/app/

  test:
    description: Run tests with race detector
    cmd: go test -race ./...

  cover:
    description: Open test coverage report
    cmd: go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out

  lint:
    description: Run golangci-lint
    cmd: golangci-lint run ./...

  tidy:
    description: Tidy go modules
    cmd: go mod tidy
```

### Node.js project

```yaml
tasks:
  install:
    description: Install dependencies
    cmd: npm install

  dev:
    description: Start dev server
    cmd: npm run dev

  build:
    description: Build for production
    cmd: npm run build

  test:
    description: Run test suite
    cmd: npm test

  lint:
    description: Lint source files
    cmd: npm run lint
```

### Docker workflow

```yaml
tasks:
  up:
    description: Start all containers
    cmd: docker compose up -d

  down:
    description: Stop all containers
    cmd: docker compose down

  logs:
    description: Tail container logs
    cmd: docker compose logs -f

  build:
    description: Rebuild images
    cmd: docker compose build --no-cache
```

---

## License

MIT © [JoaoPedr0Maciel](https://github.com/JoaoPedr0Maciel)
