# Research Plan: Session Telemetry (CPU/Mem/IO/Storage/Token)

**Research Lead**: Yggdrasil Team
**Time Budget**: 1 day
**Created**: 2026-02-01
**Status**: Draft

## Research Objectives

### Primary Research Questions

1. **How do we measure per-session CPU, memory, storage, and IO?**
   - **Why Important**: Resource tracking is a P0 requirement.
   - **Success Criteria**: Identify OS-level methods to attribute usage to agent processes.

2. **How do we capture token consumption per session?**
   - **Why Important**: Token usage is a core cost/performance metric.
   - **Success Criteria**: Determine whether token usage is reported by each agent CLI and how to extract it.

3. **What telemetry is reliable across agents?**
   - **Why Important**: We need consistent UX across agent types.
   - **Success Criteria**: Define a minimal telemetry set guaranteed for MVP.

### Knowledge Gaps
- OS-level process accounting reliability
- Token usage exposure per agent CLI
- Storage attribution for session workspaces

## Research Scope

### In Scope
- Local process resource tracking
- Token usage extraction (where available)

### Out of Scope
- Distributed telemetry pipelines
- Long-term metrics storage backends

### Assumptions
1. Yggdrasil can observe process IDs for sessions
2. Resource tracking can be derived from OS tooling

## Research Methods

### Method 1: OS Telemetry Spike
- **Objective**: Validate per-process usage tracking
- **Approach**: Use `ps`, `top`, or platform tools to capture CPU/mem/IO
- **Duration**: 0.5 day
- **Deliverable**: Metrics capture approach per OS

### Method 2: Token Usage Inspection
- **Objective**: Determine how token usage is surfaced by each CLI
- **Approach**: Run minimal agent sessions and inspect outputs/logs
- **Duration**: 0.5 day
- **Deliverable**: Token usage capture notes per agent

## Success Criteria
- [ ] CPU/memory/IO tracking approach documented
- [ ] Storage usage attribution approach documented
- [ ] Token usage capture documented per agent
- [ ] Minimal telemetry set defined for MVP

---
**Next Steps**: Start OS telemetry spike after CLI inventory.
