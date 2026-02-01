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

## Limitations
- Only `--help` commands were executed to avoid consuming tokens or triggering logins.
- Interactive TUI behavior and session attach/detach are not yet validated.

## Next Steps
- Run interactive sessions in tmux for each agent to confirm TUI attach/detach behavior.
- Observe whether login flows or subscriptions are required for interactive sessions.
