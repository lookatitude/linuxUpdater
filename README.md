# Linux Updater

Linux Updater is a command-line interface (CLI) tool for updating Linux system and managing packages. It supports the `apt`, `brew`, `snap`, and `flatpak` package managers.

## Installation

To install Linux Updater, you can download the latest binary from the Releases page and add it to your PATH, or you can build it from source by cloning this repository and running `go build`.

## Usage

You can use Linux Updater via CLI arguments or via a configuration file.

### CLI Arguments

Here are some examples of how you can use Linux Updater with CLI arguments:

```bash
# Update and manage apt packages
linux-updater apt --update --upgrade --dist-upgrade --autoremove --autoclean

# Update and manage brew packages
linux-updater brew --install --update --upgrade --cleanup --doctor

# Refresh snap packages
linux-updater snap --refresh

# Update flatpak packages
linux-updater flatpak --update
```

### Configuration File

You can also specify your options in a `config.yaml` file. Here is an example:

```yaml
apt:
update: true
upgrade: true
distUpgrade: true
autoremove: true
autoclean: true

brew:
install: false
update: true
upgrade: true
cleanup: true
doctor: true

snap:
refresh: true

flatpak:
update: true
```

You can run Linux Updater with this configuration file using the `--config` option:

```bash
linux-updater --config config.yaml
```

This will update and manage your packages according to the options specified in the configuration file.

## Command Documentation

For detailed usage information for each command, see the following documentation:

- [apt](./docs/apt.md)
- [brew](./docs/brew.md)
- [snap](./docs/snap.md)
- [flatpak](./docs/flatpak.md)

## Generating Man Pages

You can generate man pages for all commands by running the following command:

```bash
go run main.go doc --man
```

This will generate man pages in the `./man` directory.

## Contributing

If you would like to contribute to Linux Updater, please feel free to fork this repository, create a feature branch, and open a pull request.

## License

Linux Updater is released under the MIT License. See the `LICENSE` file for more information.
