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

not available yet.

---

## Installation

not available yet.

---

## How to Run

not available yet.
---

## Roadmap

Detailed execution roadmap is in [ROADMAP.md](ROADMAP.md).

---

## License

MIT — see [LICENSE](LICENSE).
