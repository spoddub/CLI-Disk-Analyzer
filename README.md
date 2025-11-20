### Hexlet tests and linter status:
[![Actions Status](https://github.com/spoddub/go-project-242/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/spoddub/go-project-242/actions)
[![lint](https://github.com/spoddub/go-project-242/actions/workflows/lint.yml/badge.svg)](https://github.com/spoddub/go-project-242/actions/workflows/lint.yml)
[![tests](https://github.com/spoddub/go-project-242/actions/workflows/tests.yml/badge.svg)](https://github.com/spoddub/go-project-242/actions/workflows/tests.yml)
[![asciicast](https://asciinema.org/a/Jx8NIKNcBiAD7lj1nCauJtro2.svg)](https://asciinema.org/a/Jx8NIKNcBiAD7lj1nCauJtro2)

# CLI Disk Analyzer

`CLI Disk Analyzer` is a small training CLI utility written in Go that prints the size of a file or directory.  
It supports recursive size calculation, human readable output and optional inclusion of hidden files.

---

## Tools used

| Tool                                                                          | What it is used for                                             |
| ----------------------------------------------------------------------------- | ---------------------------------------------------------------- |
| [Go](https://go.dev/)                                                         | Language and toolchain.                                         |
| [urfave/cli v3](https://github.com/urfave/cli)                                | Building the CLI interface and flags.                           |
| [golangci-lint](https://golangci-lint.run/)                                   | Fast all in one Go linter.                                      |
| [testify](https://github.com/stretchr/testify)                                | Assertions in unit tests.                                       |
| [GitHub Actions](https://docs.github.com/actions)                             | CI for linting and running tests.                               |
| [Make](https://www.gnu.org/software/make/)                                    | Simple developer tasks: build, test, lint.                      |

---

## Installation and local development

### Requirements

- Go (modern version, for example 1.21 or newer)
- `make`
- `golangci-lint` installed locally (or use `go install` if you prefer)

### Clone the repository

```bash
git clone this-repo
cd this-repo
```

### Build

```bash
make build
# builds ./bin/hexlet-path-size
```
## Usage

Basic examples assuming you are in the project root and built the binary with `make build`:

```bash
# Build the binary
make build

# Size of a single file in bytes
./bin/hexlet-path-size testdata/file1.txt

# Recursive size of a directory
./bin/hexlet-path-size -r testdata

# Human readable size of a file
./bin/hexlet-path-size -H testdata/file1.txt

# Recursive size of a directory including hidden files
./bin/hexlet-path-size -r -a testdata
```

### CLI help

```text
$ hexlet-path-size -h

NAME:
   hexlet-path-size - print size of a file or directory; supports -r (recursive), -H (human-readable), -a (include hidden)

USAGE:
   hexlet-path-size [global options]

GLOBAL OPTIONS:
   --recursive, -r  recursive size of directories (default: false)
   --human, -H      human-readable sizes (auto-select unit) (default: false)
   --all, -a        include hidden files and directories (default: false)
   --help, -h       show help
```

### Flags behavior

- `--recursive, -r`
    - If the path is a directory, sums sizes of all nested files and subdirectories recursively.
- `--human, -H`
    - Prints sizes in a human readable format: B, KB, MB, GB, TB, PB.
    - Examples: `123B`, `1.2MB`, `24.0MB`.
- `--all, -a`
    - Includes hidden files and directories (names starting with `.`).
    - Without this flag hidden entries are ignored.

---

## FormatSize and size calculation

- `GetPathSize(path string, all bool) (int64, error)`
    - For a file: returns its size in bytes.
    - For a directory: sums sizes of files inside (respecting `all` and `recursive` flags in the CLI logic).
    - Hidden files and directories are skipped when `all` is false.

- `FormatSize(size int64, human bool) string`
    - If `human == false` returns plain bytes with unit, for example `123B`.
    - If `human == true` converts to a human readable format with automatic unit selection (B, KB, MB, GB, TB, PB) and one decimal place where needed, for example:
        - `512` -> `512B`
        - `1024` -> `1.0KB`
        - `1234567` -> `1.2MB`

---

## Tests and quality

Run unit tests:

```bash
make test
# or directly:
# go test ./...
```

Run linters:

```bash
make lint
# or directly:
# golangci-lint run
```