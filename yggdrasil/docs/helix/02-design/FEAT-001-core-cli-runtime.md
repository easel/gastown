# Design: FEAT-001 - Core CLI & Local Runtime

**Feature ID**: FEAT-001
**Status**: Draft
**Owner**: Yggdrasil Team
**Created**: 2026-02-01
**Updated**: 2026-02-01

## Design Goals
- Single-binary CLI with deterministic startup and routing
- YAML configuration with clear precedence and minimal surprises
- OpenTelemetry logging with human-readable output to XDG paths by default
- Strict, testable diagnostics via `yg doctor`
- Fast startup (< 100 ms) with minimal dependency overhead

## Non-Goals
- GUI or web UI
- Remote execution or multi-machine orchestration
- Plugin systems or dynamic command loading

## Architecture Overview
The CLI is a single Go binary with a small runtime core that initializes configuration, logging, and module routing. Each feature area registers commands through a minimal command registry. The runtime passes a shared `RuntimeContext` to all modules, enabling consistent config, logging, and diagnostics.

```
[yg binary]
  -> bootstrap (config + logging)
  -> module registry
  -> command router
  -> domain module handlers
```

## Key Decisions
- Config format: YAML
- Config precedence: CLI flags > env vars > project config > $XDG_CONFIG_HOME/yggdrasil/config.yaml > defaults
- Logging: OpenTelemetry logs; default human-readable logs written to XDG state path
- Diagnostics: `yg doctor` checks all external dependencies and is fully test-covered

## Command Surface (MVP)
- `yg --help`
- `yg version`
- `yg doctor`
- `yg <domain> ...` (project, queue, task, session, workspace, message)

## Configuration Design

### Search Paths
- Project config: `.yggdrasil/config.yaml` resolved from repo root or current working directory
- User config: `$XDG_CONFIG_HOME/yggdrasil/config.yaml` (fallback `~/.config/yggdrasil/config.yaml`)

### Precedence
1. CLI flags
2. Environment variables (prefix `YG_`)
3. Project config
4. User config
5. Defaults

### Example (conceptual)
```
# ~/.config/yggdrasil/config.yaml
log:
  format: human
  path: ~/.local/state/yggdrasil/logs/yg.log
```

## Logging and Diagnostics

### Logging
- Use OpenTelemetry logging APIs with a small adapter for human-readable output.
- Default output: `$XDG_STATE_HOME/yggdrasil/logs/yg.log` (fallback `~/.local/state/yggdrasil/logs/yg.log`).
- No secrets in logs by default; redact known tokens and env vars when present.

### `yg doctor`
Checks are grouped and individually testable. MVP checks:
- `git` available
- `tmux` available
- Agent CLIs available (`claude`, `gemini`, `codex`, `opencode`)
- Config file parse/validation
- XDG paths writable

## Module Interface
Modules register commands and run through a shared runtime context.

```
type RuntimeContext struct {
  Config Config
  Logger Logger
  Paths  Paths
}

type Module interface {
  Name() string
  Register(reg *CommandRegistry)
}
```

The command registry binds command names to handlers, enabling predictable routing without deep CLI framework coupling.

## Runtime Paths
- Config: `$XDG_CONFIG_HOME/yggdrasil/config.yaml`
- State: `$XDG_STATE_HOME/yggdrasil/`
- Logs: `$XDG_STATE_HOME/yggdrasil/logs/`
- Cache: `$XDG_CACHE_HOME/yggdrasil/` (if needed later)

## Error Handling and Exit Codes
- Unknown command -> exit code 2, with help summary
- Invalid config -> exit code 1, with file/line hints
- Missing dependency -> exit code 1, with `yg doctor` suggestion

## Performance Considerations
- Lazy-load domain modules only when invoked
- Avoid network calls during startup
- Keep config parsing and logging initialization lightweight

## Test Strategy
- Unit tests for config precedence and parsing
- Golden tests for `yg doctor` output and error formatting
- CLI routing tests (unknown command, help output)
- Startup timing test (target < 100 ms)

## Risks and Mitigations
- Startup time regression -> benchmark in CI
- Config ambiguity -> strict precedence and explicit errors
- Logging overhead -> defer log sinks unless used

## Open Questions
None

---
*Design phase document for FEAT-001.*
