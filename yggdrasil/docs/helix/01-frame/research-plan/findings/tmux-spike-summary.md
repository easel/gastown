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
- **Gemini/Codex/OpenCode**: No output captured within 2s; likely alt-screen usage or longer startup.

## Limitations
- TUI capture may require longer wait or explicit non-alt-screen mode.
- Input injection behavior not verified beyond basic `tmux send-keys`.

## Next Steps
- Re-run interactive sessions with longer capture windows.
- Explore alt-screen flags or environment settings per agent.
