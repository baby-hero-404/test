# Prompt Base - Test Repository (Go Algorithm Implementation)

This repository contains a flat, minimalist Go project structure focused on algorithmic implementation, managed by the Prompt Base AI Orchestrator. It adheres strictly to the **Librarian Protocol** for autonomous, highly structured development workflows.

## Table of Contents

- [Overview](#overview)
- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
- [Development Workflow](#development-workflow)
- [Testing](#testing)
- [License](#license)

## Overview

This project provides a standardized workspace designed to run and verify code-generation tasks, automated planning, and system-level refactoring. It is specifically tailored for Go-based algorithmic implementations, centering around a root-level solution package with auxiliary utility sub-packages.

## Project Structure

```text
.
├── code/
│   └── repos/
│       └── test/
│           └── main/
│               ├── solution.go      # Core algorithmic solution
│               ├── solution_test.go # Solution unit tests
│               ├── go.mod           # Go module specification
│               ├── readme.md        # Project documentation (this file)
│               └── numbers/         # Auxiliary logic or helper functions for numeric processing
├── openspec/                        # Technical specifications and proposals
└── task.json                        # Current agent task context and state
```

## Getting Started

### Prerequisites

- **Go Compiler**: Version 1.18 or higher is recommended.

### Setup

Clone the repository and inspect the tasks:

```bash
git clone <repository-url>
cd code/repos/test/main
```

## Development Workflow

Development follows the **Librarian Protocol**, which enforces three primary phases:

1. **Discovery**: Look up registered skills in `registry.min.json`.
2. **Socratic Gate**: Engage in clarification rounds for ambiguous requirements.
3. **Execution & Verification**: Implement surgical, targeted patches and run automated validations.

### Writing Code

All code edits must be surgical, preserving the original code style and existing comments. Always ensure comprehensive unit tests are added for new features.

### Security Guidelines

Security is a primary concern. Ensure your contributions follow these principles:
- **No Hardcoded Secrets**: Do not check in credentials, API keys, or private tokens.
- **Input Validation**: Validate and sanitize all external inputs before processing.
- **Safe Execution**: Avoid unsafe operations (e.g., `eval`, dynamic SQL execution).

## Testing

To run the automated tests within this workspace, use the following commands:

```bash
# Execute the Go test suite
go test ./...

# Run tests with verbose output
go test -v ./...
```

## License

This project is licensed under the MIT License.
