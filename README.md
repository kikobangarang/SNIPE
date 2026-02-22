# SNIPE

A CLI to kill processes that are using a certain port

## Installation

install release

```bash
mkdir ./snipe-install
cd ./snipe-install

# download, replace with your version and architecture
curl -LO https://github.com/kikobangarang/snipe/releases/download/v1.0.0/snipe_1.0.0_linux_amd64.tar.gz

# extract and install
tar -xzf snipe_1.0.0_linux_amd64.tar.gz
sudo ./install.sh
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
