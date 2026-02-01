# Product Requirements Document: Yggdrasil

**Version**: 0.1.0
**Date**: 2026-02-01
**Status**: Draft
**Author**: Team Yggdrasil

## Executive Summary

Yggdrasil is a CLI-first, single-binary orchestrator for multi-agent development workflows. It unifies project and repository management, merge queue coordination, task delegation, task/state tracking, and agent session management into one coherent tool. The initial release focuses on single-machine operation with clear modular boundaries, while preserving a path toward multi-machine orchestration later.

The product targets engineers and AI-augmented teams who need a reliable, test-driven orchestration layer for complex workflows. Success is defined by fast local onboarding, predictable task execution, and strong test coverage that drives implementation quality.

## Problem Statement

### The Problem
Teams coordinating AI agents, repositories, and task state must stitch together multiple tools with inconsistent interfaces and workflows. This leads to:
- Fragmented operations across project management, merge queues, and task delegation
- Poor traceability from task intent to execution and outcomes
- Increased setup time and operational overhead
- Limited testability and inconsistent automation

### Current State
Existing tools and scripts (e.g., separate systems for repo management, merge queues, and agent sessions) are powerful but disjointed. Operators often rely on manual steps and ad-hoc conventions, making workflows brittle and hard to scale.

### Opportunity
A cohesive, test-driven CLI that unifies these domains can reduce operational friction, improve reliability, and create a single source of truth for task and agent orchestration. Starting with a single-machine MVP provides immediate value and builds a stable base for eventual multi-machine orchestration.

## Goals and Objectives

### Business Goals
1. Reduce time-to-start for orchestrated projects
2. Provide a unified, reliable orchestration experience for developers
3. Ensure high-quality, test-driven implementation from day one

### Success Metrics
| Metric | Target | Measurement Method | Timeline |
|--------|--------|-------------------|----------|
| Time to initialize a project workspace | < 5 minutes | Manual benchmark | MVP |
| Time to run a full orchestrated task flow locally | < 2 minutes | Automated benchmark | MVP |
| End-to-end test coverage of core workflows | >= 80% | CI test reports | MVP |
| Operator setup errors | < 1 per 10 runs | CLI telemetry/log review | MVP |

### Non-Goals
- Multi-machine orchestration in the initial release
- GUI or web UI in the initial release
- Large-scale distributed scheduling and autoscaling

## Users and Personas

### Primary Persona: Orchestration Engineer
**Role**: Developer/Operator managing AI-assisted workflows
**Goals**:
- Run reliable task pipelines locally
- Coordinate agents and tasks without manual glue
- Keep orchestration state consistent and observable

### Secondary Persona: Release/Integration Manager
**Role**: Engineer responsible for merges and release readiness
**Goals**:
- Manage merge queues and validate integration readiness
- Track task state and completion signals
- Reduce merge conflicts and late integration surprises

## Requirements Overview

### Must Have (P0)
1. **CLI-first workflow** for all orchestration operations
2. **Single-binary distribution** with embedded or on-demand resources
3. **Project and repository management** (rigs equivalent)
4. **Merge queue management** (refinery equivalent)
5. **Task delegation and breakdown** (dun integration)
6. **Task/state management** aligned with dun + helix concepts (definition deferred until core session and workspace behaviors are validated)
7. **Agent session management** (crew/polecats equivalent) with durable session control and usage tracking
8. **Cross-agent communication** for task coordination and handoffs
9. **Workspace and worktree management** for task isolation and reproducibility
10. **Modular architecture** with clear, testable interfaces
11. **Test-first development** with end-to-end flows executed early

### Should Have (P1)
1. Pluggable policy system for routing and task assignment
2. Observability for task flows (logs, summaries, metrics)
3. Expandable integration surface for future multi-machine orchestration

### Nice to Have (P2)
1. Remote orchestration (multi-machine)
2. UI or dashboard surfaces
3. Advanced scheduling policies and workload optimization

## User Journey

### Primary Flow
1. **Entry Point**: Operator runs `yg` to initialize a project workspace
2. **First Action**: Configure a rig and register repositories
3. **Core Loop**: Delegate tasks, manage merge queue, run agents, track state
4. **Success State**: Tasks complete with traceable outputs and merged changes
5. **Exit**: Operator reviews logs and closes tasks

### Alternative Flows
- Recover from failed or interrupted agent sessions
- Pause or re-order merge queue tasks
- Re-run tasks with updated requirements

## Constraints and Assumptions

### Constraints
- **Technical**: Go implementation, single-binary CLI
- **Business**: Deliver single-machine MVP before scaling out
- **User**: CLI competence and willingness to run local workflows
 - **Operational**: Must support orchestrating multiple agent CLIs on a single machine

### Assumptions
- Operators are comfortable with terminal-based workflows
- Projects have git repositories accessible locally
- Tasks can be expressed in a standard format by dun/helix (format to be defined)

### Dependencies
- Integration with dun task definitions and helix artifacts
- Git-based repositories for rigs and merge queues

## Risks and Mitigation

| Risk | Probability | Impact | Mitigation Strategy |
|------|------------|--------|-------------------|
| Over-scoping the MVP | Medium | High | Strict P0/P1 boundaries; defer multi-machine work |
| Inconsistent task/state semantics | Medium | High | Align early with dun + helix definitions |
| Agent session instability | Medium | Medium | Standardize session lifecycle and recovery workflows |
| Test suite too slow | Medium | Medium | Optimize E2E tests and limit scope to core flows |
| Domain naming confusion | Medium | Medium | Define domains by use case; defer final names until consensus |
| Agent control surface mismatch | Medium | High | Run spikes per agent to confirm supported controls and monitoring |

## Timeline and Milestones

### Phase 1: Single-Machine MVP
- Core CLI scaffold and configuration
- Project/repo management
- Task/state management baseline
- Agent session control

### Phase 2: Workflow Integration
- Merge queue management
- Task delegation orchestration
- End-to-end tests for full workflows

### Phase 3: Hardening and Extensibility
- Observability improvements
- Policy hooks and extension points
- Preparation for future multi-machine support

## Success Criteria

### Definition of Done
- [ ] All P0 requirements implemented
- [ ] End-to-end tests cover core workflows
- [ ] Single-binary distribution built and documented
- [ ] Stakeholder review and approval

### Launch Criteria
- [ ] Local orchestration flows are stable and repeatable
- [ ] Error recovery is documented and testable
- [ ] Task/state tracking is consistent and observable

## Appendices

### A. Reference Concepts
- Project/repo management (rigs)
- Merge queue management (refinery)
- Task delegation and breakdown (dun)
- Task/state management (beads/tasks)
- Agent session management (crew/polecats)
 - Domain naming: use case first; final names deferred

---
*This PRD is a living document and will be updated as we learn more.*
