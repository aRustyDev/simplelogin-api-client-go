#!/usr/bin/env bash

# Install Go Dependencies
if ! command -v go &> /dev/null; then
    echo "Install 'Go/Golang' first"
    exit 1
else
    go get -u golang.org/x/lint/golint
    go install -v github.com/go-critic/go-critic/cmd/gocritic@latest
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.59.1
    go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
    go install golang.org/x/tools/cmd/goimports@latest
fi

# Install pre-commit hooks
if ! command -v pip &> /dev/null; then
    echo "Install 'pip' first"
    exit 1
else
    pip install pre-commit
    pre-commit install --install-hooks
fi

# Install Docker
if ! command -v docker &> /dev/null; then
    echo "Install 'docker' first"
fi

# Install one-password-cli
if ! command -v op &> /dev/null; then
    echo "(Optional) Install '1password-cli@beta' first; otherwise handle secrets manually"
fi
