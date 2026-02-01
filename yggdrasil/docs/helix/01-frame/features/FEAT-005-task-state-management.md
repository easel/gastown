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

### Non-Functional Requirements
- **Performance**: [NEEDS CLARIFICATION: Task state read/write latency targets]
- **Security**: Access limited to local user; no secret leakage in task logs
- **Reliability**: Atomic state updates and recovery on crash

## Schema Reference

See `docs/helix/01-frame/research-plan/findings/task-state-schema-v0.md` (v1 draft) for the current proposed fields, transitions, and links.

## Edge Cases and Error Handling
- Partial task updates and conflicting state transitions
- Corrupted or missing task state data
- Tasks with missing agent session context

## Success Metrics
- [NEEDS CLARIFICATION: Task state consistency checks pass rate]
- [NEEDS CLARIFICATION: Time to query task status]

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

## Open Questions
1. [NEEDS CLARIFICATION: What is the minimal task state schema required for MVP?]
2. [NEEDS CLARIFICATION: How are task state updates synchronized with agent outputs?]
3. [NEEDS CLARIFICATION: What is the canonical mapping from HELIX artifacts to task state?]

## Traceability

### Related Artifacts
- **Parent PRD Section**: Requirements Overview (P0)
- **Feature Registry**: `docs/helix/01-frame/feature-registry.md`
- **Schema Draft**: `docs/helix/01-frame/research-plan/findings/task-state-schema-v0.md`

---
*Note: Task/state schema design is deferred until core session and workspace behaviors are validated.*
