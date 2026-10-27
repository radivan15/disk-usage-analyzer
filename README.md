# Disk Usage Analyzer

A CLI that recursively calculates directory sizes and lists them sorted from largest to smallest — a Go clone of the Unix `du` command.

> Status: work in progress

---

## What it does

Walks a directory tree recursively with `os.ReadDir`, aggregates file sizes per directory, and prints results sorted largest-first in human-readable units (KB/MB/GB) — starting as a sequential implementation, then adding goroutine-based parallelism for large directory trees.

---

## Why I'm building it

With an SRE/Cloud Engineering background, I've dealt with disk-space-full incidents in production more than once — it's one of the most common causes of outages. I'm building this to learn Go concurrency fundamentals (goroutines, worker pools, race conditions) from scratch, instead of just running `du` and moving on.

---

## Stack

Planned — nothing wired up yet, see [ROADMAP.md](ROADMAP.md) for sequencing.

| Component | Technology | Notes |
|---|---|---|
| Language | Go 1.27 | |
| Directory traversal | `os.ReadDir` + manual recursion (stdlib) | `filepath.WalkDir` built once as a comparison |
| Concurrency | goroutines + `sync.WaitGroup`, semaphore (buffered channel) | bounded; `errgroup.SetLimit` explored (deadlock pitfall with recursion) |
| Linting | golangci-lint | |
| Security scan | govulncheck, gosec | |
| Testing | `go test -race`, benchmarks (`b.Loop`) | |
| Task runner | Just (`Justfile`) | |

---

## Folder Structure

```
disk-usage-analyzer/
├── main.go
├── go.mod
├── README.md
├── ROADMAP.md
├── LICENSE
└── .gitignore
```

---

## Installation

Requires Go 1.27 or newer.

```powershell
git clone https://github.com/radivan15/disk-usage-analyzer.git
cd disk-usage-analyzer
```

---

## How to Run

Currently lists the entries of a single directory (non-recursive), marking each as `file` or `dir`:

```powershell
go run . C:\Users
```

```
file All Users
dir Default
file Default User
dir Public
file desktop.ini
dir rrdvan
```

Missing argument or an invalid path prints an error to stderr and exits with code 1.

---

## Roadmap

Detailed execution roadmap is in [ROADMAP.md](ROADMAP.md).

---

## License

MIT — see [LICENSE](LICENSE).
