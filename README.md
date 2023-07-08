# Linux Updater

Linux Updater is a command-line interface (CLI) tool for updating Linux system and managing packages using both apt and brew.

## Prerequisites

- Go 1.20 or higher
- Homebrew (for brew operations)
- Cobra and Viper Go libraries

## Installation

First, clone this repository to your local machine:

```bash
git clone https://github.com/lookatitude/linuxUpdater.git
```

Navigate to the project directory:

```bash
cd linuxUpdater
```

Build the project:

```bash
go build .
```

This will create an executable file in the current directory.

## Usage

You can run the tool using the generated executable. Here are some of the commands you can use:

```bash
# To update apt packages
./linux-updater apt

# To manage brew packages
./linux-updater brew
```

## Configuration

You can configure the behavior of the tool using a `config.yaml` file located in the `~/.linux-updater` directory.

Here is an example `config.yaml`:

```yaml
apt:
update: true
upgrade: true
distUpgrade: true
autoremove: true
autoclean: true

brew:
install: true
update: true
upgrade: true
cleanup: true
doctor: true
```

In this configuration file, you can set whether to perform certain operations when the respective command is run.
