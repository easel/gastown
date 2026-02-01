# Open Questions Checklist

**Date**: 2026-02-01
**Scope**: Aggregated [NEEDS CLARIFICATION] items from Frame artifacts.

## features/FEAT-001-core-cli-runtime.md
- **Performance**: CLI startup under [NEEDS CLARIFICATION: target ms]
- [NEEDS CLARIFICATION: CLI startup time target]
- [NEEDS CLARIFICATION: Error rate on command parsing]
- [NEEDS CLARIFICATION: What is the config file format and location precedence?]
- [NEEDS CLARIFICATION: Should there be a global `yg doctor` command?]
- [NEEDS CLARIFICATION: What logging format is required?]

## features/FEAT-002-project-repo-management.md
- **Performance**: [NEEDS CLARIFICATION: Max repo registration time]
- [NEEDS CLARIFICATION: Project registration success rate]
- [NEEDS CLARIFICATION: Average workspace init time]
- [NEEDS CLARIFICATION: What is the canonical project identifier format?]
- [NEEDS CLARIFICATION: Should repo registration auto-clone by default?]
- [NEEDS CLARIFICATION: How are remotes stored and validated?]

## features/FEAT-003-merge-queue-management.md
- **Performance**: [NEEDS CLARIFICATION: Max queue operation latency]
- [NEEDS CLARIFICATION: Queue processing success rate]
- [NEEDS CLARIFICATION: Avg time from add to completion]
- [NEEDS CLARIFICATION: What defines a "change" reference? PR, branch, commit?]
- [NEEDS CLARIFICATION: Should queue entries auto-expire?]
- [NEEDS CLARIFICATION: What metadata is required for each entry?]

## features/FEAT-004-task-delegation.md
- **Performance**: [NEEDS CLARIFICATION: Max task creation latency]
- [NEEDS CLARIFICATION: Delegation success rate]
- [NEEDS CLARIFICATION: Avg time to assign a task]
- [NEEDS CLARIFICATION: What is the minimal task descriptor format?]
- [NEEDS CLARIFICATION: How should dependencies between sub-tasks be expressed?]
- [NEEDS CLARIFICATION: Should tasks be versioned?]

## features/FEAT-005-task-state-management.md
- **Performance**: [NEEDS CLARIFICATION: Task state read/write latency targets]
- [NEEDS CLARIFICATION: Task state consistency checks pass rate]
- [NEEDS CLARIFICATION: Time to query task status]
- [NEEDS CLARIFICATION: What are the task state consistency checks required for MVP?]
- [NEEDS CLARIFICATION: Should task state updates be event-sourced only or allow direct mutation?]

## features/FEAT-006-agent-session-management.md
- **Performance**: [NEEDS CLARIFICATION: Startup and command latency targets]
- [NEEDS CLARIFICATION: Max acceptable session start time]
- [NEEDS CLARIFICATION: Maximum acceptable session failure rate]
- [NEEDS CLARIFICATION: Maximum concurrent sessions supported]
- [NEEDS CLARIFICATION: What telemetry fields are required per session beyond usage metrics?]
- [NEEDS CLARIFICATION: What control surfaces are available per agent (CLI, TUI, APIs)?]
- [NEEDS CLARIFICATION: Do we need to wrap agents or can we attach/control them without wrapping?]
- [NEEDS CLARIFICATION: How do we preserve subscription benefits (Claude Max, ChatGPT Pro, etc.)?]
- [NEEDS CLARIFICATION: What default sampling interval is acceptable for MVP?]

## features/FEAT-008-workspace-worktree-management.md
- **Performance**: [NEEDS CLARIFICATION: Max acceptable workspace creation time]
- [NEEDS CLARIFICATION: Workspace creation success rate]
- [NEEDS CLARIFICATION: Average workspace creation time]
- [NEEDS CLARIFICATION: What metadata fields are mandatory for workspace records?]
- [NEEDS CLARIFICATION: What are the default cleanup policies?]
- [NEEDS CLARIFICATION: Should workspace reuse be time-bounded?]

## features/FEAT-009-cross-agent-communication.md
- [NEEDS CLARIFICATION: Maximum acceptable message delivery latency]
- [NEEDS CLARIFICATION: Max message backlog size]
- [NEEDS CLARIFICATION: What are the minimal message formats required?]
- [NEEDS CLARIFICATION: How are message delivery failures surfaced?]
- [NEEDS CLARIFICATION: Should message history be stored per task or globally?]

