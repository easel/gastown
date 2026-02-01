# Feature Specification: FEAT-006 - Agent Session Management

**Feature ID**: FEAT-006
**Status**: Draft
**Priority**: P0
**Owner**: Yggdrasil Team
**Created**: 2026-02-01
**Updated**: 2026-02-01

## Overview
Manage local agent sessions for multiple CLI-based AI tools. Provide a consistent lifecycle model (start, attach, monitor, send commands, stop) with observable state, captured outputs, and resource limits.

## Problem Statement

Operators need to control multiple agent CLIs within a single orchestrator. Today, sessions are manual and inconsistent, which makes it hard to monitor progress, recover from failures, or automate workflows.

## Requirements

### Functional Requirements
- Start agent sessions for supported CLI tools (Claude Code, Gemini CLI, Codex CLI, OpenCode CLI)
- Monitor agent session state (running, idle, failed, exited)
- Attach to and observe agent TUI or output stream
- Send commands or input to a running session
- Stop or terminate sessions cleanly
- Capture outputs and logs for audit and debugging
- Track resource usage (CPU, memory, storage, IO) per session
- Track token consumption per session (when available)
- Recover or restart sessions after failure with explicit operator intent
- Preserve subscription-based access where applicable (e.g., Claude Max, ChatGPT Pro) without requiring separate API keys for MVP

### Non-Functional Requirements
- **Performance**: [NEEDS CLARIFICATION: Startup and command latency targets]
- **Security**: No secret leakage in session logs
- **Scalability**: Support multiple concurrent local sessions
- **Reliability**: Predictable session teardown and cleanup
- **Usability**: Consistent CLI behavior across agent types

## User Stories

### Story US-001: Start Agent Session [FEAT-006]
**As a** local operator
**I want** to start a new agent session for a selected CLI tool
**So that** I can begin a task in a controlled, trackable environment

**Acceptance Criteria:**
- [ ] Given a supported agent type, when I start a session, then a unique session ID is returned
- [ ] Given a start request, when the agent launches, then session state becomes "running"
- [ ] Given an unsupported agent type, then the CLI returns a clear error

### Story US-002: Observe Agent Output [FEAT-006]
**As a** local operator
**I want** to attach to a session and observe its output or TUI
**So that** I can monitor progress and intervene when needed

**Acceptance Criteria:**
- [ ] Given a running session, when I attach, then I can view live output
- [ ] Given an attached session, when I detach, then the session continues running

### Story US-003: Send Commands to Session [FEAT-006]
**As a** local operator
**I want** to send commands/input to a running session
**So that** I can guide the agent during execution

**Acceptance Criteria:**
- [ ] Given a running session, when I send input, then the agent receives it
- [ ] Given a non-running session, when I send input, then the CLI returns a clear error

### Story US-004: Stop or Terminate Session [FEAT-006]
**As a** local operator
**I want** to stop a session safely
**So that** resources are reclaimed and state is consistent

**Acceptance Criteria:**
- [ ] Given a running session, when I stop it, then the session transitions to "stopped" or "exited"
- [ ] Given a stopped session, then logs and outputs remain accessible

### Story US-005: Track Resource and Token Usage [FEAT-006]
**As a** local operator
**I want** to track resource and token usage for each session
**So that** I can observe cost and performance characteristics

**Acceptance Criteria:**
- [ ] Given a running session, resource usage (CPU, memory, storage, IO) is recorded
- [ ] Given a running session, token usage is recorded when the agent provides it
- [ ] Usage data is accessible via CLI and persisted with session logs

## Edge Cases and Error Handling
- Attempting to start two sessions with the same identifier
- Agent CLI exits immediately with an error
- Agent CLI produces TUI output that cannot be captured
- Usage data is unavailable or partial for a given agent

## Success Metrics
- [NEEDS CLARIFICATION: Max acceptable session start time]
- [NEEDS CLARIFICATION: Maximum acceptable session failure rate]
- [NEEDS CLARIFICATION: Maximum concurrent sessions supported]

## Constraints and Assumptions

### Constraints
- Local-only sessions for MVP
- CLI-first control surfaces
 - Must be compatible with subscription-based agent access flows

### Assumptions
- Agent CLIs are installed and available locally
- Operators can provide required credentials for agent CLIs

## Research/Spikes Required
- For each supported agent, identify available control surfaces (CLI flags, stdin, TUI attach, APIs)
- Determine how to monitor session state and usage metrics per agent
- Evaluate whether a wrapper is required or if attach/control without wrapping is viable

**Related Plans**:
- `docs/helix/01-frame/research-plan/RP-001-agent-control-surfaces.md`
- `docs/helix/01-frame/research-plan/RP-002-session-telemetry.md`
- `docs/helix/01-frame/research-plan/RP-003-session-control-approaches.md`

## Dependencies
- Core CLI runtime (FEAT-001)
- Task delegation workflows (FEAT-004)
- Cross-agent communication (FEAT-009)

## Out of Scope
- Distributed or remote session management
- Agent-specific UI customization beyond attach/detach

## Open Questions
1. [NEEDS CLARIFICATION: What session states are required beyond running/idle/failed/exited?]
2. [NEEDS CLARIFICATION: What telemetry fields are required per session beyond usage metrics?]
3. [NEEDS CLARIFICATION: What control surfaces are available per agent (CLI, TUI, APIs)?]
4. [NEEDS CLARIFICATION: Do we need to wrap agents or can we attach/control them without wrapping?]
5. [NEEDS CLARIFICATION: How do we preserve subscription benefits (Claude Max, ChatGPT Pro, etc.)?]

## Traceability

### Related Artifacts
- **Parent PRD Section**: Requirements Overview (P0)
- **Feature Registry**: `docs/helix/01-frame/feature-registry.md`
- **User Stories**: `docs/helix/01-frame/user-stories/FEAT-006-agent-sessions.md`

---
*Note: This specification focuses on local session lifecycle and observability. Implementation details are deferred to Design phase.*
