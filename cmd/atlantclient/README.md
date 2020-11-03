Atlantclient
=============

Atlantclient is a custom grpc console client which could be used for accessing to atlantserver. **NOTICE: client was not well tested yet and some stuff may not be work!** 


# Table of Contents

- [Installation](#installation)
- [Usage](#usage)
    - [Fetch](#fetch)
    - [List](#list)


# Installation

```bash
$ go get github.com/morozovcookie/atlant/cmd/atlantclient
```


# Usage

```bash
$ atlantclient --help

Usage:
  atlantclient [command]

Available Commands:
  fetch       
  help        Help about any command
  list        

Flags:
  -h, --help   help for atlantclient

Use "atlantclient [command] --help" for more information about a command.
```


## Fetch

```bash
$ atlantclient fetch --help

Usage:
  atlantclient fetch [flags]

Examples:
atlantclient fetch --host 127.0.0.1:8080 --url http://example.com/sample.csv

Flags:
  -h, --help          help for fetch
      --host string   server host
      --url string    csv resource url
```


## List

```bash
$ atlantclient list --help

Usage:
  atlantclient list [flags]

Examples:
atlantclient list --host 127.0.0.1:8080 --start 0 --limit 100 --sort name:desc,updated_at:asc

Flags:
  -h, --help           help for list
      --host string    server host
      --limit int      items per page
      --sort strings   sorting parameters
      --start int      start position
```
