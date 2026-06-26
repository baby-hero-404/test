# Prompt Base - Test Repository

This repository contains test components, configurations, and specs managed by the Prompt Base AI Orchestrator. It follows the **Librarian Protocol** for autonomous, highly structured development workflows.

## Table of Contents
- [Overview](#overview)
- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
- [Development Workflow](#development-workflow)
- [Testing](#testing)
- [License](#license)

## Overview

This project is a standardized workspace designed to run and verify code generation tasks, automated planning, and system-level refactoring. It serves as a sandbox/target environment for prompt-driven software agents.

## Project Structure

```text
.
├── code/
│   └── repos/
│       └── test/
│           └── main/
│               └── readme.md       # Project documentation (this file)
├── openspec/                       # Technical specifications and proposals
└── task.json                       # Current agent task context and state
```

## Getting Started

### Prerequisites

- Node.js (v22 or higher recommended)
- Python (v3.12 or higher recommended)

### Setup

Clone the repository and inspect the tasks:

```bash
git clone <repository-url>
cd <repository-directory>
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
# Execute test suite
npm run test
# or
python -m pytest
```

## License

This project is licensed under the MIT License.
