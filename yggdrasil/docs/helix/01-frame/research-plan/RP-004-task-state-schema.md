# Research Plan: Task/State Schema Discovery (FEAT-005)

**Research Lead**: Yggdrasil Team
**Time Budget**: 2-3 days
**Created**: 2026-02-01
**Status**: Draft

## Research Objectives

### Primary Research Questions

1. **What is the minimal task/state schema needed for MVP?**
   - **Why Important**: FEAT-005 is deferred until core session/workspace/comm behaviors are understood.
   - **Success Criteria**: A minimal schema that supports delegation, ownership, state transitions, and telemetry links.

2. **How should tasks map to HELIX artifacts and phases?**
   - **Why Important**: Task lifecycle must align with HELIX workflow outputs.
   - **Success Criteria**: Clear mapping between task states and HELIX phase gates.

3. **How should task state link to agent sessions and workspaces?**
   - **Why Important**: Task state must be traceable to session outputs and workspace context.
   - **Success Criteria**: Defined link fields (task_id -> session_id/workspace_id).

### Knowledge Gaps
- Minimal schema vs. future extensibility
- State transition rules
- Required metadata for audits and summaries

## Research Scope

### In Scope
- Local-only schema for MVP
- Mapping to delegation and session states

### Out of Scope
- Distributed task/state storage
- Multi-tenant access controls

### Assumptions
1. Session/workspace metadata is available from FEAT-006 and FEAT-008
2. HELIX artifacts will provide anchors for task definitions

## Research Methods

### Method 1: Schema Drafting Workshop
- **Objective**: Define minimal field set and states
- **Approach**: Draft schema proposals and review against use cases
- **Duration**: 1 day
- **Deliverable**: Schema draft v1

### Method 2: Traceability Mapping
- **Objective**: Map tasks to HELIX artifacts and phases
- **Approach**: Compare PRD/user stories/feature specs to required task fields
- **Duration**: 1 day
- **Deliverable**: Mapping document

### Method 3: Integration Alignment
- **Objective**: Align schema with sessions/workspaces
- **Approach**: Identify required foreign keys and references
- **Duration**: 0.5 day
- **Deliverable**: Linking proposal

## Success Criteria
- [ ] Minimal schema defined
- [ ] State transition model documented
- [ ] Links to sessions/workspaces defined
- [ ] Mapping to HELIX artifacts completed

## Expected Outcomes

### Impact on Product Development
- **Feature Specs**: FEAT-005 updated with concrete schema
- **Design Phase**: Input for contracts and storage design
- **Testing**: Enables task/state contract tests

---
**Next Steps**: Begin schema drafting workshop.
