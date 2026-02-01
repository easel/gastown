# Feature Specification: FEAT-001 - Core CLI & Local Runtime

**Feature ID**: FEAT-001
**Status**: Draft
**Priority**: P0
**Owner**: Yggdrasil Team
**Created**: 2026-02-01
**Updated**: 2026-02-01

## Overview
Provide the foundational CLI and local runtime that all orchestration features build on. This includes command routing, configuration loading, lifecycle management, and a consistent execution environment for subdomains.

## Problem Statement

Without a stable core CLI/runtime, all orchestration features become inconsistent and hard to test. The system needs a single-binary CLI that initializes, routes, and executes commands reliably and predictably.

## Requirements

### Functional Requirements
- Provide a single-binary CLI entrypoint (`yg`)
- Load YAML configuration from standard locations with clear precedence, cascading from `$HOME/.config/yggdrasil/config.yaml`
- Route commands to domain modules (project/repo, merge queue, delegation, task/state, agents)
- Provide OpenTelemetry logging with consistent error reporting
- Support version reporting and diagnostics
- Support dry-run or safe modes where applicable

### Non-Functional Requirements
- **Performance**: CLI startup under 100 ms
- **Security**: No secrets logged by default
- **Usability**: Clear help output and command discovery; default human-readable logs written to a suitable XDG location
- **Reliability**: Deterministic command execution

## User Stories

### Story US-001: Invoke CLI [FEAT-001]
**As a** local operator
**I want** to run the `yg` CLI
**So that** I can access orchestration features

**Acceptance Criteria:**
- [ ] Given a CLI invocation, the binary starts successfully
- [ ] Given `--help`, I can discover available commands

### Story US-002: Load Configuration [FEAT-001]
**As a** local operator
**I want** configuration to load predictably
**So that** I can control runtime behavior

**Acceptance Criteria:**
- [ ] Given config files, precedence is documented and enforced (CLI flags > env vars > project config > `$HOME/.config/yggdrasil/config.yaml` > defaults)
- [ ] Given invalid config, errors are clear and actionable

### Story US-003: Route Commands [FEAT-001]
**As a** local operator
**I want** commands routed to domain modules
**So that** feature areas behave consistently

**Acceptance Criteria:**
- [ ] Given a command, it is routed to the correct domain module
- [ ] Given an unknown command, a clear error is shown

### Story US-004: Diagnostics [FEAT-001]
**As a** local operator
**I want** diagnostic output
**So that** I can debug runtime issues

**Acceptance Criteria:**
- [ ] Given `yg version`, I see version/build info
- [ ] Given `yg doctor`, I see comprehensive checks for all external dependencies
- [ ] Given `yg doctor`, every check has a test from day 1

## Edge Cases and Error Handling
- Missing config files
- Invalid configuration
- Unknown command
- Missing dependencies (e.g., git)

## Success Metrics
- CLI startup time target: 100 ms
- Parsing error rate for valid commands: 0% (invalid input should return a clear error)

## Constraints and Assumptions

### Constraints
- Single binary distribution
- Local-only runtime

### Assumptions
- Go runtime availability on build machines
- Operators are CLI-proficient

## Dependencies
- None (foundational feature)

## Out of Scope
- GUI or web UI
- Remote execution environment

## Open Questions
None

## Traceability

### Related Artifacts
- **Parent PRD Section**: Requirements Overview (P0)
- **Feature Registry**: `docs/helix/01-frame/feature-registry.md`

---
*Note: This specification avoids implementation details. Technical design is deferred to Design phase.*
