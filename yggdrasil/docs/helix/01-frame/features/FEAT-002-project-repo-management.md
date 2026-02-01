# Feature Specification: FEAT-002 - Project/Repo Management

**Feature ID**: FEAT-002
**Status**: Draft
**Priority**: P0
**Owner**: Yggdrasil Team
**Created**: 2026-02-01
**Updated**: 2026-02-01

## Overview
Manage projects and repositories (rigs equivalent). Provide local discovery, registration, and workspace setup for repositories used by the orchestrator.

## Problem Statement

Operators need a consistent way to register, list, and prepare repositories for orchestration workflows. Without a uniform project/repo layer, downstream features (merge queue, tasks, agents) lack a stable reference point.

## Requirements

### Functional Requirements
- Register a repository as a project with a unique human-friendly identifier (slug)
- List registered projects and their repo metadata
- Initialize a project workspace on demand
- Registration should bare-clone repositories to bootstrap worktrees
- Validate repository availability and remote configuration
- Update or re-sync project metadata
- Remove or archive project entries
- Support a simple in-tree metadata format for Yggdrasil-specific project configuration (agent naming conventions, contribution policies, etc.)

### Non-Functional Requirements
- **Security**: Do not leak credentials or sensitive repo data
- **Reliability**: Repository metadata remains consistent across updates
- **Usability**: Clear CLI output for project list and status

## User Stories

### Story US-001: Register Project [FEAT-002]
**As a** local operator
**I want** to register a repository as a project
**So that** it can be orchestrated

**Acceptance Criteria:**
- [ ] Given a repo path or URL, registration creates a unique human-friendly project ID (slug)
- [ ] Given a registered project, metadata includes repo path and remotes

### Story US-002: List Projects [FEAT-002]
**As a** local operator
**I want** to list registered projects
**So that** I can see available repositories

**Acceptance Criteria:**
- [ ] Given registered projects, list shows IDs, names, and repo paths

### Story US-003: Initialize Project Workspace [FEAT-002]
**As a** local operator
**I want** to initialize a project workspace
**So that** downstream tasks can use a prepared repo

**Acceptance Criteria:**
- [ ] Given a project, initialization prepares a workspace directory
- [ ] Given a missing repo, initialization returns a clear error

### Story US-004: Update Project Metadata [FEAT-002]
**As a** local operator
**I want** to update project metadata
**So that** repo changes are reflected

**Acceptance Criteria:**
- [ ] Given a project, refresh updates remotes and in-tree metadata
- [ ] Given a failure, errors are recorded clearly

### Story US-005: Remove Project [FEAT-002]
**As a** local operator
**I want** to remove or archive a project
**So that** unused repos are cleaned up

**Acceptance Criteria:**
- [ ] Given a project, removal requires explicit confirmation
- [ ] Given removal, metadata is deleted or archived

## Edge Cases and Error Handling
- Repo path not found
- Repo missing git metadata
- Duplicate project identifiers
- Remote config mismatch

## Success Metrics
None

## Constraints and Assumptions

### Constraints
- Local-only repo management for MVP
- Registration can perform a bare clone by default to bootstrap worktrees

### Assumptions
- Operators have local git access
- Projects map to a single primary repository

## Dependencies
- Core CLI runtime (FEAT-001)
- Workspace/worktree management (FEAT-008)

## Out of Scope
- Multi-repo projects (monorepo + submodules) for MVP
- Remote repo orchestration

## Open Questions
None

## Traceability

### Related Artifacts
- **Parent PRD Section**: Requirements Overview (P0)
- **Feature Registry**: `docs/helix/01-frame/feature-registry.md`

---
*Note: This specification avoids implementation details. Technical design is deferred to Design phase.*
