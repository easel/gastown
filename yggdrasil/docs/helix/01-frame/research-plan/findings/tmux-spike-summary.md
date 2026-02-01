# Tmux PTY Spike Summary

**Date**: 2026-02-01
**Goal**: Validate basic PTY/tmux control (launch, send keys, capture output) for agent CLIs.

## Results
- Successfully launched each agent CLI in a detached tmux session.
- Sent input via `tmux send-keys` and captured pane output.
- Captured outputs:
  - `docs/helix/01-frame/research-plan/findings/tmux-claude-help.txt`
  - `docs/helix/01-frame/research-plan/findings/tmux-gemini-help.txt`
  - `docs/helix/01-frame/research-plan/findings/tmux-codex-help.txt`
  - `docs/helix/01-frame/research-plan/findings/tmux-opencode-help.txt`
  - `docs/helix/01-frame/research-plan/findings/tmux-claude-tui.txt`
  - `docs/helix/01-frame/research-plan/findings/tmux-gemini-tui.txt`
  - `docs/helix/01-frame/research-plan/findings/tmux-codex-tui.txt`
  - `docs/helix/01-frame/research-plan/findings/tmux-opencode-tui.txt`

## TUI Observations
- **Claude Code**: Trust prompt displayed on launch.
- **Gemini CLI**: `--screen-reader` surfaces a text UI prompt and status line.
- **Codex CLI**: Trust/approval prompt displayed on launch with `--no-alt-screen`.
- **OpenCode CLI**: TUI rendered within 8s and captured in tmux.

## Limitations
- Input injection behavior not verified beyond basic `tmux send-keys`.
- Gemini output required `--screen-reader` to render in text capture.

## Next Steps
- Re-run interactive sessions with longer capture windows.
- Explore alt-screen flags or environment settings per agent.
