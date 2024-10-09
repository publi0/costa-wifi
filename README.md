# Costa WiFi CLI

Costa WiFi CLI is a command-line interface tool for managing your Costa Cruise WiFi connection.

## Features

- Login with your Costa card info
- Connect to WiFi
- Disconnect from WiFi
- View active WiFi sessions
- Get version information

## Installation

There are two ways to install Costa WiFi CLI:

### Option 1: Using `go install`

If you have Go installed on your system, you can use the `go install` command to quickly install Costa WiFi CLI:

```bash
go install github.com/publi0/costa-wifi@latest
```

This will install the latest version of Costa WiFi CLI. After installation, make sure your Go bin directory is in your system's PATH.

### Option 2: Building from source

To build from source, you need to have Go installed on your system. Then, you can clone the repository and build the project:

```bash
git clone https://github.com/publi0/costa-wifi.git
cd costa-wifi
go build
```

After building, you can run the CLI tool using:

```bash
./costa-wifi [command]
```

If you want to install it system-wide after building from source, you can use:

```bash
go install
```

This will install the `costa-wifi` binary in your Go bin directory.

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
- github.com/golang-jwt/jwt/v4: For JWT token handling

## Disclaimer

This tool is not officially associated with Costa Cruises. Use at your own risk.
