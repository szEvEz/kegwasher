# kegwasher

```sh
$ kegwasher help

NAME:
   kegwasher - Housekeeping for Homebrew

USAGE:
   kegwasher [global options] command [command options] [arguments...]

VERSION:
   v0.0.2

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --noprune      Do not remove all cached files (default: false)
   --noupdate     Do not run 'brew update' and 'brew upgrade' before cleanup (default: false)
   --help, -h     show help
   --version, -v  print only the version (default: false)
```

## Installation

```sh
$ wget -O kegwasher https://github.com/szEvEz/kegwasher/releases/download/v0.0.2/kegwasher-darwin-arm64
$ chmod +x kegwasher
$ mv kegwasher /usr/local/bin/
```
