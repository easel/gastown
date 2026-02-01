# Research Plan: Agent Control Surfaces (CLI/TUI)

**Research Lead**: Yggdrasil Team
**Time Budget**: 1-2 days
**Created**: 2026-02-01
**Status**: Completed

## Research Objectives

### Primary Research Questions

1. **What control surfaces are available per agent CLI?**
   - **Why Important**: Determines whether Yggdrasil can start/attach/drive sessions without a wrapper.
   - **Success Criteria**: Documented control options for Claude Code, Gemini CLI, Codex CLI, OpenCode CLI.

2. **What session signals are observable?**
   - **Why Important**: Defines which session states can be tracked reliably.
   - **Success Criteria**: For each agent, list signals for running/idle/failed/exited and how to detect them.

3. **What usage telemetry can be collected from agents?**
   - **Why Important**: We must track token usage and resource usage per session.
   - **Success Criteria**: Identify which agents expose token usage and how to capture it.

4. **How do we preserve subscription benefits (Claude Max, ChatGPT Pro, etc.)?**
   - **Why Important**: We need to leverage paid subscriptions without forcing API keys.
   - **Success Criteria**: Clear guidance per agent on whether CLI access uses subscription entitlements.

### Knowledge Gaps
- Control surfaces (stdin, flags, APIs, TUI attach)
- Session state observability
- Token usage visibility and export formats
- Subscription entitlement paths for CLI usage

## Research Scope

### In Scope
- Claude Code CLI
- Gemini CLI
- Codex CLI
- OpenCode CLI

### Out of Scope
- Custom API wrappers
- Multi-machine orchestration

### Assumptions
1. CLI tools are installable locally
2. At least one control surface exists (stdin or TUI attach)
3. Token usage may be partially visible depending on agent

## Research Methods

### Method 1: Local CLI Inspection
- **Objective**: Identify control flags, stdin behavior, and status outputs
- **Approach**: Run `--help`, `--version`, and minimal sessions
- **Participants/Sources**: Local CLI binaries
- **Duration**: 0.5 day
- **Deliverable**: Per-agent control surface matrix

### Method 2: TUI Session Attach Spike
- **Objective**: Validate attach/detach behavior via PTY/tmux
- **Approach**: Launch session in a PTY, attempt attach/detach, confirm IO
- **Participants/Sources**: Local CLI binaries, tmux/pty
- **Duration**: 0.5 day
- **Deliverable**: Attach/detach feasibility notes

### Method 3: Subscription Validation
- **Objective**: Confirm whether CLIs honor subscription entitlements
- **Approach**: Review auth/credential flows and login methods
- **Participants/Sources**: CLI docs, local auth prompts
- **Duration**: 0.5 day
- **Deliverable**: Subscription compatibility notes

## Success Criteria

### Research Completion Criteria
- [x] Control surface matrix completed for each agent
- [x] Session state observability documented per agent
- [x] Token usage exposure documented per agent
- [x] Subscription compatibility assessed per agent

## Timeline and Milestones

| Phase | Duration | Activities | Deliverables | Responsible |
|-------|----------|------------|--------------|-------------|
| Planning | 0.5 day | Finalize plan and tools | Approved plan | Team |
| Investigation | 1 day | CLI inspection + TUI attach | Control surface matrix | Team |
| Analysis | 0.5 day | Summarize findings | Findings report | Team |

## Resource Requirements

### Tools and Materials
- Local CLI installs (claude, gemini, codex, opencode)
- PTY or tmux
- Process inspection tools (ps, top)

## Expected Outcomes

### Impact on Product Development
- **PRD Impact**: Confirm agent session scope
- **Feature Specifications**: Finalize session states and control methods
- **Technical Decisions**: Decide wrapper vs attach strategies

---
**Next Steps**: See findings and incorporate into FEAT-006 decisions.

## Findings
- Control surface matrix: `docs/helix/01-frame/research-plan/findings/agent-control-surface-matrix.md`
- CLI inventory: `docs/helix/01-frame/research-plan/findings/agent-cli-inventory.md`
- TUI attach notes: `docs/helix/01-frame/research-plan/findings/tmux-spike-summary.md`
