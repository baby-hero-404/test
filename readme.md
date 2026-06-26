# Prompt Base - Test Repository (Go Algorithm Implementation)

This repository contains a flat, minimalist Go project structure focused on algorithmic implementation, managed by the Prompt Base AI Orchestrator. It follows the **Librarian Protocol** for autonomous, highly structured development workflows.

## Table of Contents
- [Overview](#overview)
- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
- [Development Workflow](#development-workflow)
- [Testing](#testing)
- [License](#license)

## Overview

This project is a standardized workspace designed to run and verify code generation tasks, automated planning, and system-level refactoring, specifically tailored for Go-based algorithmic implementations. The architecture centers around a root-level solution package with auxiliary utility sub-packages.

## Project Structure

```text
.
├── code/
│   └── repos/
│       └── test/
│           └── main/
│               ├── solution.go     # Core algorithmic solution
│               ├── solution_test.go # Solution unit tests
│               ├── go.mod          # Go module specification
│               ├── readme.md       # Project documentation (this file)
│               └── numbers/        # Auxiliary logic or helper functions for numeric processing
├── openspec/                       # Technical specifications and proposals
└── task.json                       # Current agent task context and state
```

## Getting Started

### Prerequisites

- Go (v1.18 or higher recommended)

### Setup

Clone the repository and inspect the tasks:

```bash
git clone <repository-url>
cd code/repos/test/main
```

## Development Workflow

We follow the **Librarian Protocol** which enforces:
1. **Discovery**: Look up registered skills in `registry.min.json`.
2. **Socratic Gate**: Engage in clarification rounds for ambiguous requirements.
3. **Execution & Verification**: Implement surgical, targeted patches and run automated validations.

### Writing Code
All code edits should be surgical, preserving original code style and existing comments. Always add tests for new features.

## Testing

To run the automated tests within this workspace, run:

```bash
# Execute Go test suite
go test ./...
```

## License

This project is licensed under the MIT License.
