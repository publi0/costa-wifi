# Costa WiFi CLI

Costa WiFi CLI is a command-line interface tool for managing your Costa Cruise WiFi connection.

## Features

- Login with your Costa card info
- Connect to WiFi
- Disconnect from WiFi
- View active WiFi sessions
- Get version information

## Installation

To install Costa WiFi CLI, you need to have Go installed on your system. Then, you can clone the repository and build the project:

```bash
git clone https://github.com/publi0/costa-wifi.git
cd costa-wifi
go build
```

## Usage

After building the project, you can run the CLI tool using:

```bash
./costa-wifi [command]
```

Available commands:

- `login`: Authenticate with your Costa card info
- `connect`: Connect to WiFi
- `disconnect`: List and disconnect a WiFi session
- `sessions`: Get all active sessions
- `version`: Print the version number
- `help`: Display help information for all commands

For more details on each command, use:

```bash
./costa-wifi [command] --help
```

## Examples

1. Login:
   ```
   ./costa-wifi login
   ```

2. Connect to WiFi:
   ```
   ./costa-wifi connect
   ```
   Or with a specific IP:
   ```
   ./costa-wifi connect --ip 192.168.1.100
   ```

3. View active sessions:
   ```
   ./costa-wifi sessions
   ```

4. Disconnect from WiFi:
   ```
   ./costa-wifi disconnect
   ```

## Configuration

The CLI tool stores configuration data in a JSON file located in your home directory. The file is named `.costa-wifi`.

## Dependencies

This project uses the following main dependencies:

- github.com/spf13/cobra: For creating powerful modern CLI applications
- github.com/pterm/pterm: For beautiful console output
- github.com/goccy/go-yaml: For YAML support
- github.com/golang-jwt/jwt/v4: For JWT token handling

## Disclaimer

This tool is not officially associated with Costa Cruises. Use at your own risk.
