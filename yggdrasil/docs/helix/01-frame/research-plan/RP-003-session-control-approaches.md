# Research Plan: Session Control Approaches (Wrap vs Attach)

**Research Lead**: Yggdrasil Team
**Time Budget**: 1 day
**Created**: 2026-02-01
**Status**: In Progress

## Research Objectives

### Primary Research Questions

1. **Do we need to wrap agent CLIs to control them?**
   - **Why Important**: Wrapping may break subscription entitlements or TUI behavior.
   - **Success Criteria**: Clear recommendation on wrap vs attach per agent.

2. **What control options exist without wrapping?**
   - **Why Important**: Attaching to PTY/tmux may preserve official CLI behavior.
   - **Success Criteria**: Feasibility validated for attach/detach and input injection.

3. **What are the operational tradeoffs?**
   - **Why Important**: We need stable, testable automation without losing UX.
   - **Success Criteria**: Documented tradeoffs (reliability, portability, complexity).

### Knowledge Gaps
- TUI stability when managed via PTY/tmux
- Agent behavior when stdin is piped
- Compatibility with subscription login flows

## Research Scope

### In Scope
- PTY/tmux/session management options
- Wrap vs attach tradeoffs

### Out of Scope
- API-based integrations
- Remote orchestration controls

### Assumptions
1. PTY or tmux can be installed on target systems
2. CLI tools tolerate IO redirection

## Research Methods

### Method 1: PTY/Tmux Control Spike
- **Objective**: Validate attach/detach and input injection
- **Approach**: Launch agents in tmux or PTY, attach and send commands
- **Duration**: 0.5 day
- **Deliverable**: Feasibility report

### Method 2: Wrapper Spike
- **Objective**: Test basic wrapper viability
- **Approach**: Run agents with stdin/stdout capture and measure behavior
- **Duration**: 0.5 day
- **Deliverable**: Wrapper feasibility notes

## Success Criteria
- [ ] A recommended control strategy per agent
- [ ] Known failure modes documented
- [ ] Subscription compatibility assessed

---
**Next Steps**: Run PTY and wrapper experiments after CLI inventory.
