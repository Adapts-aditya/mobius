# Mobius prompt design

Prompts are split on purpose (harness playbook + Anthropic/OpenAI agent guidance):

| Layer | File | Role |
|-------|------|------|
| Core behavior | `pkg/agentctx/prompt.go` | Identity, workflow, tools, safety — stable across projects |
| REPL / repo | `config/repl_guide.md` | Overrides IDE tutor rules; verification commands; examples |
| Optional context | `AGENTS.md` (low priority in `config/guides.toml`) | Background architecture for Cursor; not the REPL persona |

**Precedence:** `config/guides.toml` loads multiple files; higher `priority` is appended **last** and wins on conflict (`pkg/guides/guides.go`).

## Principles used

1. **Minimal but complete** — sections with clear headers; avoid laundry lists of edge cases ([Anthropic context engineering](https://www.anthropic.com/engineering/effective-context-engineering-for-ai-agents)).
2. **Verify → act → re-verify** — gather ground truth before edits; re-run tests after ([AgentsCamp patterns](https://agentscamp.com/guides/prompting/prompt-patterns)).
3. **Tool-first** — dedicated tools over shell; parallel independent reads ([OpenAI Codex prompting guide](https://github.com/openai/openai-cookbook/blob/main/examples/gpt-5/codex_prompting_guide.ipynb)).
4. **Tone matching** — short user messages → short replies; no unsolicited onboarding.
5. **Few-shot shape** — canonical examples in `repl_guide.md` instead of vague adjectives.

## Iterating

1. Start simple; add rules only for **observed** failures (ratchet principle in `docs/HARNESS_PLAYBOOK.md`).
2. After a bad run, check `.mobius/events/<thread>.jsonl` and add one line to `repl_guide.md` or a sensor — not a paragraph in the system prompt.
3. Run `/init` or maintain `## Verification Commands` in `AGENTS.md` so sensors stay accurate.
