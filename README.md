# trk

Lightweight time tracking and invoicing tool for freelancers/contractors

NOTE: This tool is very early in development still and will have an unstable API/data format.

# Setup

Install trk with:

    go install github.com/minism/trk@latest

Initialize the data folder:

    trk init

# Basic usage

`trk --help` will print all commands, but basics are:

    # Log 4 hours to `my_project` for today
    trk add my_project 4

    # Log another 2 hours for the same project with a message
    trk add my_project 2 -m "Something I did today"

    # Set the total hours on another project from yesterday to 8
    trk set another_project 8 -d "yesterday"

    # Show all work logged this month
    trk log --since "this month"

    # Generate an invoice automatically from the work log
    trk invoice generate

# Design Goals

- Human-readable storage format (YAML and CSV)
- Git for persistence/syncing layer
- Friendly CLI with natural language parsing terms such as "yesterday"
- Per-project and global views
- CLI and TUI both communicate through a core API shell

# Planned features - not yet implemented

See the TODO file for now.
