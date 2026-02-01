# Agent Control Surface Matrix (Draft)

**Date**: 2026-02-01
**Status**: Draft (pending analysis of help output and spikes)

| Agent | Start Method | Attach/TUI | Send Input | Stop/Terminate | Session State Signal | Token Usage Exposure | Notes |
|------|--------------|------------|------------|----------------|----------------------|---------------------|-------|
| Claude Code | Interactive default; `-p/--print` for non-interactive | No explicit attach; interactive TUI via PTY/tmux | Stdin (print mode) and interactive input | Process stop | `--session-id`, `--resume`, `--continue`, `--fork-session` | Not explicit; `--output-format json/stream-json` may help | `setup-token` requires subscription |
| Gemini CLI | Interactive default; `-p/--prompt` for non-interactive | No explicit attach; interactive via PTY/tmux | Stdin + `--prompt`; `-i` continues interactive | Process stop | `--list-sessions`, `--resume`, `--delete-session` | Not explicit; `--output-format json/stream-json` | Sessions indexed by project |
| Codex CLI | Interactive default; `exec` for non-interactive | TUI with `--no-alt-screen`; no explicit attach | Non-interactive via `exec` | Process stop | `resume`/`fork` commands | Not explicit in help | `login/logout` subcommands |
| OpenCode CLI | `opencode` TUI; `run` non-interactive; `serve` headless | `attach <url>` for running server; TUI default | `run [message..]` or interactive TUI | `session` management; process stop | `session` management; `export` JSON | `stats` shows token usage/cost | Provides `export/import` JSON |

## References
- `docs/helix/01-frame/research-plan/findings/claude-help.txt`
- `docs/helix/01-frame/research-plan/findings/gemini-help.txt`
- `docs/helix/01-frame/research-plan/findings/codex-help.txt`
- `docs/helix/01-frame/research-plan/findings/opencode-help.txt`
