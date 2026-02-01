# Feature Specification: FEAT-009 - Cross-Agent Communication

**Feature ID**: FEAT-009
**Status**: Draft
**Priority**: P0
**Owner**: Yggdrasil Team
**Created**: 2026-02-01
**Updated**: 2026-02-01

## Overview
Enable agents to exchange messages and handoff signals during task execution. Provide a minimal, local-only message bus that supports agent-to-agent coordination and task state updates.

## Problem Statement

When multiple agents collaborate on a task, they need a reliable way to pass context, status, and handoff signals. Without cross-agent communication, coordination becomes manual and brittle.

## Requirements

### Functional Requirements
- Send a message from one agent session to another
- Support broadcast messages to all agents in a task scope
- Provide message history per task/session
- Record message metadata (sender, recipient, timestamps)
- Support handoff signals to indicate task ownership changes
- Allow operator to view or replay message streams

### Non-Functional Requirements
- **Performance**: Low-latency local message delivery
- **Security**: Messages remain local and accessible only to the operator
- **Reliability**: No message loss within a local session
- **Usability**: Clear CLI commands for send/list/replay

## User Stories

### Story US-001: Send Message [FEAT-009]
**As a** local operator
**I want** to send a message from one agent to another
**So that** agents can coordinate work

**Acceptance Criteria:**
- [ ] Given two sessions, when I send a message, then the recipient session receives it
- [ ] Given a message send, metadata records sender, recipient, and timestamp

### Story US-002: Broadcast Message [FEAT-009]
**As a** local operator
**I want** to broadcast a message to all agents in a task scope
**So that** I can coordinate group actions

**Acceptance Criteria:**
- [ ] Given multiple sessions in a task, broadcast delivers to each session
- [ ] Broadcast messages are recorded once per recipient

### Story US-003: View Message History [FEAT-009]
**As a** local operator
**I want** to view message history for a task or session
**So that** I can audit coordination and handoffs

**Acceptance Criteria:**
- [ ] Given a task/session, list shows ordered messages with metadata
- [ ] Given a message history, it is retrievable after session exit

### Story US-004: Handoff Signal [FEAT-009]
**As a** local operator
**I want** to send a handoff signal
**So that** task ownership changes are explicit

**Acceptance Criteria:**
- [ ] Given a handoff signal, it is recorded in the message stream
- [ ] Given a handoff signal, the target session is notified

## Edge Cases and Error Handling
- Sending to a non-existent session
- Broadcasting with zero active sessions
- Message history storage failure

## Success Metrics
- [NEEDS CLARIFICATION: Maximum acceptable message delivery latency]
- [NEEDS CLARIFICATION: Max message backlog size]

## Constraints and Assumptions

### Constraints
- Local-only message bus for MVP
- No external message brokers

### Assumptions
- Agent sessions can receive messages through their control surfaces
- Operators can query message history locally

## Dependencies
- Core CLI runtime (FEAT-001)
- Agent session management (FEAT-006)
- Task/state management (FEAT-005)

## Out of Scope
- Distributed message routing
- External service integration

## Open Questions
1. [NEEDS CLARIFICATION: What are the minimal message formats required?]
2. [NEEDS CLARIFICATION: How are message delivery failures surfaced?]
3. [NEEDS CLARIFICATION: Should message history be stored per task or globally?]

## Traceability

### Related Artifacts
- **Parent PRD Section**: Requirements Overview (P0)
- **Feature Registry**: `docs/helix/01-frame/feature-registry.md`

---
*Note: This specification focuses on local-only communication. Technical design is deferred to Design phase.*
