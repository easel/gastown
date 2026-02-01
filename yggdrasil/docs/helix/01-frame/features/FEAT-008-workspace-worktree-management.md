# Feature Specification: FEAT-008 - Workspace and Worktree Management

**Feature ID**: FEAT-008
**Status**: Draft
**Priority**: P0
**Owner**: Yggdrasil Team
**Created**: 2026-02-01
**Updated**: 2026-02-01

## Overview
Provide consistent, local-only workspace and worktree management for task isolation and reproducibility. This feature defines how tasks are provisioned with clean working contexts and how those contexts are listed, reused, and cleaned up.

## Problem Statement

Orchestrated tasks need isolated working contexts to avoid cross-task contamination, minimize merge conflicts, and support reproducible agent runs. Without a consistent workspace model, task delegation and agent sessions become brittle and hard to test.

## Requirements

### Functional Requirements
- Create a workspace for a task with a unique identifier
- Create an isolated worktree for a task within a workspace
- List existing workspaces and worktrees with associated task references
- Attach an agent session to a workspace/worktree
- Reuse a workspace/worktree when explicitly requested
- Clean up workspaces/worktrees on demand
- Record workspace/worktree metadata for traceability
- Workspace records must include a name and upstream repository reference

### Non-Functional Requirements
- **Security**: Workspaces are local to the user and not shared across accounts
- **Reliability**: Workspace creation is atomic and recoverable on failure
- **Usability**: CLI output clearly shows workspace/task mapping

## User Stories

### Story US-001: Create Workspace [FEAT-008]
**As a** local operator
**I want** to create a workspace for a task
**So that** I can run agents in an isolated context

**Acceptance Criteria:**
- [ ] Given a task identifier, when I create a workspace, then a unique workspace ID is returned
- [ ] Given a created workspace, metadata links it to the task

### Story US-002: Create Worktree [FEAT-008]
**As a** local operator
**I want** to create a worktree within a workspace
**So that** task work is isolated from other tasks

**Acceptance Criteria:**
- [ ] Given a workspace, when I create a worktree, then a unique worktree ID is returned
- [ ] Given a worktree, it is linked to a workspace and task

### Story US-003: List Workspaces and Worktrees [FEAT-008]
**As a** local operator
**I want** to list workspaces and worktrees
**So that** I can see current task contexts

**Acceptance Criteria:**
- [ ] Given existing workspaces, list shows workspace IDs and task references
- [ ] Given existing worktrees, list shows worktree IDs and workspace references

### Story US-004: Attach Session to Workspace [FEAT-008]
**As a** local operator
**I want** to attach an agent session to a workspace/worktree
**So that** the agent operates in the correct context

**Acceptance Criteria:**
- [ ] Given a workspace/worktree, an agent session can be launched with that context
- [ ] Given a running session, the workspace/worktree link is recorded

### Story US-005: Reuse Workspace [FEAT-008]
**As a** local operator
**I want** to reuse a workspace/worktree when appropriate
**So that** I can continue existing work without re-provisioning

**Acceptance Criteria:**
- [ ] Given a workspace, reuse requires explicit operator intent
- [ ] Given a reuse request, metadata records the reuse event

### Story US-006: Cleanup Workspace [FEAT-008]
**As a** local operator
**I want** to clean up workspaces/worktrees
**So that** local resources are reclaimed

**Acceptance Criteria:**
- [ ] Given a workspace, cleanup removes or archives the workspace
- [ ] Given a workspace with active sessions, cleanup requires explicit confirmation

## Edge Cases and Error Handling
- Workspace creation fails mid-process
- Worktree creation requested for unknown workspace
- Cleanup requested while an agent session is active
- Duplicate workspace identifiers

## Success Metrics
None

## Constraints and Assumptions

### Constraints
- Local-only workspaces for MVP
- No multi-machine workspace management
- No default cleanup policies; cleanup is explicit only
- Workspace reuse is allowed without time bounds until explicitly cleaned up

### Assumptions
- Tasks reference a single primary workspace at a time
- Workspace metadata can be stored locally

## Dependencies
- Core CLI runtime (FEAT-001)
- Project/Repo management (FEAT-002)
- Task delegation (FEAT-004)
- Task/state management (FEAT-005)
- Agent session management (FEAT-006)

## Out of Scope
- Distributed or shared workspaces
- Automatic garbage collection without operator intent

## Open Questions
None

## Traceability

### Related Artifacts
- **Parent PRD Section**: Requirements Overview (P0)
- **Feature Registry**: `docs/helix/01-frame/feature-registry.md`

---
*Note: This specification avoids implementation details. Technical design is deferred to Design phase.*
