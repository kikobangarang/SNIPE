# SNIPE

A CLI to kill processes that are using a certain port

## Installation

install release

```bash

./install.sh

```

## Usage

```bash
snipe [port] [--force]
```

- `port`: The port number you want to free up.
- `--force`: (Optional) If provided, the process will be killed without asking for confirmation.
- `--help`: Display help information.

## Dependencies

This tool requires `ss` (socket statistics) command to be available on the system. It is typically part of the `iproute2` package on most Linux distributions.
