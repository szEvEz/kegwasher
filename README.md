# kegwasher

```sh
$ kegwasher help

NAME:
   kegwasher - Housekeeping for Homebrew

USAGE:
   kegwasher [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --prune        Remove all cache files (default: false)
   --update       Run 'brew update' and 'brew upgrade' before cleanup (default: false)
   --help, -h     show help
   --version, -v  print only the version (default: false)
```

## Installation

```sh
$ wget -O kegwasher https://github.com/szEvEz/kegwasher/releases/download/v0.0.1/kegwasher-darwin-arm64
$ chmod +x kegwasher
$ mv kegwasher /usr/local/bin/
```

## Usage

```sh
$ kegwasher --update --prune
```

