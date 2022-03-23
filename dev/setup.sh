#!/usr/bin/env bash

install_hooks() {
  git config core.hooksPath \
    || (git config core.hooksPath "./dev/git_hooks")
}

echo "Setting up Git hooks"
install_hooks
