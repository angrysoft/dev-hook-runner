# dhr — Dev Hook Runner

[Polski](README.pl.md)

`dhr` runs hooks defined in YAML files. Hooks can execute commands and shell scripts, use environment variables, and declare dependencies on other hooks.

## Requirements

- Go version compatible with `go.mod` (1.27.1 or newer)
- Bash, if a hook includes a `script`

## Build and test

```sh
make build
make test
```

The executable is created at `bin/dhr`. You can also run the program directly:

```sh
go run . [options] [hook-name]
```

Other available Makefile targets:

```sh
make help
make install
make cross-build
make clean
```

`make install` installs the program to `$GOPATH/bin`, and `make cross-build` builds binaries for Linux, macOS, and Windows.

## Hook configuration

By default, `dhr` looks for hooks in the `.hooks.d` directory under the current directory. Each hook should be defined in a file named `<name>.hook.yml`; use the name without the extension when running a hook.

Example directory structure:

```text
project/
├── .hooks.d/
│   ├── build.hook.yml
│   └── prepare.hook.yml
└── ...
```

Example `.hooks.d/build.hook.yml`:

```yaml
name: build
description: Builds the project
depends_on:
  - prepare
env_file: .hooks.d/build.env.yml
execs:
  - cmd: go
    args:
      - build
      - ./...
script: echo "Build complete"
needs_args: false
```

All fields are optional:

| Field | Description |
| --- | --- |
| `name` | Name displayed when the hook runs. |
| `description` | Short description displayed when listing hooks. |
| `depends_on` | Names of hooks to run before this hook. |
| `env_file` | Path to a YAML file containing environment variables. |
| `execs` | Commands to run in sequence. Each item contains `cmd` and an `args` list. |
| `script` | A script executed using `bash -c`. |
| `needs_args` | If `true`, the hook requires an argument passed with `--arg` or `-a`; the value is available as `HOOK_ARG`. |

Commands in `execs` run in order, followed by the `script`. Commands and scripts inherit the process environment and receive variables from `env_file`.

Example `.hooks.d/build.env.yml`:

```yaml
vars:
  - key: BUILD_MODE
    value: release
```

## Usage

Run a hook from the current directory:

```sh
dhr build
```

Pass an argument to a hook that requires one:

```sh
dhr --arg release build
```

Specify the project directory (an absolute path is recommended):

```sh
dhr --dir /path/to/project build
```

List hooks available in a directory:

```sh
dhr --list
dhr --dir /path/to/project --list
```

| Option | Shorthand | Description |
| --- | --- | --- |
| `--dir` | `-d` | Project directory containing `.hooks.d`; defaults to `.`. |
| `--list` | `-l` | Lists hook names and descriptions. |
| `--arg` | `-a` | Argument to pass to a hook. |

Dependencies run before the hook that declares them. Names in `depends_on` correspond to filenames without the `.hook.yml` suffix.
