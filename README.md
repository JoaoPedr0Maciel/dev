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
  - [Linux and macOS — curl](#linux-and-macos--curl)
  - [Windows — PowerShell](#windows--powershell)
  - [go install](#go-install)
  - [Build from source](#build-from-source)
- [Configuration](#configuration)
- [Usage](#usage)
- [Keybindings](#keybindings)
- [Examples](#examples)
- [Updating](#updating)

---

## Installation

### Linux and macOS — curl

The fastest way. No Go required.

```bash
curl -fsSL https://raw.githubusercontent.com/JoaoPedr0Maciel/dev/main/install.sh | sh
```

The script will:
1. Detect your OS and architecture
2. Download the correct binary from the [latest release](https://github.com/JoaoPedr0Maciel/dev/releases/latest)
3. Install it to `/usr/local/bin` (or `~/.local/bin` if no write permission)

If `dev` is not found after install, add the directory to your `PATH`:

```bash
# Add to ~/.bashrc or ~/.zshrc
export PATH="$PATH:$HOME/.local/bin"
source ~/.bashrc   # or source ~/.zshrc
```

---

### Windows — PowerShell

```powershell
$url = "https://github.com/JoaoPedr0Maciel/dev/releases/latest/download/dev_windows_amd64.exe"
$dest = "$env:USERPROFILE\.local\bin\dev.exe"
New-Item -ItemType Directory -Force -Path (Split-Path $dest) | Out-Null
Invoke-WebRequest -Uri $url -OutFile $dest
```

Then add `%USERPROFILE%\.local\bin` to your `PATH`:

1. Open **Start** → search **"Environment Variables"**
2. Under **User variables**, select `Path` → **Edit**
3. Click **New** and add: `%USERPROFILE%\.local\bin`
4. Click **OK** and restart your terminal

---

### go install

If you have Go 1.24+ installed:

```bash
go install github.com/JoaoPedr0Maciel/dev/cmd/dev@latest
```

Make sure `$GOPATH/bin` (usually `~/go/bin`) is in your `PATH`:

```bash
export PATH="$PATH:$HOME/go/bin"
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

# On Windows, move dev.exe to a folder already in PATH
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

Navigate the task list, read the description, and press `Enter` to run. The TUI closes and the command's output streams live to your terminal — long-running processes like `npm run dev` or `docker compose up` work naturally.

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

### Version

```bash
dev version
```

### Update

Update `dev` to the latest release:

```bash
dev update
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

## Updating

```bash
dev update
```

Or re-run the install script:

```bash
curl -fsSL https://raw.githubusercontent.com/JoaoPedr0Maciel/dev/main/install.sh | sh
```

---

## License

MIT © [JoaoPedr0Maciel](https://github.com/JoaoPedr0Maciel)
