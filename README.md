# Number Guessing Game

Simple number guess game written in Go. Computer randomly selects a number, and
you have to guess it in a limited number of attempts.

This is totally useless and pointless game, but it was made just in sake of
learning Go and its libraries. See [Tech stack](#tech-stack).

## Installation

Pre-built binaries are available for Windows, Linux, macOS and Android on
[GitHub Releases](https://github.com/wadrodrog/guess/releases/latest).

## Usage

```sh
# Show help
guess --help

# Play the game. Available difficulties: easy, medium, hard
guess play --difficulty easy
```

## Configuration

### Lookup order

Configuration lookup order (first found is being used):

1. Flags (`--difficulty hard`)
2. Environment variables (`GUESS_DIFFICULTY=hard`)
3. Config file
    1. `./config.toml`
    2. `guess/config.toml` in config dir
        - Windows: `C:\Users\%USER%\AppData\Roaming`
        - macOS: `$HOME/Library/Application Support`
        - Linux: `$XDG_CONFIG_HOME` or `$HOME/.config`

You can use this lookup order to override values.

### Available keys and values

|Flag|Environment|Config file|Value|Description|
|----|-----------|-----------|-----|-----------|
|`-d, --difficulty`|`GUESS_DIFFICULTY`|`difficulty`|`easy` / `medium` / `hard`|Number of attempts (easy: 10, medium: 5, hard: 3)

### Examples

TOML configuration:
```toml
# config.toml
difficulty = 'medium'
```

Environment variables:
```sh
GUESS_DIFFICULTY=medium guess play
```

Flags:
```sh
guess play --difficulty medium
```

## Build

To build this program, you need [Go] installed on your system.

[Go]: https://go.dev/doc/install

```sh
git clone https://github.com/wadrodrog/guess
cd guess

# Build just for your platform
go build .

# Build for all platforms (binaries in target/)
./build.sh  # Linux/macOS
```

## Tech stack

- [Cobra](https://cobra.dev)
- [Viper](https://github.com/spf13/viper)

## License

This project is licensed under MIT License. See [LICENSE](/LICENSE) file.

