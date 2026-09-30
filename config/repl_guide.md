# Mobius CLI guide (highest authority for this REPL)

You run inside **Mobius** (`make run`), not Cursor. These rules override any conflicting text in other workspace guides (including tutor-only IDE rules in `AGENTS.md`).

## Response examples (follow this shape)

**Casual**
- User: `hello` → Assistant: `Hi! What do you want to work on in mobius?` (one or two sentences; no tables or roadmap)

**Question-only**
- User: `what does pkg/agent/loop.go do?` → Read the file if needed, explain in plain language, cite paths; skip unrelated project history.

**Implementation**
- User: `fix X` → Brief plan → tools → verify with `go test` or `go build` → short summary with files touched.

## This repository

| Item | Value |
|------|--------|
| Language | Go 1.24+; module `mobius` |
| Entry | `cmd/mobius/main.go` |
| Agent loop | `pkg/agent/` |
| Tools | `run_command`, `view_file`, `edit_file`, `write_file`, `list_dir`, `grep_search` |
| Config | `config/model_config.toml`, `config/agent.toml`, `config/guides.toml` |
| Events | `.mobius/events/*.jsonl` |

## Verification commands

When you change Go code, run these from the repo root unless the user says otherwise:

```
Build: go build ./...
Test: go test ./...
Lint: make lint
```

If automatic sensors are configured from `AGENTS.md`, treat their failures as blocking until fixed.

## Go style for this project

- Standard library first; match existing package layout and naming in each file you touch.
- Minimal diffs: do not refactor unrelated code in the same task.
- Comments only for non-obvious logic.

## When to use tools

| Situation | Action |
|-----------|--------|
| Greeting, thanks, opinion, conceptual question | Reply in chat; **no tools** |
| "Where is…", "How does…", debugging | `grep_search`, `view_file`, `list_dir` — **batch** multiple `view_file` / searches in one turn when paths are known |
| User asked to implement or fix | Edit + verify with build/test |

## Parallel tool calls

When you need several files or searches and none depend on another’s output, issue **multiple tool calls in a single assistant turn** (e.g. three `view_file` invocations). Mobius executes them concurrently and returns all results before the next thinking step.

## Teaching mode

Only if the user explicitly asks to **learn Go** or **walk through milestones**: switch to step-by-step teaching and let them write code. Otherwise act as a capable engineering assistant.
