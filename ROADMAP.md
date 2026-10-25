# Roadmap — Disk Usage Analyzer

My execution plan for this project, broken into phases. I'll check items off as I go.

## Phase 1 — Small Project (sequential)

> Traversal uses manual recursion with `os.ReadDir` — a function that calls itself for each
> subdirectory. It makes per-directory sizes natural (the function just returns its size) and
> Phase 2 only needs to *add* goroutines instead of rewriting the traversal.

- [ ] Set up project (`go mod init`) + hello world → verify: `main.go` & `go.mod` created, `go run .` prints output
- [ ] Read directory path from CLI arg (`os.Args`) → verify: path printed back, clear error if missing
- [ ] List one directory's contents, non-recursive (`os.ReadDir`) → verify: entry names + type (file/dir) printed
- [ ] Sum file sizes in one directory (`entry.Info().Size()`) → verify: total matches File Explorer's "Size" (not "Size on disk")
- [ ] Make it recursive: a function that calls itself for each subdirectory and returns the total size → verify: total size of a large folder matches a reference tool; symlinks/junctions are not followed (no double-counting)
- [ ] Handle `permission denied` gracefully — skip that directory and keep going instead of crashing → verify: scanning a restricted folder (e.g. `C:\Users\<me>`) doesn't stop the whole program
- [ ] Size per top-level subdirectory (struct `{Path, Size}` + slice) → verify: size of each top-level subdirectory printed
- [ ] Sort descending (`slices.SortFunc`) + human-readable output (KB/MB/GB) in its own function → verify: output sorted largest → smallest
- [ ] Replace `fmt.Println` debug output with `slog` (e.g. log which dirs were skipped and why) → verify: skipped-dir messages show up as structured log lines
- [ ] Comparison: compute the same total once with `filepath.WalkDir` in a separate function → verify: same total as the recursive version; can explain the callback/closure and why `WalkDir` is hard to parallelize

## Phase 2a — Medium Project: Safety Net

> Before touching concurrency, lock down the current behavior with tests and tooling — like
> taking a baseline before a risky production change. The parallel version in 2b must pass
> the same tests.

- [ ] Move scan logic into `internal/scanner` package → verify: `go build ./...` works, output unchanged
- [ ] Table-driven tests for the human-readable formatter (`t.Run` subtests) → verify: `go test -v ./...` passes
- [ ] Test the scanner against a temp directory tree with known file sizes (`t.TempDir`) → verify: expected totals match — this becomes the reference test for the parallel version
- [ ] Set up `.golangci.yml` (golangci-lint v2 config) → verify: `golangci-lint run` passes
- [ ] Set up `gosec` (insecure pattern scan) + `govulncheck` (dependency vuln scan) → verify: both run clean
- [ ] `Justfile` (`build`, `test`, `lint`, `format`, `security`, `check`) wrapping `go mod tidy`, `go vet ./...`, `go fix -diff ./...` (modernizers), `golangci-lint run`, `goimports -w .`, `go test -race ./...`, `govulncheck ./...`
- [ ] Verify: `just check` all green, push

## Phase 2b — Medium Project: Concurrency

> Concurrency goes here, not Phase 3 like my Go rules default to — it's the whole point of this
> project. Each step deliberately shows the bug first, then fixes it.

- [ ] Identify what can be parallelized → verify: can explain which subdirectories are independent, and why manual recursion parallelizes easily
- [ ] Naive version: one goroutine per subdirectory (`sync.WaitGroup.Go`), unbounded, writing to a shared total without locking → verify: `go test -race ./...` *catches* the data race (expected failure — that's the lesson)
- [ ] Fix the race (mutex, channel, or `sync/atomic`) → verify: `go test -race ./...` clean, scanner test still passes
- [ ] Try bounding with `errgroup.SetLimit` inside the recursion → verify: observe it hang (deadlock) on a deep tree and can explain why (parents hold slots while waiting for slots for their children)
- [ ] Bound it correctly: a semaphore (buffered channel) around only the `os.ReadDir` call → verify: no deadlock, results identical to sequential, `-race` clean
- [ ] Benchmark sequential vs parallel (`go test -bench`, using `b.Loop`) → verify: have numbers showing when parallel wins
- [ ] Verify: `just check` all green, push per feature

## Phase 3 — Large Project (optional)

This CLI is realistically "done" after Phase 2 — Phase 3 is optional polish. I'd only come back to it if I want to harden it further, and concurrency shows up once more here (Ctrl+C handling).

- [ ] Handle Ctrl+C (SIGINT) during a long scan via `signal.NotifyContext` → pass the `context.Context` into the recursive function and cancel in-flight goroutines instead of an abrupt kill → verify: pressing Ctrl+C mid-scan stops within ~1s, no leftover goroutines (proven with the `goroutineleak` profile in `runtime/pprof`, GA since Go 1.27)
- [ ] Build with version info (`-ldflags`), log version at startup
- [ ] Track tool dependencies with `tool` directive in `go.mod` (goimports, golangci-lint, govulncheck, gosec) instead of a `tools.go` hack

## Backlog

Not committing to these — just don't want to lose the idea.

- JSON output mode, so results can be piped into other tools
- Filter by extension/pattern
- Exclude patterns (`.git`, `node_modules`, ...)
- Interactive TUI (like `ncdu`)
