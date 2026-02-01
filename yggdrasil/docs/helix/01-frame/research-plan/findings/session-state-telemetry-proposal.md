# Session State and Telemetry Proposal (Draft)

**Date**: 2026-02-01
**Scope**: Minimal session state machine and telemetry schema for MVP.

## Session State Model

### Core States (MVP)
- **starting**: Session process launched, pre-flight steps in progress (trust prompts, auth checks).
- **running**: Agent is active and can receive input.
- **waiting_input**: Agent is awaiting user/operator input.
- **stopping**: Session is shutting down (graceful stop requested).
- **stopped**: Session terminated normally.
- **failed**: Session terminated with error or abnormal exit.
- **unknown**: State cannot be determined reliably (fallback).

### Optional States (Post-MVP)
- **paused**: Operator paused session without terminating.
- **detached**: Session running without active attachment.
- **resuming**: Session resuming from persisted state.

### State Transition Notes
- `starting -> running` when agent accepts input or renders prompt.
- `running -> waiting_input` when prompt visible and idle for configured timeout.
- `running|waiting_input -> stopping -> stopped` on explicit stop.
- Any state -> `failed` on non-zero exit or crash.

## Telemetry Schema (MVP)

### Session Identity
- `session_id` (string)
- `agent_type` (enum: claude|gemini|codex|opencode|unknown)
- `started_at` (timestamp)
- `ended_at` (timestamp, optional)
- `state` (enum from above)

### Resource Usage (sampled)
- `cpu_percent` (float)
- `memory_rss_bytes` (int)
- `memory_vsz_bytes` (int)
- `io_read_bytes` (int)
- `io_write_bytes` (int)
- `io_read_chars` (int)
- `io_write_chars` (int)

### Workspace/Storage
- `workspace_path` (string)
- `workspace_bytes_used` (int, sampled; derived from `du -s`)

### Token Usage (when available)
- `tokens_input` (int, optional)
- `tokens_output` (int, optional)
- `tokens_cache_read` (int, optional)
- `tokens_cache_write` (int, optional)
- `cost_usd` (float, optional)

### IO and Attachment
- `attached` (bool)
- `last_input_at` (timestamp, optional)
- `last_output_at` (timestamp, optional)

## Derived/Computed Fields (Optional)
- `duration_ms`
- `avg_cpu_percent`
- `peak_memory_rss_bytes`
- `exit_code`

## Agent-Specific Mapping Notes

- **Claude Code**: `--output-format json` provides usage and cost per session.
- **Codex CLI**: `exec --json` provides usage data in JSONL.
- **OpenCode CLI**: `opencode stats` provides aggregate token usage; may need session-level mapping.
- **Gemini CLI**: Token usage not confirmed due to quota; fallback to resource-only telemetry.

## Open Questions
1. How do we detect `waiting_input` vs `running` across agents consistently?
2. Do we need to persist telemetry at high frequency, or snapshot at intervals?
3. How do we map OpenCode stats to specific sessions?
4. What default sampling interval is acceptable for MVP?

