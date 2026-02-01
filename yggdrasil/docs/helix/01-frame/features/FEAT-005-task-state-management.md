# Feature Specification: FEAT-005 - Task/State Management

**Feature ID**: FEAT-005
**Status**: Draft
**Priority**: P0
**Owner**: Yggdrasil Team
**Created**: 2026-02-01
**Updated**: 2026-02-01

## Overview
Define and manage task and state semantics aligned with dun and HELIX. This feature establishes how tasks are represented, tracked, and updated across the orchestration lifecycle.

## Problem Statement

Without a shared task/state model, coordination across delegation, agent sessions, and merge queues becomes inconsistent and hard to test. A clear, testable task/state model is required, informed by validated session, workspace, and communication behaviors.

## Requirements

### Functional Requirements
- Define task lifecycle states and transitions aligned with orchestration workflows
- Persist task and execution state locally with auditability
- Provide queryable task status for CLI operations
- Support linking tasks to workspaces, agents, and repositories
- Support task dependencies and ownership changes
- Synchronize task state updates with agent session outputs

### Non-Functional Requirements
- **Security**: Access limited to local user; no secret leakage in task logs
- **Reliability**: Atomic state updates and recovery on crash

## Schema Reference

See `docs/helix/01-frame/research-plan/findings/task-state-schema-v0.md` (v1 draft) for the current proposed fields, transitions, and links.

## Minimal Task Schema (MVP)

**Required fields**:
- `task_id`
- `title`
- `status`
- `owner_session_id`
- `workspace_id`
- `feature_id`
- `created_at`
- `updated_at`

**Optional fields** (MVP):
- `description`
- `priority`
- `story_id`
- `helix_phase`
- `artifact_refs`
- `started_at`
- `completed_at`
- `tags`
- `notes`

## Task State Sync with Agent Outputs (MVP)

- **Source of truth**: Task record + TaskEvent history
- **Sync triggers**:
  - Session start -> set `status` to `in_progress` if assigned
  - Session exit -> set `status` to `completed` or `failed` based on exit status
  - Operator override -> allow explicit state change with audit event
- **Evidence model**:
  - Append agent output metadata to TaskEvent logs (timestamp, session_id, outcome)
  - Do not overwrite task fields without explicit state transitions

## Task State Consistency Checks (MVP)

- **Required fields present**: all required task fields populated
- **Valid transitions**: status changes must follow the allowed transition graph
- **Referential integrity**: referenced `session_id`, `workspace_id`, and `feature_id` exist
- **Timestamps monotonic**: `created_at <= updated_at <= completed_at` (when present)
- **Event audit**: any state change must produce a TaskEvent entry

## Event-Sourcing Stance (MVP)

- **Hybrid approach**: Task is a materialized view; TaskEvent is append-only audit log.
- **State changes**: MUST be recorded via TaskEvent and then reflected in Task.
- **Metadata updates**: MAY be direct mutation (notes/tags) but still emit TaskEvent.

## HELIX Artifact Mapping (Resolved)

Tasks map to HELIX artifacts and phases using explicit references and lightweight fields:
- **Feature** -> `feature_id` (FEAT-XXX)
- **User Story** -> `story_id` (US-XXX, optional)
- **HELIX Phase** -> `helix_phase` (frame/design/test/build/deploy/iterate, optional)
- **Artifacts** -> `artifact_refs` (optional list of doc paths, e.g., PRD section or spec)

This provides traceability without embedding full artifact content in task state.

## Edge Cases and Error Handling
- Partial task updates and conflicting state transitions
- Corrupted or missing task state data
- Tasks with missing agent session context

## Success Metrics
- Task status query latency target: 100 ms

## Constraints and Assumptions

### Constraints
- Local-only persistence for MVP
- Schema may evolve once FEAT-006/FEAT-008 are finalized

### Assumptions
- Dun and HELIX artifacts provide baseline structure, but need refinement

## Dependencies
- Core CLI runtime (FEAT-001)
- Task delegation (FEAT-004)
- Agent session management (FEAT-006)
- Workspace & worktree management (FEAT-008)
- Cross-agent communication (FEAT-009)

## Out of Scope
- Distributed task state across multiple machines
- Centralized or hosted task state services

## Traceability

### Related Artifacts
- **Parent PRD Section**: Requirements Overview (P0)
- **Feature Registry**: `docs/helix/01-frame/feature-registry.md`
- **Schema Draft**: `docs/helix/01-frame/research-plan/findings/task-state-schema-v0.md`

---
*Note: Task/state schema design is deferred until core session and workspace behaviors are validated.*
