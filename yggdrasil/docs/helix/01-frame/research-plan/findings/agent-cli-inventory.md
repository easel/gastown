# Agent CLI Inventory

**Date**: 2026-02-01
**Purpose**: Initial spike implementation for installed agent CLIs.

## Detected Binaries

| Agent | Path | Version |
|------|------|---------|
| Claude Code | /home/erik/.local/bin/claude | 2.1.29 (Claude Code) |
| Gemini CLI | /home/linuxbrew/.linuxbrew/bin/gemini | 0.26.0 |
| Codex CLI | /home/linuxbrew/.linuxbrew/bin/codex | codex-cli 0.91.0 |
| OpenCode CLI | /home/erik/.opencode/bin/opencode | 1.1.23 |

## Notes
- All four target CLIs are installed locally.
- Next step: inspect `--help` / subcommands to identify control surfaces and state signals.

## Help Output Captured
- `docs/helix/01-frame/research-plan/findings/claude-help.txt`
- `docs/helix/01-frame/research-plan/findings/gemini-help.txt`
- `docs/helix/01-frame/research-plan/findings/codex-help.txt`
- `docs/helix/01-frame/research-plan/findings/codex-exec-help.txt`
- `docs/helix/01-frame/research-plan/findings/opencode-help.txt`

## Tmux PTY Outputs
- `docs/helix/01-frame/research-plan/findings/tmux-claude-help.txt`
- `docs/helix/01-frame/research-plan/findings/tmux-gemini-help.txt`
- `docs/helix/01-frame/research-plan/findings/tmux-codex-help.txt`
- `docs/helix/01-frame/research-plan/findings/tmux-opencode-help.txt`

## Tmux TUI Outputs
- `docs/helix/01-frame/research-plan/findings/tmux-claude-tui.txt`
- `docs/helix/01-frame/research-plan/findings/tmux-gemini-tui.txt`
- `docs/helix/01-frame/research-plan/findings/tmux-codex-tui.txt`
- `docs/helix/01-frame/research-plan/findings/tmux-opencode-tui.txt`

## Session Telemetry Samples
- `docs/helix/01-frame/research-plan/findings/metrics-claude.txt`
- `docs/helix/01-frame/research-plan/findings/metrics-gemini.txt`
- `docs/helix/01-frame/research-plan/findings/metrics-codex.txt`
- `docs/helix/01-frame/research-plan/findings/metrics-opencode.txt`

## Token Usage Samples
- `docs/helix/01-frame/research-plan/findings/claude-print.json`
- `docs/helix/01-frame/research-plan/findings/codex-exec.jsonl`
- `docs/helix/01-frame/research-plan/findings/opencode-stats.txt`
- `docs/helix/01-frame/research-plan/findings/gemini-print.json`
