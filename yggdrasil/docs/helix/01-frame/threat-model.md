# Threat Model

**Project**: Yggdrasil
**Version**: 0.1.0
**Date**: 2026-02-01
**Threat Modeling Team**: Yggdrasil Team
**Review Date**: 2026-03-01

---

## Executive Summary

**System Overview**: CLI-first local orchestrator managing repos, tasks, and agent sessions.
**Key Assets**: Repo contents, task/state data, credentials, logs, agent outputs.
**Primary Threats**: Secret leakage, command injection, tampering with task state, malicious agent output.
**Risk Level**: Medium (local-first, but handles sensitive data and shell operations).

## System Description

### System Boundaries
**In Scope**: Local CLI, configuration, repo operations, task state, agent sessions, logs.
**Out of Scope**: Multi-machine orchestration, external scheduling services, hosted UI.
**Trust Boundaries**: User input boundaries, external command execution, git/network operations.

### System Components
1. **CLI Runtime**: Orchestration logic and command handling
2. **Task/State Store**: Local persistence for task and execution state
3. **Repo Manager**: Git operations for rigs and merge queue
4. **Agent Session Manager**: Process/session control and IO capture

### Data Flows
- User input to CLI commands
- CLI invoking git and external tools
- Task/state updates stored locally
- Agent session output captured to logs

## Assets and Protection Goals

### Data Assets
| Asset | Classification | Confidentiality | Integrity | Availability | Owner |
|-------|---------------|-----------------|-----------|--------------|-------|
| Repo contents | Sensitive | Medium | High | High | User |
| Credentials/tokens | Highly Sensitive | Critical | High | Medium | User |
| Task/state data | Internal | Medium | High | High | User |
| Logs/outputs | Internal | Medium | Medium | Medium | User |

### System Assets
| Asset | Description | Criticality | Dependencies |
|-------|-------------|-------------|--------------|
| CLI runtime | Orchestration and command handling | High | Go binary |
| Task/state store | Local persistence | High | Filesystem |
| Repo manager | Git operations | High | git |
| Agent manager | Session lifecycle control | Medium | OS process/session tooling |

## STRIDE Threat Analysis (High-Level)

### Spoofing Identity
- Malicious inputs masquerading as trusted config
- Risk: Medium; Mitigation: strict config validation and path scoping

### Tampering with Data
- Task/state file tampering
- Risk: Medium; Mitigation: checksums and controlled write access

### Repudiation
- Lack of audit trail for repo or task updates
- Risk: Medium; Mitigation: structured local logging

### Information Disclosure
- Secrets in logs or task outputs
- Risk: High; Mitigation: redaction, no secret printing, secure storage

### Denial of Service
- Resource exhaustion via runaway agents or tasks
- Risk: Medium; Mitigation: process limits and timeouts

### Elevation of Privilege
- Command injection via CLI input
- Risk: High; Mitigation: strict argument handling and validation

## Top Risks Identified

| Risk ID | Threat | Impact | Likelihood | Priority |
|---------|--------|--------|------------|----------|
| TM-001 | Secret leakage in logs | High | Medium | High |
| TM-002 | Command injection | High | Medium | High |
| TM-003 | Task/state tampering | Medium | Medium | Medium |
| TM-004 | Malicious agent output | Medium | Medium | Medium |

## Mitigation Strategies

1. **Secrets Handling**
   - Redact secrets in logs
   - Avoid printing tokens or config secrets

2. **Safe Command Execution**
   - Use explicit argument lists
   - Validate paths and inputs

3. **State Integrity**
   - Atomic writes for task/state files
   - Optional checksums for state snapshots

4. **Session Safety**
   - Timeouts and resource limits for agent sessions
   - Explicit user control for starting/stopping sessions

## Maintenance and Updates

- Review after MVP completion
- Update for any networked or multi-machine features

## Approval and Sign-off

| Role | Name | Signature | Date |
|------|------|-----------|------|
| Security Champion | TBD | | |
| Technical Lead | TBD | | |
| Product Owner | TBD | | |

---
**Document Control**
- **Template Version**: 0.1
- **Last Updated**: 2026-02-01
- **Next Review Date**: 2026-03-01
