# zend — Zend CLI

**zend** is a lightweight, cross-platform command-line tool to interact with **Zend**, a customizable browser homepage. Using zend, you can open your Zend homepage, manage configuration, and integrate it with your terminal workflow.

## Features

- Open your Zend homepage directly from the terminal
- Apply temporary configuration overrides via JSON
- Manage Zend YAML configuration files
- Generate shell autocompletion scripts
- Cross-platform support (Linux, macOS, Windows)

## Installation

Clone the repository and navigate to the CLI directory:

```bash
git clone https://github.com/hoppxi/zend.git
cd zend
```

Make zend executable:

```bash
# Linux / macOS
chmod +x zend
./zend --help

# Windows (PowerShell)
.\zend.exe --help
```

You can optionally move `zend` to a directory in your `PATH` for global access.

## Usage

```bash
zend [flags]
zend [command]
```

### Global Flags

- `-a, --addr string` — Address or port to run the server on (default: random)
- `-h, --help` — Show help for `zend`
- `-v, --version` — Show the version of `zend`

### Commands

#### `open`

Open your Zend homepage in the default browser with optional temporary configuration overrides:

```bash
zend open [flags]

zend open --override '{"theme":"dark"}'
```

---

#### `config`

View, modify, validate, or apply Zend configuration files:

```bash
zend config [file] [flags]
```

**Flags:**

- `-a, --apply` — Apply a config file as the default
- `-g, --get string` — Get a config key value (dot notation supported)
- `-p, --print` — Print current/default config in JSON
- `-s, --set string` — Set a config key=value (dot notation supported)
- `-v, --validate` — Validate config file syntax
- `-h, --help` — Show help for `config`

---

#### `completion`

Generate shell autocompletion scripts for zend:

```bash
zend completion [command]
```

Supported shells: `bash`, `zsh`, `fish`, `powershell`

```bash
zend completion bash > ~/.bashrc
source ~/.bashrc
```
