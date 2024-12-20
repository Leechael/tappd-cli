# Tappd CLI POC

This is a POC (Proof of Concept) for the Tappd CLI, a command-line tool for secure data processing in TEE environments.

## Installation

```bash
go install github.com/your-org/tappd-cli@latest
```

## Usage

Basic usage for quoting files:

```bash
tappd-cli quote -i input.txt -o output.txt
```

### Options

- `-i, --input`: Input file path (required)
- `-o, --output`: Output file path (required)
- `--endpoint`: Custom endpoint for the Tappd service (optional)

## TEE Environment

This tool is designed to work inside the TEE (Trusted Execution Environment) provided by [DStack](https://github.com/Dstack-TEE/dstack).

### Using the Simulator

For development and testing purposes, you can use the [Tappd Simulator](https://github.com/Leechael/tappd-simulator) without a TEE environment.

1. Get the simulator:
```bash
git clone https://github.com/Leechael/tappd-simulator
cd tappd-simulator
```

2. Start the simulator:
```bash
tappd-simulator --cert-file tmp-ca.cert --key-file tmp-ca.key -l unix:///tmp/tappd.sock
```

3. In another terminal, run the CLI with the simulator:
```bash
tappd-cli quote --endpoint /tmp/tappd.sock -i input.txt -o output.txt
```

## Development

Requirements:
- Go 1.20 or higher

To build from source:
```bash
git clone https://github.com/your-org/tappd-cli
cd tappd-cli
go build
```

## License

Apache License 2.0 - See [LICENSE](LICENSE) file for details
