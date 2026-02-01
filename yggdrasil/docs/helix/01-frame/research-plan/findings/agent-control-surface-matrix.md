# Agent Control Surface Matrix (Draft)

**Date**: 2026-02-01
**Status**: Draft (pending analysis of help output and spikes)

| Agent | Start Method | Attach/TUI | Send Input | Stop/Terminate | Session State Signal | Token Usage Exposure | Notes |
|------|--------------|------------|------------|----------------|----------------------|---------------------|-------|
| Claude Code | Interactive default; `-p/--print` for non-interactive | No explicit attach; interactive TUI via PTY/tmux (trust prompt observed) | Stdin (print mode) and interactive input | Process stop | `--session-id`, `--resume`, `--continue`, `--fork-session` | Confirmed via `--output-format json` (usage + total_cost_usd) | `setup-token` requires subscription |
| Gemini CLI | Interactive default; `-p/--prompt` for non-interactive | No explicit attach; interactive via PTY/tmux (no output in 2s capture) | Stdin + `--prompt`; `-i` continues interactive | Process stop | `--list-sessions`, `--resume`, `--delete-session` | Not confirmed; `--output-format json/stream-json` exists | `--prompt` run hit quota (capacity) |
| Codex CLI | Interactive default; `exec` for non-interactive | TUI with `--no-alt-screen`; no explicit attach | Non-interactive via `exec` | Process stop | `resume`/`fork` commands | Confirmed via `exec --json` (usage in JSONL) | `login/logout` subcommands |
| OpenCode CLI | `opencode` TUI; `run` non-interactive; `serve` headless | `attach <url>` for running server; TUI default (no output in 2s capture) | `run [message..]` or interactive TUI | `session` management; process stop | `session` management; `export` JSON | `stats` shows token usage/cost | Provides `export/import` JSON |

## References
- `docs/helix/01-frame/research-plan/findings/claude-help.txt`
- `docs/helix/01-frame/research-plan/findings/claude-print.json`
- `docs/helix/01-frame/research-plan/findings/gemini-help.txt`
- `docs/helix/01-frame/research-plan/findings/gemini-print.err`
- `docs/helix/01-frame/research-plan/findings/codex-help.txt`
- `docs/helix/01-frame/research-plan/findings/codex-exec-help.txt`
- `docs/helix/01-frame/research-plan/findings/codex-exec.jsonl`
- `docs/helix/01-frame/research-plan/findings/opencode-help.txt`
- `docs/helix/01-frame/research-plan/findings/opencode-stats.txt`
- `docs/helix/01-frame/research-plan/findings/tmux-claude-tui.txt`
- `docs/helix/01-frame/research-plan/findings/tmux-gemini-tui.txt`
- `docs/helix/01-frame/research-plan/findings/tmux-codex-tui.txt`
- `docs/helix/01-frame/research-plan/findings/tmux-opencode-tui.txt`
