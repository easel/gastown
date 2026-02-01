# Task/State Schema Draft (v0)

**Date**: 2026-02-01
**Status**: Draft (pending review)

## Core Entities

### Task
- `task_id`
- `title`
- `description`
- `status`
- `owner_session_id`
- `workspace_id`
- `created_at`
- `updated_at`

### SubTask
- `subtask_id`
- `parent_task_id`
- `title`
- `status`
- `owner_session_id`

### TaskStatus (enum)
- `new`
- `assigned`
- `in_progress`
- `blocked`
- `completed`
- `failed`
- `canceled`

## Links
- `task_id -> session_id`
- `task_id -> workspace_id`
- `task_id -> feature_id`

## Open Questions
- Do we need separate status for handoff or waiting?
- How do we represent HELIX phase alignment?
- Do we need audit events for transitions?
