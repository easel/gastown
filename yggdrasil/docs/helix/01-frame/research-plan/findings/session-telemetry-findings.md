# Session Telemetry Findings (Draft)

**Date**: 2026-02-01
**Scope**: Local OS telemetry and token usage visibility for agent sessions.

## Resource Usage (CPU/Mem/IO)

### Method
- Launch agent in tmux.
- Identify session PID via `tmux list-panes` and child process via `pgrep -P`.
- Capture `ps` metrics and `/proc/<pid>/io`.

### Samples
- `docs/helix/01-frame/research-plan/findings/metrics-claude.txt`
- `docs/helix/01-frame/research-plan/findings/metrics-gemini.txt`
- `docs/helix/01-frame/research-plan/findings/metrics-codex.txt`
- `docs/helix/01-frame/research-plan/findings/metrics-opencode.txt`

### Observations
- CPU/memory/IO metrics are available via `/proc` for the session process.
- Each agent spawns a single visible child process in this spike.
- CPU usage can exceed 100% due to multi-core reporting.

### Open Questions
- How to aggregate metrics across any child process trees (if agents spawn helpers)?
- How to sample storage usage for session workspaces (likely `du -s` on workspace root)?

## Token Usage

### Claude Code
- `claude -p --output-format json "ping"` returns usage and cost.
- Sample: `docs/helix/01-frame/research-plan/findings/claude-print.json`

### Codex CLI
- `codex exec --json "ping"` returns JSONL with `usage`.
- Sample: `docs/helix/01-frame/research-plan/findings/codex-exec.jsonl`

### OpenCode CLI
- `opencode stats` provides token usage and cost summaries.
- Sample: `docs/helix/01-frame/research-plan/findings/opencode-stats.txt`

### Gemini CLI
- `gemini --prompt "ping" --output-format json` failed due to model quota.
- Error log: `docs/helix/01-frame/research-plan/findings/gemini-print.err`

## Preliminary Recommendation
- For MVP, track CPU/memory/IO via process metrics.
- Token usage should be captured per agent via each CLI's native output where available.
- Define a minimal telemetry schema that allows partial data when agents do not expose tokens.
