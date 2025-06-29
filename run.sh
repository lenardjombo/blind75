#!/bin/bash

# Usage:
# ./run.sh test           -> Runs all Go tests
# ./run.sh push "msg"     -> Commits and pushes with message "msg"

case "$1" in
  test)
    echo "Running Go tests..."
    go test ./... -v
    ;;

  push)
    if [ -z "$2" ]; then
      echo "Commit message required. Usage: ./run.sh push \"your message\""
      exit 1
    fi
    echo "Adding all changes..."
    git add .
    echo "Committing with message: $2"
    git commit -m "$2"
    echo "Pushing to GitHub..."
    git push
    ;;

  *)
    echo "Unknown command: $1"
    echo "Usage: ./run.sh [test|push \"commit message\"]"
    exit 1
    ;;
esac
