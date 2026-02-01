# Task/State Schema Draft (v1)

**Date**: 2026-02-01
**Status**: Draft (refined)

## Goals
- Minimal schema to support delegation, ownership, and state transitions
- Traceability to sessions, workspaces, and HELIX artifacts
- Local-only storage for MVP

## Core Entities

### Task
- `task_id`
- `title`
- `description`
- `status`
- `priority` (P0/P1/P2)
- `owner_session_id`
- `workspace_id`
- `feature_id` (FEAT-XXX)
- `story_id` (US-XXX, optional)
- `helix_phase` (frame/design/test/build/deploy/iterate, optional)
- `created_at`
- `updated_at`
- `started_at` (optional)
- `completed_at` (optional)
- `tags` (optional)
- `notes` (optional)

### SubTask
- `subtask_id`
- `parent_task_id`
- `title`
- `status`
- `owner_session_id`
- `created_at`
- `updated_at`

### TaskDependency
- `task_id`
- `depends_on_task_id`
- `dependency_type` (required/blocks/related)

### TaskEvent (optional for MVP, recommended)
- `event_id`
- `task_id`
- `timestamp`
- `actor` (operator/session)
- `event_type` (assign/reassign/status_change/comment)
- `from_status` (optional)
- `to_status` (optional)
- `note` (optional)

## TaskStatus (enum)
- `new`
- `assigned`
- `in_progress`
- `blocked`
- `completed`
- `failed`
- `canceled`

## State Transitions (MVP)
- `new -> assigned`
- `assigned -> in_progress`
- `in_progress -> completed | failed | blocked`
- `blocked -> in_progress`
- Any -> `canceled` (explicit operator intent)

## Links
- `task_id -> session_id`
- `task_id -> workspace_id`
- `task_id -> feature_id`
- `task_id -> story_id` (optional)

## Open Questions
- Do we need a separate status for handoff vs waiting?
- How tightly should tasks map to HELIX phases?
- Should TaskEvent be mandatory in MVP or optional?
