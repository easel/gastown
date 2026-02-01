# Feature Registry

**Document Type**: Feature Registry
**Status**: Active
**Last Updated**: 2026-02-01
**Maintained By**: Yggdrasil Team

## Purpose

This registry tracks all features in Yggdrasil, their status, dependencies, and ownership.

## Active Features

| ID | Name | Description | Status | Priority | Owner | Created | Updated |
|----|------|-------------|--------|----------|-------|---------|---------|
| FEAT-001 | Core CLI & Local Runtime | Single-binary CLI, config, and local orchestration shell | Draft | P0 | Team | 2026-02-01 | 2026-02-01 |
| FEAT-002 | Project/Repo Management (Rigs) | Manage projects and repositories | Draft | P0 | Team | 2026-02-01 | 2026-02-01 |
| FEAT-003 | Merge Queue Management (Refinery) | Coordinate merge queues and integration readiness | Draft | P0 | Team | 2026-02-01 | 2026-02-01 |
| FEAT-004 | Task Delegation & Breakdown (Dun) | Delegate tasks and break work into units | Draft | P0 | Team | 2026-02-01 | 2026-02-01 |
| FEAT-005 | Task/State Management | Track task state aligned with dun + helix (schema deferred) | Draft | P0 | Team | 2026-02-01 | 2026-02-01 |
| FEAT-006 | Agent Session Management | Manage agent sessions (crew/polecats) | Draft | P0 | Team | 2026-02-01 | 2026-02-01 |
| FEAT-007 | Observability & Reporting | Logs, summaries, and metrics for orchestration flows | Draft | P1 | Team | 2026-02-01 | 2026-02-01 |
| FEAT-008 | Workspace & Worktree Management | Isolated workspaces and worktrees for tasks | Draft | P0 | Team | 2026-02-01 | 2026-02-01 |
| FEAT-009 | Cross-Agent Communication | Agent-to-agent messaging and handoff signals | Draft | P0 | Team | 2026-02-01 | 2026-02-01 |
| FEAT-010 | Centralized Mobile Access | Mobile-friendly access to session chat and status | Draft | P2 | Team | 2026-02-01 | 2026-02-01 |

## Feature Dependencies

| Feature | Depends On | Dependency Type | Notes |
|---------|------------|-----------------|-------|
| FEAT-002 | FEAT-001 | Required | Core CLI runtime needed first |
| FEAT-003 | FEAT-001, FEAT-002 | Required | Merge queue tied to projects and repos |
| FEAT-004 | FEAT-001 | Required | Delegation needs runtime orchestration |
| FEAT-005 | FEAT-001, FEAT-004, FEAT-008 | Required | Task state tied to delegation and workspace context |
| FEAT-006 | FEAT-001 | Required | Session control tied to runtime |
| FEAT-007 | FEAT-001..FEAT-009 | Optional | Observability spans features |
| FEAT-008 | FEAT-001, FEAT-002 | Required | Worktrees depend on repo management |
| FEAT-009 | FEAT-001, FEAT-006 | Required | Communication depends on runtime and session control |

## Feature Categories

### Orchestration Core
- FEAT-001: Core CLI & Local Runtime
- FEAT-004: Task Delegation & Breakdown
- FEAT-005: Task/State Management
- FEAT-008: Workspace & Worktree Management
- FEAT-009: Cross-Agent Communication

### Repo and Merge Management
- FEAT-002: Project/Repo Management (Rigs)
- FEAT-003: Merge Queue Management (Refinery)

### Agent Operations
- FEAT-006: Agent Session Management

### Observability
- FEAT-007: Observability & Reporting

### Interfaces (Future)
- FEAT-010: Centralized Mobile Access

## ID Assignment Rules

1. Features are numbered sequentially (FEAT-001, FEAT-002, ...)
2. IDs are permanent and never reused
3. New features are appended at the end

## Deprecated/Cancelled Features

| ID | Name | Status | Reason | Date |
|----|------|--------|--------|------|
| | | | | |

## Cross-References

- **PRD**: `docs/helix/01-frame/prd.md`
- **Principles**: `docs/helix/01-frame/principles.md`
- **Feature Specs**: `docs/helix/01-frame/features/FEAT-XXX-[name].md`

---
*Update this registry whenever features are added, modified, or their status changes.*
