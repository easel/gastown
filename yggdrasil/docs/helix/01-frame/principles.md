# Project Principles

*Project-specific principles for Yggdrasil. These extend the HELIX workflow principles.*

## Core Principles

1. **CLI-First**
   - The primary interface is a command-line UX.
   - All core workflows must be operable without a GUI.

2. **Local-First, Single-Machine MVP**
   - The first release must operate entirely on a single machine.
   - Multi-machine orchestration is a long-term goal and is explicitly deferred.

3. **Single-Binary Distribution**
   - Yggdrasil ships as one CLI binary.
   - Any supporting resources are embedded or cloned on demand.

4. **Go Implementation**
   - The system is written in Go.
   - Build and runtime dependencies are minimized.

5. **Modular, Encapsulated Design**
   - Modules have clear, testable interfaces.
   - Boundaries are explicit and enforced by packages and contracts.

6. **Test-First, Integration-Biased**
   - Tests drive implementation from the start.
   - Prefer real implementations over heavy mocking.
   - End-to-end tests must exercise the full system early.

7. **HELIX Methodology Compliance**
   - Work follows Frame → Design → Test → Build → Deploy → Iterate.
   - No implementation begins before tests are written and failing.

8. **Domain-First Naming**
   - Domains are defined by use case and mapped to existing gastown concepts.
   - Final naming is deferred until domain boundaries are agreed.

9. **Deferred Task/State Schema**
   - Task/state schema is defined after core session, workspace, and communication behaviors are validated.
   - Early work must avoid locking the schema prematurely.

## Technology Constraints

- **Primary Language**: Go
- **Primary Interface**: CLI
- **Distribution**: Single binary

## Exceptions Log

| Date | Principle | Exception | Justification | Resolution Timeline |
|------|-----------|-----------|---------------|-------------------|
| | | | | |
